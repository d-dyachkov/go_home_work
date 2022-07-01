package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const (
	emptyToken rune = 0
	slashToken rune = '\\'
)

func Unpack(in string) (string, error) {
	escaped := false
	prevToken := emptyToken
	builder := &strings.Builder{}
	for _, token := range in {
		count := 1
		switch {
		case token == slashToken:
			if escaped = !escaped; escaped {
				token = emptyToken
			}
		case escaped && unicode.IsDigit(token):
			prevToken = emptyToken
			escaped = false
		case !escaped && unicode.IsDigit(token):
			if prevToken == emptyToken {
				return "", ErrInvalidString
			}
			count, _ = strconv.Atoi(string(token))
			token = emptyToken
		case escaped:
			return "", ErrInvalidString
		}
		if err := writeUnpacked(builder, prevToken, count); err != nil {
			return "", err
		}
		prevToken = token
	}
	err := writeUnpacked(builder, prevToken, 1)
	return builder.String(), err
}

func writeUnpacked(b *strings.Builder, r rune, count int) error {
	var err error
	if r != emptyToken {
		if count != 1 {
			unpacked := strings.Repeat(string(r), count)
			_, err = b.WriteString(unpacked)
		} else {
			_, err = b.WriteRune(r)
		}
	}
	return err
}
