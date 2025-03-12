package main

import (
	"log"
	"os"
	"strings"
)

// AsciiiMap is a map that links each ascii character to slice of 8 strings, each string trpresent one line of Ascii Art
var AsciiiMap = make(map[rune][]string)
// AsciiChars is a map that links each ascii character to an index (first line of ascii art) in standard.txt
var AsciiChars = make(map[rune]int)
// lines in standard.txt will be copied to FileLines
var FileLines []string

// LoadAscii will pre load all ascii Art into AsciiiMap for later use
func LoadAscii() {
	f, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(f), "\n")
	FileLines = append(FileLines, lines...)
	generateAsciiMap()
	for k, v := range AsciiChars {
		AsciiiMap[k] = extractChar(k, v)
	}
}

func extractChar(k rune, v int) []string {
	result := make([]string, 8, 8)
	for i := 0; i < 8; i++ {
		result[i] = FileLines[i+v]
	}
	return result
}

// generateAsciiMap function will populate AsciiChars map with (Ascii char ) rune -> (index - line number in standard.txt) int
func generateAsciiMap() {
	var asciiCode = ' '
	const totalChars = 95
	for i := 0; i < totalChars; i++ {
		AsciiChars[asciiCode] = (9 * i) + 1
		asciiCode += 1
	}
}
