package main

import (
	"strings"
)

func cleanInput(text string) []string {
	s := []string{}
	trimmedString := strings.Fields(text)
	for _, word := range trimmedString {
		lowerStr := strings.ToLower(word)
		s = append(s, lowerStr)
	}
	return s
}
