package helper

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	WORD TokenType = iota
	COMMAND
	PUNCT
	QUOTE
)

type Token struct {
	Type     TokenType
	Value    string
	Warning  string
	Children []Token
}

// -----------------------------------------------------------------------------
// PUBLIC ENTRY POINT
// -----------------------------------------------------------------------------

func Tokenize(text string) []Token {
	tokens, _ := tokenizeRunes([]rune(text), 0)
	PrintTokens(tokens, 0)
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

		case isWordRune(c):
			start := i
			for i < len(runes) && isWordRune(runes[i]) {
				i++
			}
			tokens = append(tokens, Token{Type: WORD, Value: string(runes[start:i])})

		case c == '\'':
			tokens = append(tokens, Token{Type: QUOTE, Value: "'"})
			i++

		case isPunctuation(c):
			tokens = append(tokens, Token{Type: PUNCT, Value: string(c)})
			i++

		default:
			i++
		}
	}

	return tokens, i
}

// -----------------------------------------------------------------------------
// COMMAND READER
// -----------------------------------------------------------------------------

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

func isWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\n' || unicode.IsSymbol(r)
}

func isPunctuation(r rune) bool {
	switch r {
	case '.', '!', '?', ',', ':', ';':
		return true
	}
	return false
}

func hasTopLevelUnexpectedSpaces(inner string) bool {
	depth := 0
	prevSpace := false

	for _, r := range inner {
		switch r {
		case '(':
			depth++
			prevSpace = false
		case ')':
			depth--
			prevSpace = false
		default:
			if unicode.IsSpace(r) {
				if depth == 0 {
					if prevSpace {
						return true
					}
					prevSpace = true
				}
			} else {
				prevSpace = false
			}
		}
	}
	return false
}

func PrintTokens(tokens []Token, indent int) {
	prefix := func() string {
		s := ""
		for i := 0; i < indent; i++ {
			s += "  "
		}
		return s
	}()

	for i, t := range tokens {
		name := [...]string{"WORD", "COMMAND", "PUNCT", "QUOTE"}[t.Type]

		if t.Warning != "" {
			fmt.Printf("%s[%d] %s: '%s'  ⚠ %s\n", prefix, i, name, t.Value, t.Warning)
		} else {
			fmt.Printf("%s[%d] %s: '%s'\n", prefix, i, name, t.Value)
		}

		if len(t.Children) > 0 {
			PrintTokens(t.Children, indent+1)
		}
	}
}
