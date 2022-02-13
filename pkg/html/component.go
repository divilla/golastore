package html

import "strings"

type (
	Component struct {
		html []Renderer
	}

	IComponent interface {
		Render(depth int, bb *strings.Builder)
	}
)

func NewComponent(html ...Renderer) *Component {
	return &Component{
		html: html,
	}
}

func (v *Component) HTML(html ...Renderer) {
	v.html = html
}

func (v *Component) Render(depth int, bb *strings.Builder) {
	for k, elm := range v.html {
		if k > 0 {
			bb.WriteString("\n")
		}
		elm.Render(depth, bb)
	}
}
