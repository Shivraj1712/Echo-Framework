package middleware

import (
	"net/http"

	"github.com/Shivraj1712/Echo-Framework/response"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func RateLimiter(e *echo.Echo) {
	e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStore(20),
		IdentifierExtractor: func(c *echo.Context) (string, error) {
			return c.RealIP(), nil
		},
		DenyHandler: func(c *echo.Context, identifier string, err error) error {
			return response.FailureResponse(c, "Too Many Requests! Try again later", http.StatusTooManyRequests)
		},
	}))
}
