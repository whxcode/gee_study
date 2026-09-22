package gee

import "testing"

func TestR1(t *testing.T) {
	add := func() int {
		return 3
	}

	result := add()

	if result == 3 {
		t.Errorf("期望结果 3，实际 %d", result)
	}
}
