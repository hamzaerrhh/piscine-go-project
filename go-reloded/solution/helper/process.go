package helper

import (
	"strings"
	"unicode"
)


type TokenType int

const (
	WORD TokenType = iota
	COMMAND
	PUNCT
	QUOTE
	SPACE
)

type Token struct {
	Type     TokenType
	Value    string
	Warning  string
	Children []Token
}

func Tokenize(text string) []Token {
	tokens, _ := tokenizeRunes([]rune(text), 0)
	return tokens
}
func tokenizeRunes(runes []rune, i int) ([]Token, int) {
	var tokens []Token

	for i < len(runes) {
		c := runes[i]

		switch {
		case c == '(':
			tok, next := readCommand(runes, i)
			tokens = append(tokens, tok)
			i = next

		case c == '\n':
			// Handle newline as a separate token
			tokens = append(tokens, Token{Type: SPACE, Value: "\n"})
			i++

		case IsWordRune(c):
			start := i
			for i < len(runes) && IsWordRune(runes[i]) {
				i++
			}
			tokens = append(tokens, Token{Type: WORD, Value: string(runes[start:i])})

		case c == '\'':
			tokens = append(tokens, Token{Type: QUOTE, Value: "'"})
			i++

		case IsPunctuation(c):
			// Handle punctuation groups like "...", "!!", "?!"
			start := i
			for i < len(runes) && IsPunctuation(runes[i]) {
				i++
			}
			tokens = append(tokens, Token{Type: PUNCT, Value: string(runes[start:i])})

		case unicode.IsSpace(c) && c != '\n': // Handle spaces but not newlines
			start := i
			for i < len(runes) && unicode.IsSpace(runes[i]) && runes[i] != '\n' {
				i++
			}
			tokens = append(tokens, Token{Type: SPACE, Value: string(runes[start:i])})

		default:
			// Handle any other characters (including emojis that might not be caught above)
			tokens = append(tokens, Token{Type: WORD, Value: string([]rune{c})})
			i++
		}
	}

	return tokens, i
}

func readCommand(runes []rune, start int) (Token, int) {
	i := start + 1
	depth := 1

	// find matching parentheses
	for i < len(runes) && depth > 0 {
		if runes[i] == '(' {
			depth++
		} else if runes[i] == ')' {
			depth--
		}
		i++
	}

	raw := string(runes[start:i])
	inner := raw[1 : len(raw)-1]

	children, _ := tokenizeRunes([]rune(inner), 0)

	warning := ""
	if hasTopLevelUnexpectedSpaces(inner) {
		warning = "Unexpected spaces inside command"
	}

	return Token{
		Type:     COMMAND,
		Value:    raw,
		Warning:  warning,
		Children: children,
	}, i
}


//etaps 2

func ProcessTokens(tokens []Token) []Token {
	tokens = handleCommands(tokens)
	tokens = HandleAAn(tokens)
	tokens = HandlePunctuation(tokens)
	tokens = HandleQuotes(tokens)
	return tokens
}

//etaps 3
func TokensToString(tokens []Token) string {
	var result strings.Builder
	
	for i, token := range tokens {
		switch token.Type {
		case WORD:
			result.WriteString(token.Value)
		case PUNCT:
			result.WriteString(token.Value)
		case QUOTE:
			result.WriteString(token.Value)
		case SPACE:
			if token.Value == "\n" {
				result.WriteString("\n")
			} else {
				result.WriteString(" ")
			}
		case COMMAND:
			// Commands are removed from final output
		}
		
		// Improved spacing logic
		if i < len(tokens)-1 {
			currentToken := token
			nextToken := tokens[i+1]
			
			// Add space after words unless next is punctuation or we're at newline
			if currentToken.Type == WORD {
				if nextToken.Type == SPACE && nextToken.Value == "\n" {
					// Don't add space before newline
					continue
				} else if nextToken.Type == WORD || nextToken.Type == QUOTE {
					result.WriteString(" ")
				}
			}else if currentToken.Type == PUNCT {
				if nextToken.Type == WORD || nextToken.Type == QUOTE {
					result.WriteString(" ")
				}
			}			else if currentToken.Type == QUOTE {
				// Count quotes to determine if this is opening or closing
				quoteCount := 0
				for j := 0; j <= i; j++ {
					if tokens[j].Type == QUOTE {
						quoteCount++
					}
				}
				if quoteCount%2 == 1 && nextToken.Type == WORD { // Opening quote
					// No space - quote attaches to word
				} else if nextToken.Type == WORD {
					result.WriteString(" ")
				}
			}
		}
	}
	
	return result.String()
}
//etaps 4
func CleanText(text string) string {
	// Fix punctuation spacing but preserve newlines
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		// Apply cleanup to each line individually
		// Fix spaces before punctuation (but preserve intended spaces)
		line = strings.ReplaceAll(line, " .", ".")
		line = strings.ReplaceAll(line, " ,", ",")
		line = strings.ReplaceAll(line, " !", "!")
		line = strings.ReplaceAll(line, " ?", "?")
		line = strings.ReplaceAll(line, " :", ":")
		line = strings.ReplaceAll(line, " ;", ";")

		// Fix spaces around quotes - be more careful
		line = strings.ReplaceAll(line, " ' ", "' ")
		line = strings.ReplaceAll(line, "' ", "'")
		line = strings.ReplaceAll(line, " '", "'")

		// Ensure space after colon when followed by quote
		line = strings.ReplaceAll(line, ":'", ": '")

		// Clean up multiple spaces (but preserve single spaces)
		for strings.Contains(line, "  ") {
			line = strings.ReplaceAll(line, "  ", " ")
		}

		lines[i] = strings.TrimSpace(line)
	}

	// Join lines back with newlines
	return strings.Join(lines, "\n")
}