package c

import (
	"bytes"
	"github.com/divilla/golastore/pkg/html"
)

type (
	TextComponent struct {
		Text string
	}

	HTMLComponent struct {
		HTML string
	}
)

func Text(val string) *TextComponent {
	return &TextComponent{
		Text: val,
	}
}

func (c *TextComponent) Render(depth int, bb *bytes.Buffer) {
	bb.WriteString(c.Text)
}

func HTML(val string) *HTMLComponent {
	return &HTMLComponent{
		HTML: val,
	}
}

func (c *HTMLComponent) Render(depth int, bb *bytes.Buffer) {
	html.Tabs(depth, bb)
	bb.WriteString(c.HTML + "\n")
}
