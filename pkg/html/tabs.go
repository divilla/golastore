package html

import (
	"strings"
)

func Tabs(depth int, bb *strings.Builder) {
	for i := 0; i < depth; i++ {
		bb.WriteString("  ")
	}
}
