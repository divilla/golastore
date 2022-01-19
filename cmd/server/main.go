package main

import (
	"bytes"
	"github.com/divilla/golastore/framework/di"
	"github.com/divilla/golastore/framework/middleware"
	"github.com/divilla/golastore/internal/views/layouts"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
)

func main() {
	mc := di.InitMainContainer()

	e := echo.New()
	e.Use(middleware.SetCustomContext())
	e.Use(middleware.ZapLogger(mc.Logger()))
	e.Static("/assets", "assets")

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.GET("/", test)
	e.GET("/bb", bb)
	e.GET("/sw", sw)

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
	jr := gjson.Parse(`{
	"app": {},
	"ctx": {},
	"mod": {
		"title": "Hello Bulma"
	}
}`)

	layouts.MainLayout(&sb, &jr)

	return ctx.HTML(http.StatusOK, sb.String())
}
