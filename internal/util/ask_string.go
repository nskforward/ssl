package util

import (
	"fmt"
	"strings"
)

func AskString(prompt string, required, allowSpaces bool) string {
	fmt.Printf("- %s: ", prompt)
	var answer string
	fmt.Scanln(&answer)
	answer = strings.TrimSpace(answer)
	if !allowSpaces && strings.Contains(answer, " ") {
		fmt.Println("error: spaces not allowed")
		return AskString(prompt, required, allowSpaces)
	}
	if answer == "" && required {
		fmt.Println("error: empty value not allowed")
		return AskString(prompt, required, allowSpaces)
	}
	return answer
}
