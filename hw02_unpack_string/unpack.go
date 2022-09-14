package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var prevRune rune
	strBuilder := strings.Builder{}
	isChar := true
	escaped := false
	for pos, r := range str {
		isIndex := unicode.IsDigit(r) && !escaped
		currStr := string(r)
		prevStr := string(prevRune)

		if !escaped && currStr == "\\" {
			escaped = true
			isChar = false
			continue
		}
		if escaped && !(unicode.IsDigit(r) || currStr == "\\") {
			return "", ErrInvalidString
		} else if escaped {
			prevRune = r
			escaped = false
			isChar = true
			strBuilder.WriteString(prevStr)
			if pos == len(str)-1 {
				strBuilder.WriteString(currStr)
			}
			continue
		}
		if isIndex && (prevRune == 0 || (unicode.IsDigit(prevRune) && !isChar)) {
			return "", ErrInvalidString
		}
		if isIndex {
			isChar = false
			count, _ := strconv.Atoi(currStr)
			strBuilder.WriteString(strings.Repeat(prevStr, count))
			prevRune = r
			continue
		}
		if !unicode.IsDigit(prevRune) && pos != 0 {
			strBuilder.WriteString(prevStr)
		}
		prevRune = r
		if pos == len(str)-1 {
			strBuilder.WriteString(currStr)
		}
	}

	return strBuilder.String(), nil
}
