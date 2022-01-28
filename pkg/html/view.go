package html

import "strings"

type (
	View struct {
		children []Renderer
	}

	IView interface {
		Render(depth int, bb *strings.Builder)
	}
)

func NewView(children ...Renderer) *View {
	return &View{
		children: children,
	}
}

func (v *View) Render(depth int, bb *strings.Builder) {
	for k, elm := range v.children {
		if k > 0 {
			bb.WriteString("\n")
		}
		elm.Render(depth, bb)
	}
}
