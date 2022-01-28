package main

import (
	"bytes"
	"github.com/divilla/golastore/framework/di"
	"github.com/divilla/golastore/framework/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func main() {
	e := echo.New()
	c := di.NewContainer(e)
	defer c.Close()

	e.Use(middleware.CustomContextMiddleware())
	e.Use(middleware.ZapLoggerMiddleware(c.Logger()))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Static("/assets", "assets")
	e.Debug = true

	//e.GET("/", test)
	//e.GET("/bb", bb)
	//e.GET("/sw", sw)

	e.Logger.Fatal(e.Start(":8000"))
}

func bb(ctx echo.Context) error {
	var bb bytes.Buffer
	for i := 0; i < 100000; i++ {
		func(bb *bytes.Buffer) {
			bb.WriteString("a")
		}(&bb)
	}
	return ctx.HTMLBlob(http.StatusOK, bb.Bytes())
}

func sw(ctx echo.Context) error {
	var sw strings.Builder
	for i := 0; i < 100000; i++ {
		func(bb *strings.Builder) {
			bb.WriteString("a")
		}(&sw)
	}
	return ctx.HTML(http.StatusOK, sw.String())
}

func test(ctx echo.Context) error {
	var sb strings.Builder
	sb.WriteString("test")
	//catalog.NewIndexView(layouts.NewMainLayout()).Render(0, &sb)

	return ctx.HTML(http.StatusOK, sb.String())
}
