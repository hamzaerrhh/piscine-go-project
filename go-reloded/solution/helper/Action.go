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
func HandlePunctuation(tokens []Token) []Token {
	var result []Token
	
	for i := 0; i < len(tokens); i++ {
		current := tokens[i]
		
		if current.Type == PUNCT {
			// If there's a previous word, attach punctuation to it
			if len(result) > 0 && result[len(result)-1].Type == WORD {
				result[len(result)-1].Value += current.Value
			} else {
				result = append(result, current)
			}
		} else if current.Type == SPACE {
			// Keep all spaces including newlines
			if current.Value == "\n" {
				// Always keep newlines
				result = append(result, current)
			} else {
				// For regular spaces, only keep if it's between words
				if len(result) > 0 {
					lastToken := result[len(result)-1]
					if lastToken.Type == WORD || lastToken.Type == QUOTE {
						// Add space only if next token is a word
						if i+1 < len(tokens) && (tokens[i+1].Type == WORD || tokens[i+1].Type == QUOTE) {
							result = append(result, current)
						}
					} else {
						result = append(result, current)
					}
				}
			}
		} else {
			result = append(result, current)
		}
	}
	
	return result
}
func HandleQuotes(tokens []Token) []Token {
	var result []Token
	quoteCount := 0
	
	for i := 0; i < len(tokens); i++ {
		current := tokens[i]
		
		if current.Type == QUOTE {
			quoteCount++
			
			if quoteCount%2 == 1 {
				// Opening quote - don't add space before it
				if len(result) > 0 && result[len(result)-1].Type == SPACE {
					result = result[:len(result)-1]
				}
				result = append(result, current)
			} else {
				// Closing quote - attach to previous word if exists
				if len(result) > 0 && result[len(result)-1].Type == WORD {
					result[len(result)-1].Value += current.Value
				} else {
					result = append(result, current)
				}
			}
		} else if current.Type == PUNCT {
			// Handle punctuation - don't let it interfere with quote spacing
			result = append(result, current)
		} else if current.Type == SPACE && quoteCount%2 == 1 {
			// Inside quotes - keep the space
			result = append(result, current)
		} else if current.Type == WORD {
			if quoteCount%2 == 1 && len(result) > 0 && result[len(result)-1].Type == QUOTE {
				// Word after opening quote - attach quote to word
				current.Value = result[len(result)-1].Value + current.Value
				result = result[:len(result)-1]
			}
			result = append(result, current)
		} else {
			result = append(result, current)
		}
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