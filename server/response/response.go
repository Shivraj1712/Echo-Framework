package response

import "github.com/labstack/echo/v5"

type Success struct {
	Data map[string]any
}

type Failure struct {
	ErrorStruct struct {
		Message string
	}
}

func SuccessResponse(ctx *echo.Context, Data map[string]any, statusCode int) error {
	return ctx.JSON(statusCode, &Success{
		Data: Data,
	})
}

func FailureResponse(ctx *echo.Context, message string, statusCode int) error {
	return ctx.JSON(statusCode, &Failure{
		ErrorStruct: struct{ Message string }{
			Message: message,
		},
	})
}
