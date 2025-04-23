package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware интерфейс для всех middleware
type Middleware interface {
	TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc
}

// provider структура, реализующая Middleware
type provider struct {
}

// NewMiddlewareProvider создает новый экземпляр провайдера middleware
func NewMiddlewareProvider() Middleware {
	return &provider{}
}

// TimeoutMiddleware ограничивает выполнение обработчика по времени
func (p *provider) TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		ch := make(chan struct{})

		c.Request = c.Request.WithContext(ctx)

		go func() {
			c.Next()
			close(ch)
		}()

		select {
		case <-ch:
			return
		case <-ctx.Done():
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "request timeout"})
		}
	}
}
