package d

import (
	"github.com/divilla/golastore/pkg/html"
	"strings"
)

type BlockDirective struct {
	Content string
}

func Block(content string) *BlockDirective {
	return &BlockDirective{
		Content: content,
	}
}

func (b *BlockDirective) Render(depth int, bb *strings.Builder) {
	html.Tabs(depth, bb)
	bb.WriteString(b.Content)
}
