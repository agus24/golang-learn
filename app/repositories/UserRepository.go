package repositories

import (
	"database/sql"
	"golang_gin/app/ginapp_2/model"
	. "golang_gin/app/ginapp_2/table"

	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/go-jet/jet/v2/mysql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (self *UserRepository) GetUserById(ctx *gin.Context, id int64) (*model.Users, error) {
	stmt := SELECT(
		Users.AllColumns,
	).
		FROM(Users).
		WHERE(Users.ID.EQ(Int64(id)))

	var results []model.Users

	err := stmt.QueryContext(ctx, self.db, &results)
	if err != nil {
		return nil, err
	}

	return &results[0], nil
}

func (self *UserRepository) GetUserByUsername(ctx *gin.Context, username string) (*model.Users, error) {
	stmt := SELECT(
		Users.AllColumns,
	).
		FROM(Users).
		WHERE(Users.Username.EQ(String(username)))

	var results []model.Users

	err := stmt.QueryContext(ctx, self.db, &results)
	if err != nil {
		return nil, err
	}

	fmt.Println(results[0])

	return &results[0], nil
}
