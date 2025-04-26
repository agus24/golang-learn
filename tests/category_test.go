package tests

import (
	"golang_gin/tests/helpers"
	"golang_gin/tests/seeders"
	"golang_gin/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryIndex(test *testing.T) {
	testServer, db := GetTestServer()
	defer testServer.Close()

	test.Run("It should return categories", func(test *testing.T) {
		ResetDB(db)

		user := seeders.SeedUser(db, "user1", "password", "user 1")

		builder := helpers.NewHttpRequestBuilder(testServer.URL).
			SetToken(GetToken(user))

		category := seeders.SeedCategory(db, "Category 1")

		resp, _ := builder.BuildAndRun("GET", "/api/v1/app/categories", nil)
		data, err := ParseRequestBody(resp)
		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		categoryData := data["data"].([]any)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, category.ID, utils.ParseID(categoryData[0].(map[string]any)["id"]))
		assert.Equal(test, category.Name, categoryData[0].(map[string]any)["name"])
	})
}
