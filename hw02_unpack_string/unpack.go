package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder
	// make a slice from the string to determine
	// a last symbol in loop
	list := []rune(input)

	// mark char as escaped between iterations
	var charIsEscaped bool

	for i, currentRune := range list {
		currentChar := string(currentRune)
		currentCharIsDigit := unicode.IsDigit(currentRune)

		isLastRune := i == len(list)-1

		currentCharIsEscape := currentChar == "\\"
		lastCharIsEscape := isLastRune && currentCharIsEscape

		//   ↓
		// aa\
		if lastCharIsEscape {
			return "", ErrInvalidString
		}

		//     ↓ ↓
		// qwe\\\5
		if charIsEscaped {
			// now it is a letter
			currentCharIsEscape = false
			// now it is a letter
			currentCharIsDigit = false
			// reset for next iteration
			charIsEscaped = false
		}

		//    ↓
		// qwe\\5
		if currentCharIsEscape {
			// set for next iteration
			charIsEscaped = true
			continue
		}

		var nextRune rune
		if !isLastRune {
			nextRune = list[i+1]
		}
		nextChar := string(nextRune)
		nextCharIsDigit := unicode.IsDigit(nextRune)

		currentAndNextCharsAreDigit := currentCharIsDigit && nextCharIsDigit
		firstCharIsDigit := i == 0 && currentCharIsDigit

		// ↓     ↓      ↓
		// 1a || 12 || a11a
		if firstCharIsDigit || currentAndNextCharsAreDigit {
			return "", ErrInvalidString
		}

		// ↓
		// a2
		if nextCharIsDigit {
			repeatCount, _ := strconv.Atoi(nextChar)
			result.WriteString(strings.Repeat(currentChar, repeatCount))
			continue
		}

		//  ↓
		// a3a
		if currentCharIsDigit {
			continue
		}

		// ↓
		// aa
		result.WriteString(currentChar)
	}

	return result.String(), nil
}
