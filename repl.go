package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	s := []string{}
	trimmedString := strings.Fields(text)
	for _, word := range trimmedString {
		lowerStr := strings.ToLower(word)
		s = append(s, lowerStr)
	}
	fmt.Printf("%q\n", s)
	return s
}
