package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"

	"github.com/emilia-miki/become-a-developer-test-task/go/uniqueness"
)

// Similar to bufio.ScanWords, but words only consist of unicode letters and marks
func ScanStrippedWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces and punctuation.
	start := 0

	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])

		if unicode.IsLetter(r) {
			break
		}
	}

	// Scan until non-letter, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		if !unicode.IsLetter(r) {
			return i + width, data[start:i], nil
		}
	}

	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}

func main() {
	reader := norm.NFC.Reader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(ScanStrippedWords)

	outerUniquenessMap := uniqueness.NewMap()
	for scanner.Scan() {
		word := scanner.Text()

		uniquenessMap := uniqueness.NewMap()
		for _, r := range word {
			uniquenessMap.Check(r)
		}

		firstUnique, err := uniquenessMap.GetFirstUniqueRune()
		if err != nil {
			fmt.Printf("The word %s has no unique runes", word)
		}
		if err == nil {
			outerUniquenessMap.Check(firstUnique)
		}
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	firstUnique, err := outerUniquenessMap.GetFirstUniqueRune()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%c\n", firstUnique)
}
