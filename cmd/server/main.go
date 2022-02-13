package main

import (
	"github.com/divilla/golastore/framework/di"
	"github.com/divilla/golastore/framework/middleware"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	e := echo.New()
	c := di.NewContainer(e)
	defer c.Close()

	e.Use(middleware.CustomContextMiddleware())
	e.Use(middleware.ZapLoggerMiddleware(c.Logger()))
	//e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Static("/assets", "assets")
	e.Debug = true

	c.Logger().Sugar().With(zap.Stack("stack")).Fatal(e.Start(":8000"))
}
