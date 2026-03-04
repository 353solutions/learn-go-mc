package main

import (
	"fmt"
	"unicode/utf8"
)

func PrintLine(name string, amount int) {
	if utf8.RuneCountInString(name) > 5 {
		name = name[:2] + "..."
	}
	fmt.Printf("%-5s|%4d\n", name, amount)
}

func main() {
	PrintLine("Alice", 42)
	PrintLine("Bob", 100)
	PrintLine("Alexander", 7)
	PrintLine("Chloè", 17)

	for i, c := range "Chloè!" {
		fmt.Printf("%d: %c\n", i, c)
	}
}
