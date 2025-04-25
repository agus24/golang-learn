package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(test *testing.T) {
	testServer, _ := GetTestServer()
	defer testServer.Close()

	test.Run("Addition", func(test *testing.T) {
		result := 2 + 2
		expected := 4

		assert.Equal(test, result, expected)
	})

	test.Run("it should return 200 when health is ok", func(test *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/__health", testServer.URL))

		if err != nil {
			test.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(test, 200, resp.StatusCode)
	})
}
