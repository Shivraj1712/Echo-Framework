package middleware

import (
	"net/http"

	"github.com/Shivraj1712/Echo-Framework/response"
	"github.com/labstack/echo/v5"
)

func ErrorHandler(c *echo.Context, err error) error {
	code := http.StatusInternalServerError
	message := echo.ErrInternalServerError.Error()
	if Err, ok := err.(*echo.HTTPError); ok {
		code = Err.StatusCode()
		message = Err.Message
	}
	return response.FailureResponse(c, message, code)
}
