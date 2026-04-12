package middlewares

import (
	"encoding/json"

	tollbooth "github.com/didip/tollbooth/v7"
	"github.com/jarusuraj/schoolsystem/models"

	tollbooth_gin "github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)



func RateLimit(cfg models.RateLimitConfig) gin.HandlerFunc {
	msg := models.Message{
		Status: cfg.Status,
		Body:   cfg.Body,
	}
	jsonMsg, _ := json.Marshal(msg)

	limiter := tollbooth.NewLimiter(cfg.Limit, nil)
	limiter.SetMessageContentType("application/json")
	limiter.SetMessage(string(jsonMsg))

	return tollbooth_gin.LimitHandler(limiter)
}
