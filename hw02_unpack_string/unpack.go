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
	strLen := len([]rune(str))
	for pos, r := range []rune(str) {
		isIndex := unicode.IsDigit(r) && !escaped
		currStr := string(r)
		prevStr := string(prevRune)

		switch {
		case !escaped && currStr == "\\":
			escaped = true
			isChar = false
			continue
		case escaped && !(unicode.IsDigit(r) || currStr == "\\"):
			return "", ErrInvalidString
		case escaped:
			prevRune = r
			escaped = false
			isChar = true
			strBuilder.WriteString(prevStr)
			if pos == strLen-1 {
				strBuilder.WriteString(currStr)
			}
			continue
		}

		if isIndex && (prevRune == 0 || (unicode.IsDigit(prevRune) && !isChar)) {
			return "", ErrInvalidString
		} else if isIndex {
			isChar = false
			count, err := strconv.Atoi(currStr)
			if err != nil {
				return "", err
			}
			strBuilder.WriteString(strings.Repeat(prevStr, count))
			prevRune = r
			continue
		}
		if !unicode.IsDigit(prevRune) && pos != 0 {
			strBuilder.WriteString(prevStr)
		}
		prevRune = r
		if pos == strLen-1 {
			strBuilder.WriteString(currStr)
		}
	}

	return strBuilder.String(), nil
}
