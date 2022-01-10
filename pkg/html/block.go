package html

import "strings"

type B struct {
	Content string
}

func Block(content string) *B {
	return &B{
		Content: content,
	}
}

func (b *B) Render(depth int, bb *strings.Builder) {
	Tabs(depth, bb)
	bb.WriteString(b.Content)
}
