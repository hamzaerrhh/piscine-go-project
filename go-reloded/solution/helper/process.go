package helper

import (
	"regexp"
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

	
		}
	}

	return tokens, i
}

func readCommand(runes []rune, start int) (Token, int) {
	i := start + 1
	depth := 1

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



func ProcessTokens(tokens []Token) []Token {
	tokens = handleCommands(tokens)

	return tokens
}

func TokensToString(tokens []Token) string {
	var result strings.Builder
	
	for _, token := range tokens {
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

	
		}
		
	
	}
	
	return result.String()
}
func CleanUpText(input string) string {
	if input == "" {
		return input
	}

	// Split input into lines
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		// Remove extra spaces
		for strings.Contains(line, "  ") {
			line = strings.ReplaceAll(line, "  ", " ")
		}
		line = strings.TrimSpace(line)

	

		// Fix punctuation
		line = FixPunctuation(line)
			// Handle quotes
		reQuotes := regexp.MustCompile(`'\s*([^']*?)\s*'`)
		line = reQuotes.ReplaceAllStringFunc(line, func(match string) string {
			start := strings.Index(match, "'") + 1
			end := strings.LastIndex(match, "'")
			content := match[start:end]
			content = strings.TrimSpace(content) // keep internal spaces
			return "'" + content + "'"
		})

		// Handle 'a' → 'an'
		reAtoAn := regexp.MustCompile(`\b([Aa])\s+([aeiouhAEIOUH]\w*)`)
		line = reAtoAn.ReplaceAllString(line, "an $2")

		lines[i] = line
	}

	// Join lines back preserving newlines
	return strings.Join(lines, "\n")
}

func FixPunctuation(input string) string {
	if input == "" {
		return input
	}

	// Merge consecutive punctuation groups
	reMerge := regexp.MustCompile(`([.,!?;:]+)\s+([.,!?;:]+)`)
	input = reMerge.ReplaceAllString(input, "$1$2")

	// Remove space before punctuation
	reBefore := regexp.MustCompile(`\s+([.,!?;:]+)`)
	input = reBefore.ReplaceAllString(input, "$1")

	// Ensure exactly one space after punctuation if not end of string
	reAfter := regexp.MustCompile(`([.,!?;:]+)([^\s.,!?;:]|$)`)
	input = reAfter.ReplaceAllString(input, "$1 $2")

	// Collapse multiple spaces, but do not remove newlines (already handled line by line)
	input = strings.Join(strings.Fields(input), " ")

	return input
}
