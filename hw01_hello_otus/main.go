package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	text := "Hello, OTUS!"
	reversedText := stringutil.Reverse(text)

	fmt.Print(reversedText)
}
