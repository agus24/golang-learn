package requests

import (
	"golang_gin/app/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicValidation[T any](ctx *gin.Context) {
	var input T

	inputRaw, err := DefaultValidationRule[T](ctx)

	if err != nil {
		return
	}

	input = *inputRaw

	ctx.Set("validated", input)
	ctx.Next()
}

func DefaultValidationRule[T any](ctx *gin.Context) (*T, error) {
	var input T

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ValidationError(err))
		return nil, err
	}

	return &input, nil
}
