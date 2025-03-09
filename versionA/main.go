	package main

	import (
		"fmt"
		"os"
		"log"
		"strings"
	)
	// a map that links each ascii character to an index (first line of ascii art) in standard.txt
	var AsciiChars = make(map[rune]int)
	var Result strings.Builder

	func main() {
		if len(os.Args) != 2 {
			log.Fatal("Wrong number of argument")
		}
		input := os.Args[1]
		generateAsciiMap()
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

	// readFile open standard.txt and lookup ascii art one character at a time then store them into Result
	func readFile(word string) {
		// store standard.txt into a byte slice f
		f, err := os.ReadFile("standard.txt")
		if err != nil {
			log.Fatal(err)
		}
		// create a slice of lines from standard.txt file
		fileLines := strings.Split(string(f), "\n")
		for i := 0; i < 8; i++ {
			for _, c := range word {
				lineN, ok := AsciiChars[c]
				if !ok {
					log.Fatalf("character %c is not in standard.txt", c)
				}
				Result.WriteString(fileLines[lineN + i])
			}
			if i < 7 {
				Result.WriteString("\n")
			}
		}
	}

	// generateAsciiMap function will populate AsciiChars map with (Ascii char ) rune -> (index - line number in standard.txt) int
	func generateAsciiMap() {
		var asciiCode = ' '
		const totalChars = 95
		for i := 0; i < totalChars; i++ {
			AsciiChars[asciiCode] = (9*i)+1
			asciiCode += 1
		}
	}
