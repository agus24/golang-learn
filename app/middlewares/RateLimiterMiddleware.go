package middlewares

import (
// "fmt"
//
// "github.com/gin-gonic/gin"
// "github.com/ulule/limiter/v3"
// limitergin "github.com/ulule/limiter/v3/drivers/middleware/gin"
// memstore "github.com/ulule/limiter/v3/drivers/store/memory"
)

// func RateLimiterMiddleware() gin.HandlerFunc {
// 	rate, _ := limiter.NewRateFromFormatted("60-M")
// 	store := memstore.NewStore()
//
// 	// Custom key function: use user ID if logged in, else IP
// 	keyFunc := func(c *gin.Context) string {
// 		if userIDVal, exists := c.Get("userID"); exists {
// 			// Assuming userID is stored as string or int
// 			return "user:" + fmt.Sprint(userIDVal)
// 		}
// 		return "ip:" + c.ClientIP()
// 	}
//
// 	return limitergin.NewMiddlewareWithKeyGetter(limiter.New(store, rate), keyFunc)
// }
