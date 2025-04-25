package tests

import (
	"testing"
)

func TestCategoryIndex(test *testing.T) {
	testServer := GetTestServer()
	defer testServer.Close()

	// test.Run("It should return categories", func(test *testing.T) {
	// 	resp, err := http.Get(fmt.Sprintf("%s/api/v1/app/categories", testServer.URL))
	//
	// 	if err != nil {
	// 		test.Fatalf("Expected no error, got %v", err)
	// 	}
	//
	// 	assert.Equal(test, 200, resp.StatusCode)
	// })
}
