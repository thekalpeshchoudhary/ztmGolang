//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type Stats struct {
	letters     int
	digits      int
	spaces      int
	punctuation int
}

type StatRunner func(eachLine string, stats *Stats)

func iterateOver(lines []string, statsFunc StatRunner) {
	stats := Stats{}
	for _, eachLine := range lines {
		statsFunc(eachLine, &stats)
	}
	fmt.Printf("---Stats\nLetter : %v\nDigits : %v\nSpaces : %v\nPunctuation Marks : %v\n", stats.letters, stats.digits, stats.spaces, stats.punctuation)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	statsFunc := func(line string, stats *Stats) {
		for _, eachRune := range line {
			if unicode.IsLetter(eachRune) {
				stats.letters += 1
			}
			if unicode.IsDigit(eachRune) {
				stats.digits += 1
			}
			if unicode.IsSpace(eachRune) {
				stats.spaces += 1
			}
			if unicode.IsPunct(eachRune) {
				stats.punctuation += 1
			}
		}
	}

	lines2 := []string{
		"Te are",
		"6877 letters,",
		"five digits,",
		"12 spaces ,",
		"and 4 punctuation marks in these lines of text!",
	}
	iterateOver(lines, statsFunc)
	iterateOver(lines2, statsFunc)
}
