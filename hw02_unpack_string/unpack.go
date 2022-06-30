package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const emptyToken = rune(0)

func Unpack(in string) (string, error) {
	var err error
	token := emptyToken
	builder := strings.Builder{}
	for _, r := range []rune(in) {
		if !unicode.IsDigit(r) {
			if token != emptyToken {
				builder.WriteRune(token)
			}
			token = r
			continue
		}
		if token == emptyToken {
			err = ErrInvalidString
			break
		}
		unpacked := strings.Repeat(string(token), int(r-'0'))
		_, err = builder.WriteString(unpacked)
		token = emptyToken
	}
	if token != emptyToken {
		builder.WriteRune(token)
	}
	return builder.String(), err
}
