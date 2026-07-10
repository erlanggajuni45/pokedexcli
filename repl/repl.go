package repl

import "strings"

func CleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")
}
