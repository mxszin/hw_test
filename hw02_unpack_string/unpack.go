package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var builder strings.Builder
	builder.Grow(len(input) * 2)
	err := unpack([]rune(input), &builder)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func unpack(runes []rune, builder *strings.Builder) error {
	var allowStartingWithInt bool

	if len(runes) == 0 {
		return nil
	}

	// escaping.
	if runes[0] == '\\' {
		allowStartingWithInt = true
		runes = runes[1:]
	}

	// if after escaping str is empty it is invalid string.
	if len(runes) == 0 {
		return ErrInvalidString
	}

	// if first letter is digit but not escaped digit it is invalid string.
	if unicode.IsDigit(runes[0]) && !allowStartingWithInt {
		return ErrInvalidString
	}

	// if it is the last rune.
	if len(runes) == 1 {
		builder.WriteRune(runes[0])
		return nil
	}

	// repeat rune.
	if unicode.IsDigit(runes[1]) {
		repeat := int(runes[1] - '0')
		for i := 0; i < repeat; i++ {
			builder.WriteRune(runes[0])
		}
		return unpack(runes[2:], builder)
	}

	// just put rune into result without repeating.
	builder.WriteRune(runes[0])
	return unpack(runes[1:], builder)
}
