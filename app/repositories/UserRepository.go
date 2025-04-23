package repositories

import (
	"database/sql"
	"golang_gin/app/ginapp_2/model"
	. "golang_gin/app/ginapp_2/table"
	"golang_gin/utils"

	"github.com/gin-gonic/gin"
	. "github.com/go-jet/jet/v2/mysql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (self *UserRepository) getMultipleUsers(ctx *gin.Context, stmt SelectStatement) ([]model.Users, error) {
	var results []model.Users

	err := stmt.QueryContext(ctx, self.db, &results)

	return results, err
}

func (self *UserRepository) getSingleUser(ctx *gin.Context, stmt SelectStatement) (*model.Users, error) {
	var results []model.Users

	err := stmt.QueryContext(ctx, self.db, &results)

	return &results[0], err
}

func (self *UserRepository) GetUserById(ctx *gin.Context, id int64) (*model.Users, error) {
	stmt := SELECT(
		Users.AllColumns,
	).
		FROM(Users).
		WHERE(Users.ID.EQ(Int64(id)))

	return self.getSingleUser(ctx, stmt)
}

func (self *UserRepository) GetUserByUsername(ctx *gin.Context, username string) (*model.Users, error) {
	stmt := SELECT(
		Users.AllColumns,
	).
		FROM(Users).
		WHERE(Users.Username.EQ(String(username)))

	return self.getSingleUser(ctx, stmt)
}

func (self *UserRepository) CreateUser(ctx *gin.Context, username string, password string, name string) (*model.Users, error) {
	utils.StartTransaction(self.db)
	stmt := Users.INSERT(Users.Username, Users.Password, Users.Name).VALUES(username, password, name)

	res, err := stmt.Exec(self.db)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	user, err := self.GetUserById(ctx, id)
	return user, err
}
