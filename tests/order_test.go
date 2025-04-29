package tests

import (
	"golang_gin/app/databases/model"
	"golang_gin/app/helpers"
	"golang_gin/tests/seeders"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderIndex(test *testing.T) {
	testServer, db := GetTestServer()
	defer testServer.Close()

	ResetDB(db)

	builder := helpers.NewHttpRequestBuilder(testServer.URL, db).SetToken(nil)

	test.Run("It should return empty", func(test *testing.T) {
		ResetDB(db)

		resp, _ := builder.BuildAndRun("GET", "/api/v1/app/categories", nil)
		data, err := ParseResponseBody(resp)
		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, 0, len(data["data"].([]any)))
	})

	test.Run("It should return some data with pagination", func(test *testing.T) {
		ResetDB(db)

		orderDate := time.Now()
		seeders.SeedOrder(db, orderDate, "ORD123", 100_000, "Customer 1")

		resp, _ := builder.BuildAndRun("GET", "/api/v1/app/orders", nil)

		data, err := ParseResponseBody(resp)
		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		orderData := data["data"].([]any)
		meta := data["meta"].(map[string]any)

		assert.Equal(test, 1, len(orderData))
		assert.Equal(test, "ORD123", orderData[0].(map[string]any)["order_number"])
		assert.Equal(test, 1, int(meta["page"].(float64)))
		assert.Equal(test, 15, int(meta["per_page"].(float64)))
	})

	test.Run("it can do pagination well", func(test *testing.T) {
		ResetDB(db)

		orderDate := time.Now()

		for i := range 20 {
			seeders.SeedOrder(db, orderDate, "ORD12-"+strconv.Itoa(i), 100_000, "Customer "+strconv.Itoa(i))
		}

		body := map[string]any{"page": "1"}

		resp, _ := builder.BuildAndRun("GET", "/api/v1/app/orders", body)
		data, err := ParseResponseBody(resp)
		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		orderData := data["data"].([]any)
		meta := data["meta"].(map[string]any)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, 15, len(orderData))
		assert.Equal(test, 1, int(meta["page"].(float64)))
		assert.Equal(test, 15, int(meta["per_page"].(float64)))

		body = map[string]any{"page": "2"}
		resp, _ = builder.BuildAndRun("GET", "/api/v1/app/orders", body)

		data, err = ParseResponseBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		orderData = data["data"].([]any)
		meta = data["meta"].(map[string]any)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, 5, len(orderData))
		assert.Equal(test, 15, int(meta["per_page"].(float64)))
		assert.Equal(test, 2, int(meta["page"].(float64)))

		// test pagination error

		body = map[string]any{"page": "asd"}
		resp, _ = builder.BuildAndRun("GET", "/api/v1/app/orders", body)

		data, err = ParseResponseBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "invalid_page_or_per_page", data["error_code"])
	})

	test.Run("it can do search well", func(test *testing.T) {
		ResetDB(db)

		var body map[string]any
		var data map[string]any
		var orderData []any
		var resp *http.Response
		var err error
		var orders []model.Orders

		orderDate := time.Now()

		for i := range 20 {
			order := *seeders.SeedOrder(db, orderDate, "ORD-"+strconv.Itoa(i), 100_000, "Customer "+strconv.Itoa(i))
			orders = append(orders, order)
		}

		body = map[string]any{"search": "ORD-0"}
		resp, _ = builder.BuildAndRun("GET", "/api/v1/app/orders", body)

		data, err = ParseResponseBody(resp)

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		orderData = data["data"].([]any)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, 1, len(data["data"].([]any)))
		assert.Equal(test, "ORD-0", orderData[0].(map[string]any)["order_number"])
		assert.Equal(test, orders[0].ID, int64(orderData[0].(map[string]any)["id"].(float64)))
	})
}
