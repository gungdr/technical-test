package refactor

import "strings"

// findFirstStringInBracket returns string inside a bracket
func findFirstStringInBracket(word string) string {
	length := len(word)
	if length > 0 {
		openBracketIndex := strings.LastIndex(word, "(")
		closingBrackerIndex := strings.Index(word, ")")
		if openBracketIndex >= 0 && closingBrackerIndex >= 0 {
			// return the string between '(' and ')'
			return string(word[openBracketIndex+1 : closingBrackerIndex])
		}
	}
	return ""
}
