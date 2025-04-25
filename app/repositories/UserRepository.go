package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"golang_gin/app/databases/model"
	. "golang_gin/app/databases/table"
	"golang_gin/utils"

	. "github.com/go-jet/jet/v2/mysql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (self *UserRepository) getMultipleUsers(stmt SelectStatement) ([]model.Users, error) {
	var results []model.Users

	err := stmt.Query(self.db, &results)

	return results, err
}

func (self *UserRepository) getSingleUser(stmt SelectStatement) (*model.Users, error) {
	var results []model.Users

	err := stmt.Query(self.db, &results)

	if len(results) == 0 {
		return nil, errors.New("User not found")
	}

	fmt.Println("🔍 Jet SQL:", stmt.DebugSql())

	self.db.Exec("DELETE FROM users")
	self.db.Exec("INSERT INTO users (username, name, password) VALUES (?, ?, ?)", "test123", "asdfasdf", "asdfasdf")

	var dbName string
	_ = self.db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	fmt.Println("🧪 ACTUAL CONNECTED DB:", dbName)

	return &results[0], err
}

func (self *UserRepository) GetUserById(id int64) (*model.Users, error) {
	stmt := SELECT(
		Users.AllColumns,
	).
		FROM(Users).
		WHERE(Users.ID.EQ(Int64(id)))

	return self.getSingleUser(stmt)
}

func (self *UserRepository) GetUserByUsername(username string) (*model.Users, error) {
	stmt := SELECT(
		Users.AllColumns,
	).
		FROM(Users).
		WHERE(Users.Username.EQ(String(username)))

	return self.getSingleUser(stmt)
}

func (self *UserRepository) CreateUser(username string, password string, name string) (sql.Result, error) {
	utils.StartTransaction(self.db)
	stmt := Users.INSERT(Users.Username, Users.Password, Users.Name).VALUES(username, password, name)

	return stmt.Exec(self.db)
}
