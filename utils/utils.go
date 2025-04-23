package utils

import (
	"errors"
	"golang_gin/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParsePageAndPerPage(pageQuery string, perPageQuery string) (*int64, *int64, error) {
	var perPage int64
	if perPageQuery == "" {
		perPage = config.DefaultPerPage
	} else {
		var err error
		perPage, err = strconv.ParseInt(perPageQuery, 10, 64)
		if err != nil {
			return nil, nil, errors.New("Invalid per_page value")
		}
	}

	page, err := strconv.ParseInt(pageQuery, 10, 64)

	if err != nil {
		return nil, nil, errors.New("Invalid per_page value")
	}

	return &page, &perPage, nil
}

func Handle[T any](c *gin.Context, dataFunc func() T, err error, status int) {
	if err != nil {
		var errorMessage string
		println(config.Debug)

		if config.Debug {
			errorMessage = err.Error()
		} else {
			errorMessage = "Internal Server Error"
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": errorMessage})
		return
	}

	c.JSON(status, gin.H{"data": dataFunc()})
}
