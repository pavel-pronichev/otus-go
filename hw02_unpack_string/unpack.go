package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var b strings.Builder

	var current rune

	for _, r := range s {
		if num, err := strconv.Atoi(string(r)); err == nil {
			if current == 0 {
				return "", ErrInvalidString
			}

			b.WriteString(strings.Repeat(string(current), num))
			current = 0
		} else {
			if current != 0 {
				b.WriteRune(current)
			}

			current = r
		}
	}

	if current != 0 {
		b.WriteRune(current)
	}

	return b.String(), nil
}
