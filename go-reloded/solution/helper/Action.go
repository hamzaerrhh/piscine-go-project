package helper

import (
	"strconv"
	"strings"
	"unicode"
)

func HandleAAn(tokens []Token) []Token {
	var result []Token

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == WORD && strings.ToLower(tokens[i].Value) == "a" {
			j := i + 1
			for j < len(tokens) && tokens[j].Type != WORD {
				j++
			}

			if j < len(tokens) && tokens[j].Type == WORD {
				nextWord := tokens[j].Value
				if len(nextWord) > 0 {
					firstChar := unicode.ToLower(rune(nextWord[0]))
					if firstChar == 'a' || firstChar == 'e' || firstChar == 'i' ||
						firstChar == 'o' || firstChar == 'u' || firstChar == 'h' {
						result = append(result, Token{
							Type:  WORD,
							Value: "an",
						})
						continue
					}
				}
			}
		}
		result = append(result, tokens[i])
	}

	return result
}

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