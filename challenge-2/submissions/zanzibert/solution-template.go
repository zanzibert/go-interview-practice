package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	runeSlice := []rune(s)
	length := len(runeSlice)
	for i := range length / 2 {
		temp := runeSlice[i]
		runeSlice[i] = runeSlice[length-(i+1)]
		runeSlice[length-(i+1)] = temp
	}
	return string(runeSlice)
}
