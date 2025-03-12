package main

import (
	"fmt"
	"log"
	"strings"
	"os"
)

var Result strings.Builder

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Wrong number of argument")
	}
	input := os.Args[1]
	LoadAscii()
	fmt.Println(generateAscii(input))
}

func generateAscii(input string) string {
	// create a slice of words seperated by new line from input
	words := strings.Split(string(input), "\\n")

	for i := 0; i < len(words); i++ {
		if i >= 1 {
			Result.WriteString("\n")
		}
		readFile(words[i])
	}

	return Result.String()
}

func readFile(word string) {
	for i := 0; i < 8; i++ {
		for _, c := range word {
			ascii, ok := AsciiiMap[c]
			if !ok {
				log.Fatalf("character %c is not in standard.txt", c)
			}
			Result.WriteString(ascii[i])
		}
		if i < 7 {
			Result.WriteString("\n")
		}
	}
}
