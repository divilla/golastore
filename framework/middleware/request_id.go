package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestID() echo.MiddlewareFunc {
	return middleware.RequestID()
}
