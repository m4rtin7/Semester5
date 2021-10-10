package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	rv := make(map[string]int)
	arr := strings.Fields(s)
	for _, str := range arr {
		rv[str] += 1
	}
	return rv
}

func main() {
	for k, v := range WordCount("toto toto toto je je x y z") {
		fmt.Println(k, "value is", v)
	}
}
