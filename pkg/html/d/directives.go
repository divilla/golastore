package d

import (
	"bytes"
	"github.com/divilla/golastore/pkg/html"
)

type (
	IfDirective struct {
		ExpressionsChildren []expressionChildren
	}

	expressionChildren struct {
		expression bool
		children   []html.Renderer
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

func (d *IfDirective) Render(depth int, bb *bytes.Buffer) {
	for _, ec := range d.ExpressionsChildren {
		if ec.expression {
			for _, child := range ec.children {
				child.Render(depth, bb)
			}
			return
		}
	}
}
