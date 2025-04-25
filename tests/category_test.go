package tests

import (
	"fmt"
	"golang_gin/tests/seeders"
	"golang_gin/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryIndex(test *testing.T) {
	testServer, db := GetTestServer()
	defer testServer.Close()

	test.Run("It should return categories", func(test *testing.T) {
		ResetDB(db)

		category := seeders.SeedCategory(db, "Category 1")
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/app/categories", testServer.URL), nil)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		resp, _ := http.DefaultClient.Do(Authenticate(req))
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
