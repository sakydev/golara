package golaStr

import (
	"fmt"
	"strings"
)

func PadBoth(str string, length int, pad string) string {
	return applyPad(str, length, pad, "both")
}

func PadLeft(str string, length int, pad string) string {
	return applyPad(str, length, pad, "left")
}

func PadRight(str string, length int, pad string) string {
	return applyPad(str, length, pad, "right")
}

func Reverse(str string) string {
	parts := []rune(str)
	length := len(parts)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	return string(parts)
}

func applyPad(str string, length int, pad string, side string) string {
	if len(str) >= length {
		return str
	}

	padding := strings.Repeat(pad, length-len(str))

	switch side {
	case "left":
		return fmt.Sprintf("%s%s", padding, str)
	case "right":
		return fmt.Sprintf("%s%s", str, padding)
	default:
		return fmt.Sprintf("%s%s%s", padding[:len(padding)/2], str, padding[len(padding)/2:])
	}
}
