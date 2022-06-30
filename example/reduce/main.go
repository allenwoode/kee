package main

import (
	"fmt"
	"strings"
)

func main() {
	var list = []string{"Fei", "Lin", "Wu"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Println(list, "=>", x)

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})

	fmt.Println(list, "=>", y)
}

func MapStrToStr(arr []string, f func(str string) string) []string {
	var ret []string
	for _, v := range arr {
		ret = append(ret, f(v))
	}
	return ret
}

func MapStrToInt(arr []string, f func(str string) int) []int {
	var ret []int
	for _, v := range arr {
		ret = append(ret, f(v))
	}
	return ret
}
