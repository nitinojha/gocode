package main

import "fmt"

func main() {
	str := "hello world"
	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[char]++
	}

	for char, count := range charCount {
		fmt.Printf("%c: %d\n", char, count)
	}
}
