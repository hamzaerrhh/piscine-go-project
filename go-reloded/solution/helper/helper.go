package helper

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"strings"
	"unicode"
)



type commandInfo struct {
	name  string
	count int
}


func IsPunctuation(r rune) bool {
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

func IsWordRune(r rune) bool {
	// Allow letters, digits, and most Unicode symbols (including emojis)
	// but exclude specific punctuation and control characters
	return unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSymbol(r) || 
		   unicode.IsPunct(r) && r != '(' && r != ')' && r != '\'' && 
		   r != '.' && r != '!' && r != '?' && r != ',' && r != ':' && r != ';'
}
func handleCommands(tokens []Token) []Token {
	var result []Token
	
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == COMMAND {
			if len(tokens[i].Children) > 0 {
				processedChildren := handleCommands(tokens[i].Children)
				tokens[i].Children = processedChildren
			}
			
			command := extractCommand(tokens[i].Children)
			
			switch command.name {
			case "hex", "bin", "up", "low", "cap":
				// Find how many words to process
				count := 1
				if command.count > 0 {
					count = command.count
				}
				
				// Check if we have enough words before the command
				if len(result) < count {
					result = append(result, tokens[i])
					continue
				}
				
				// Process the words (from right to left)
				processed := 0
				for j := len(result) - 1; j >= 0 && processed < count; j-- {
					if result[j].Type == WORD {
						result[j].Value = TransformWord(result[j].Value, command.name)
						processed++
					}
				}
				
				// Remove the command token since we've processed it
			default:
				result = append(result, tokens[i])
			}
		} else {
			result = append(result, tokens[i])
		}
	}
	
	return result
}

func extractCommand(children []Token) commandInfo {
	if len(children) == 0 {
		return commandInfo{}
	}
	
	// First, process any nested commands in children
	var processedChildren []Token
	for _, child := range children {
		if child.Type == COMMAND {
			// Process nested command
			processedNested := handleCommands([]Token{child})
			if len(processedNested) > 0 {
				processedChildren = append(processedChildren, processedNested[0])
			}
		} else {
			processedChildren = append(processedChildren, child)
		}
	}
	
	// Find the command name (first word token)
	var nameToken *Token
	for i := range processedChildren {
		if processedChildren[i].Type == WORD {
			nameToken = &processedChildren[i]
			break
		}
	}
	
	if nameToken == nil {
		return commandInfo{}
	}
	
	cmd := commandInfo{name: strings.ToLower(nameToken.Value)}
	
	// Check for count parameter (word after comma)
	for i := 0; i < len(processedChildren)-2; i++ {
		if processedChildren[i].Type == PUNCT && processedChildren[i].Value == "," {
			// Look for number in next tokens
			for j := i + 1; j < len(processedChildren); j++ {
				if processedChildren[j].Type == WORD {
					// Try to convert to number, handle nested command results
					if count, err := strconv.Atoi(processedChildren[j].Value); err == nil {
						cmd.count = count
						break
					}
				}
			}
			break
		}
	}
	
	return cmd
}
func ProtectLayer(filePath,output string){
maxSize:= int64(200 * 1024 * 1024)
info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error checking file:", err)
		return
	}

	if info.Size() > maxSize {
		fmt.Println("Error: file is larger than 200MB")
		os.Exit(0)
	}
	// REGEX: match anything ending with main.go
	blockMain := regexp.MustCompile(`.go$`)

	if blockMain.MatchString(output) || blockMain.MatchString(filePath) {
		fmt.Println("Error: output file cannot be main.go")
		os.Exit(0)
	}


	

}
