package middleware

import (
	"github.com/labstack/echo"
)

//UserLoginRequired to check user whether login
func UserLoginRequired(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return handler(ctx)
	}
}
