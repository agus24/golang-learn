package middlewares

import (
	"errors"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func CatchAllMiddleware(ctx *gin.Context) {
	defer func() {
		var message string
		if v := recover(); v != nil {
			switch e := v.(type) {
			case string:
				message = "Recovered panic with string:"
			case error:
				message = "Recovered panic with error:" + e.Error()
			default:
				message = "Recovered panic with unknown type:"
			}

			if gin.Mode() == gin.DebugMode {
				ctx.AbortWithStatusJSON(500, gin.H{"error": message, "error_code": "internal_server_error", "stack": string(debug.Stack())})
				panic(errors.New(message))
				return
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
		}
	}()

	ctx.Next()

	// If there are errors in context
	if len(ctx.Errors) > 0 {
		ctx.JSON(-1, gin.H{
			"errors": ctx.Errors.JSON(),
		})
		return
	}
}
