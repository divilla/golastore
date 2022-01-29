package d

import (
	"github.com/divilla/golastore/pkg/html"
	"strings"
)

type (
	IfDirective struct {
		ExpressionsChildren []expressionChildren
	}

	expressionChildren struct {
		expression bool
		children   []html.Renderer
	}

	IfsDirective struct {
		ExpressionsTexts []expressionText
	}

	expressionText struct {
		expression bool
		text       string
	}
)

func If(expression bool, children ...html.Renderer) *IfDirective {
	return &IfDirective{
		ExpressionsChildren: []expressionChildren{
			{
				expression: expression,
				children:   children,
			},
		},
	}
}

func (d *IfDirective) ElseIf(expression bool, children ...html.Renderer) *IfDirective {
	d.ExpressionsChildren = append(d.ExpressionsChildren, expressionChildren{
		expression: expression,
		children:   children,
	})
	return d
}

func (d *IfDirective) Else(children ...html.Renderer) html.Renderer {
	d.ExpressionsChildren = append(d.ExpressionsChildren, expressionChildren{
		expression: true,
		children:   children,
	})
	return d
}

func (d *IfDirective) Render(depth int, bb *strings.Builder) {
	for _, ec := range d.ExpressionsChildren {
		if ec.expression {
			for _, child := range ec.children {
				child.Render(depth, bb)
			}
			return
		}
	}
}

func Ifs(expression bool, text string) *IfsDirective {
	return &IfsDirective{
		ExpressionsTexts: []expressionText{
			{
				expression: expression,
				text:       text,
			},
		},
	}
}

func (d *IfsDirective) ElseIf(expression bool, text string) *IfsDirective {
	d.ExpressionsTexts = append(d.ExpressionsTexts, expressionText{
		expression: expression,
		text:       text,
	})
	return d
}

func (d *IfsDirective) Else(text string) {
	d.ExpressionsTexts = append(d.ExpressionsTexts, expressionText{
		expression: true,
		text:       text,
	})
}

func (d *IfsDirective) String() string {
	for _, ec := range d.ExpressionsTexts {
		if ec.expression {
			return ec.text
		}
	}

	return ""
}
