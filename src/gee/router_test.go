package gee

import (
	"testing"
)

func TestParsePattern(t *testing.T) {
	r := parsePattern("/feafa/:name")

	t.Log("输出:", r)
}
