package main

import (
	"fmt"
	"regexp"
)

const (
	T1 = iota // 0
	T2        // 1
	T3        // 2
	T4        // 3
)

func init() {
	fmt.Println(10)
}

func main() {
	re := regexp.MustCompile(`ab?`)
	fmt.Println(re.FindStringIndex("tblett"))
	fmt.Println(re.FindStringIndex("foo") == nil)
}
