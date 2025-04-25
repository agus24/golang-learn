package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(test *testing.T) {
	testServer, db := GetTestServer()
	defer testServer.Close()

	test.Run("It should return validation error", func(test *testing.T) {
		ResetDB(db)
		resp, err := http.Post(fmt.Sprintf("%s/api/v1/auth/login", testServer.URL), "application/json", nil)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		data, err := ParseRequestBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, data["error"], "Invalid request")
	})

	test.Run("It should return wrong username", func(test *testing.T) {
		ResetDB(db)
		loginBody, err := PrepareBody(map[string]any{
			"username": "test",
			"password": "test",
		})

		resp, err := http.Post(fmt.Sprintf("%s/api/v1/auth/login", testServer.URL), "application/json", loginBody)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		data, err := ParseRequestBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, data["error"], "Invalid credentials")
	})
}
