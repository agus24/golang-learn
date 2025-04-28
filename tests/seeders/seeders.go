package seeders

import (
	"database/sql"
	"golang_gin/app/databases/model"
	"golang_gin/app/repositories"
	"golang_gin/app/requests"
	"golang_gin/app/services"
	"log"
	"runtime/debug"
	"time"
)

func SeedUser(db *sql.DB, username string, password string, name string) *model.Users {
	service := services.NewUserService(repositories.NewUserRepository(db))

	user, err := service.CreateUser(requests.UserCreateRequest{
		Username: username,
		Password: password,
		Name:     name,
	})

	if err != nil {
		log.Fatal(err, string(debug.Stack()))
		return nil
	}

	return user
}

func SeedCategory(db *sql.DB, name string) *model.Categories {
	service := services.NewCategoryService(repositories.NewCategoryRepository(db))

	category, err := service.CreateCategory(name)

	if err != nil {
		log.Fatal(err, string(debug.Stack()))
		return nil
	}

	return category
}

func SeedOrder(db *sql.DB, date time.Time, orderNumber string, grandTotal int, customerName string) *model.Orders {
	repository := repositories.NewOrderRepository(db)

	result, err := repository.CreateOrder(date, orderNumber, grandTotal, customerName)

	if err != nil {
		log.Fatal(err, string(debug.Stack()))
		return nil
	}

	return result
}
