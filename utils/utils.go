package utils

import (
	"errors"
	"golang_gin/config"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/goutil/dump"
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

func ParseToInt(ctx *gin.Context, param string) *int64 {
	if param == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return nil
	}

	paramInt, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return nil
	}

	return &paramInt
}

func Handle[T any](c *gin.Context, dataFunc func() T, err error, status int) {
	if err != nil {
		if gin.Mode() == gin.DebugMode {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "trace": string(debug.Stack())})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	c.JSON(status, dataFunc())
}

func Dump(data any, isFatal bool) {
	dump.P(data)
	if isFatal {
		log.Fatal("FATAL because dump")
	}
}

func GenerateValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			errors[e.Field()] = e.Tag()
		}
	} else {
		errors["error"] = err.Error()
	}

	return errors
}
