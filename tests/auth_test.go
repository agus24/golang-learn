package tests

import (
	"fmt"
	"golang_gin/tests/helpers"
	"golang_gin/tests/seeders"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(test *testing.T) {
	testServer, db := GetTestServer()
	defer testServer.Close()

	builder := helpers.NewHttpRequestBuilder(testServer.URL, db).SetToken(nil)

	test.Run("It should return validation error", func(test *testing.T) {
		ResetDB(db)
		resp, err := builder.BuildAndRun("POST", "/api/v1/auth/login", nil)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		data, err := ParseResponseBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, data["error_code"], "validation_failed")
	})

	test.Run("It should return wrong username", func(test *testing.T) {
		ResetDB(db)
		loginBody, err := helpers.PrepareBody(map[string]any{
			"username": "test",
			"password": "test",
		})

		resp, err := http.Post(fmt.Sprintf("%s/api/v1/auth/login", testServer.URL), "application/json", loginBody)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		data, err := ParseResponseBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, data["error"], "Invalid credentials")
	})

	test.Run("It should return invalid password", func(test *testing.T) {
		ResetDB(db)
		user := seeders.SeedUser(db, "user1", "password", "user 1")

		loginBody, err := helpers.PrepareBody(map[string]any{
			"username": user.Username,
			"password": "test",
		})

		resp, err := http.Post(fmt.Sprintf("%s/api/v1/auth/login", testServer.URL), "application/json", loginBody)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		data, err := ParseResponseBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, data["error"], "Invalid credentials")
	})

	test.Run("It should return valid user", func(test *testing.T) {
		ResetDB(db)
		user := seeders.SeedUser(db, "user1", "password", "user 1")

		loginBody, err := helpers.PrepareBody(map[string]any{
			"username": user.Username,
			"password": "password",
		})

		resp, err := http.Post(fmt.Sprintf("%s/api/v1/auth/login", testServer.URL), "application/json", loginBody)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		data, err := ParseResponseBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		userResponse := data["user"].(map[string]any)

		assert.Equal(test, 200, resp.StatusCode)
		assert.True(test, data["token"] != nil)
		assert.Equal(test, userResponse["username"], user.Username)
		assert.Equal(test, userResponse["name"], user.Name)
	})
}
