package main

import (
	"github.com/divilla/render/pkg/html"
	"github.com/divilla/render/pkg/html/a"
	"github.com/divilla/render/pkg/html/e"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.GET("/", test)

	e.Logger.Fatal(e.Start(":8000"))
}

func test(c echo.Context) error {
	start := time.Now()
	d := html.NewDocument(
		e.Html("en-US").C(
			e.Head().C(
				e.Meta(a.A{K: "charset", V: "utf-8"}),
				e.Meta(a.A{K: "name", V: "viewport"}, a.A{K: "content", V: "width=device-width, initial-scale=1"}),
				e.Title().T("Hello Bulma!"),
				e.Link(a.A{K: "shortcut icon", V: "https://fdn.gsmarena.com/imgroot/static/favicon.ico"}),
				e.Link(a.Rel("stylesheet"), a.Href("https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css")),
			),
		),
		e.Body().C(
			container(nil)...
		),
	)

	elapsed := time.Since(start)
	res := d.Render()
	log.Printf("Binomial took %s\n", elapsed)

	return c.HTMLBlob(http.StatusOK, res)
}

func container(params map[string]string) []*e.E {
	var c []*e.E
	for i:=0; i<60; i++ {
		c = append(c,
			e.Section(a.Class("section")).C(
				e.Div(a.Class("container")).C(
					e.H1(a.Class("title")).T("Hello pussies!"),
					e.P(a.Class("subtitle")).T("My first website with <strong>Bulma</strong>!"),
				),
			),
		)
	}

	return c
}