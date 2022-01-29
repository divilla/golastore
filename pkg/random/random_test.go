package random

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Len(t, MustString(32), 32)
	assert.Len(t, MustURI(32), 32)
	assert.Len(t, MustURINew(32), 32)
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), MustString(8, Numeric))
}
