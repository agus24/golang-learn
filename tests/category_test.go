package tests

import (
	"golang_gin/app/databases/model"
	"golang_gin/tests/helpers"
	"golang_gin/tests/seeders"
	"golang_gin/utils"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryIndex(test *testing.T) {
	testServer, db := GetTestServer()
	defer testServer.Close()

	ResetDB(db)
	builder := helpers.NewHttpRequestBuilder(testServer.URL, db).SetToken(nil)

	test.Run("It should return empty", func(test *testing.T) {
		ResetDB(db)

		resp, _ := builder.BuildAndRun("GET", "/api/v1/app/categories", nil)
		data, err := ParseRequestBody(resp)
		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, 0, len(data["data"].([]any)))
	})

	test.Run("It should return categories", func(test *testing.T) {
		ResetDB(db)

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

		assert.Equal(test, int64(1), int64(data["meta"].(map[string]any)["page"].(float64)))
		assert.Equal(test, int64(15), int64(data["meta"].(map[string]any)["per_page"].(float64)))
	})

	test.Run("it should return searched results", func(test *testing.T) {
		ResetDB(db)

		var categories []model.Categories

		for i := range 20 {
			categories = append(categories, *seeders.SeedCategory(db, "Category "+strconv.Itoa(i)))
		}

		body := map[string]any{"search": "test"}

		resp, _ := builder.BuildAndRun("GET", "/api/v1/app/categories", body)
		data, _ := ParseRequestBody(resp)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, 0, len(data["data"].([]any)))

		body = map[string]any{"search": "Category 11"}

		resp, _ = builder.BuildAndRun("GET", "/api/v1/app/categories", body)
		data, _ = ParseRequestBody(resp)

		categoryData := data["data"].([]any)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, 1, len(data["data"].([]any)))

		assert.Equal(test, categories[11].ID, utils.ParseID(categoryData[0].(map[string]any)["id"]))
		assert.Equal(test, categories[11].Name, categoryData[0].(map[string]any)["name"])
	})
}

func TestCategoryCreate(test *testing.T) {
	testServer, db := GetTestServer()
	defer testServer.Close()

	ResetDB(db)
	builder := helpers.NewHttpRequestBuilder(testServer.URL, db).SetToken(nil)

	test.Run("it should validate input", func(test *testing.T) {
		ResetDB(db)

		resp, _ := builder.BuildAndRun("POST", "/api/v1/app/categories", nil)
		data, _ := ParseRequestBody(resp)

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "validation_failed", data["error_code"])
		assert.Equal(test, "Invalid Input", data["message"])
	})

	test.Run("it should create a category", func(test *testing.T) {
		ResetDB(db)

		body := map[string]any{"name": "Category 1"}

		resp, _ := builder.BuildAndRun("POST", "/api/v1/app/categories", body)
		data, _ := ParseRequestBody(resp)

		categoryData := data["data"].(map[string]any)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "Category 1", categoryData["name"])
	})
}
