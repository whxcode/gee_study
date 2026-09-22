package gee

import (
	"fmt"
	"regexp"
	"testing"
)

func TestR1(t *testing.T) {
	re := regexp.MustCompile(`^/get/(?P<name>\w+)/doc/?(?P<id>\d+)$`)

	match := re.FindStringSubmatch("/get/1234/doc/10")

	fmt.Println(re.SubexpNames())

	if match == nil || len(match) == 0 {
		t.Fatalf("No match found")
		return
	}

	// 把分组名和值对应起来
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	fmt.Println(result)
}
