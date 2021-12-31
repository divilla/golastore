package html

import "bytes"

func Tabs(depth int, bb *bytes.Buffer) {
	for i := 0; i < depth; i++ {
		bb.WriteString("    ")
	}
}
