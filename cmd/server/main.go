package main

import (
	"bytes"
	"github.com/divilla/golastore/pkg/html/d"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/e"
	"github.com/labstack/echo/v4"
)

func main() {
	s := echo.New()
	s.Static("/assets", "assets")

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	s.GET("/", test)
	s.GET("/bb", bb)
	s.GET("/sw", sw)

	s.Logger.Fatal(s.Start(":8000"))
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
	start := time.Now()
	json := gjson.Parse(`{
	"title": "Hello Bulma"
}`)
	l := html.NewLayout(
		html.Block("<!DOCTYPE html>"),
		e.Html("en-US").Children(
			e.Head().Children(
				e.Meta(a.A{K: "charset", V: "utf-8"}),
				e.Meta(a.A{K: "name", V: "viewport"}, a.A{K: "content", V: "width=device-width, initial-scale=1"}),
				e.Title().Text(json.Get("title").String()),
				e.Link(a.A{K: "shortcut icon", V: "https://fdn.gsmarena.com/imgroot/static/favicon.ico"}),
				e.Link(a.Rel("stylesheet"), a.Href("/assets/css/style.css")),
			),
			e.Body().Children(
				e.Section(a.Class("section"), a.Style("padding-top: 1.5rem")).Children(
					e.Div(a.Class("columns")).Children(
						e.Div(a.Class("column is-2")).Children(
							e.Img(a.Src("/assets/images/market-logo.png"), a.Width("142"), a.Height("42")),
						),
						e.Div(a.Class("column is-5")).Children(
							e.Div(a.Class("field")).Children(
								e.P(a.Class("control has-icons-right")).Children(
									e.Input(a.Class("input is-rounded"), a.Type("text"), a.Placeholder("Search")),
									e.Span(a.Class("icon is-small is-right has-text-danger-dark")).Children(
										e.I(a.Class("fas fa-search")),
									),
								),
							),
						),
						e.Div(a.Class("column is-5 is-flex is-justify-content-flex-end")).Children(
							e.P(a.Class("buttons")).Children(
								e.A(a.Class("button is-white")).Children(
									e.Span(a.Class("icon is-small has-text-success-dark")).Children(
										e.I(a.Class("fas fa-sign-in-alt")),
									),
									e.Span().Text("Sign in"),
								),
								e.A(a.Class("button is-white")).Children(
									e.Span(a.Class("icon is-small has-text-info-dark")).Children(
										e.I(a.Class("fas fa-user-plus")),
									),
									e.Span().Text("Sign up"),
								),
								e.A(a.Class("button is-white")).Children(
									e.Span(a.Class("icon is-small has-text-danger-dark")).Children(
										e.I(a.Class("fas fa-sign-out-alt")),
									),
									e.Span().Text("Sign out"),
								),
								e.A(a.Class("button is-dark")).Children(
									e.Span(a.Class("badge is-danger"), a.Title("Cart Items")).Text("8"),
									e.Span().Text("0,00 kn"),
									e.Span(a.Class("icon is-small has-text-warning")).Children(
										e.I(a.Class("fas fa-shopping-cart")),
									),
								),
							),
						),
					),
					//e.Nav(a.Class("navbar"), a.Role("navigation"), a.AriaLabel("main-navigation")).Children(
					//	e.Div(a.Class("navbar-brand")).Children(
					//		e.A(a.Class("navbar-item"), a.Href("/")).Children(
					//			e.Img(a.Src("/assets/images/market-logo.png"), a.Width("142"), a.Height("42")),
					//		),
					//		e.A(a.Class("navbar-burger"), a.Role("button"), a.AriaExpanded("false"), a.DataTarget("navbarBasicExample")).Children(
					//			e.Span(a.AriaHidden("true")),
					//			e.Span(a.AriaHidden("true")),
					//			e.Span(a.AriaHidden("true")),
					//		),

					//	),
					//	e.Div(a.Id("navbarBasicExample"), a.Class("navbar-menu")).Children(
					//		e.Div(a.Class("navbar-start")).Children(
					//			e.A(a.Class("navbar-item")).T(
					//				"Home",
					//			),
					//			e.A(a.Class("navbar-item")).T(
					//				"Documentation",
					//			),
					//		),
					//	),
					//),
					//container(nil)...
				),
				e.Section(a.Class("section"), a.Style("padding-top: 1.5rem")).Children(
					d.If(false,
						html.Block("This is text"),
					).Else(
						html.Block(`<div style="color:red">This is text</div>`),
					),
				),
			),
		),
	)

	var bb strings.Builder
	l.Render(0, &bb)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s\n", elapsed)

	return ctx.HTML(http.StatusOK, bb.String())
}

func container(params map[string]string) []*e.E {
	var c []*e.E
	for i := 0; i < 60; i++ {
		c = append(c,
			e.Section(a.Class("section")).Children(
				e.Div(a.Class("container")).Children(
					e.H1(a.Class("title")).Text("Hello pussies!"),
					e.P(a.Class("subtitle")).Text("My first website with <strong>Bulma</strong>!"),
				),
			),
		)
	}

	return c
}
