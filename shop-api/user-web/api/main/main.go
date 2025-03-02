package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func GenerateSmsCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	//rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	var learn strings.Builder
	learn.WriteString("hello")
	fmt.Println(learn.String())
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}

	return sb.String()
}
func main() {
	fmt.Println(GenerateSmsCode(5))
}
