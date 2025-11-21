package helper

import (
	"strconv"
	"strings"
)



func TransformWord(word, operation string) string {
	switch operation {
	case "hex":
		if num, err := strconv.ParseInt(word, 16, 64); err == nil {
			return strconv.FormatInt(num, 10)
		}
	case "bin":
		if num, err := strconv.ParseInt(word, 2, 64); err == nil {
			return strconv.FormatInt(num, 10)
		}
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		if len(word) > 0 {
			return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return word
}