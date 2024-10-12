package fmtutil

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func ToSearchable(text string) (textInSearchableFormat string) {
	normalizedText := norm.NFD.String(
		text,
	)
	normalizedLetters := make([]rune, 0, len(normalizedText))

	for _, r := range normalizedText {
		if !unicode.Is(unicode.Mn, r) {
			normalizedLetters = append(normalizedLetters, r)
		}
	}

	textInSearchableFormat = strings.ToLower(string(normalizedLetters))

	return textInSearchableFormat
}
