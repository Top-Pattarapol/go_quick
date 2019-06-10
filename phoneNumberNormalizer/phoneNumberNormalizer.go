package main

import "strings"

// Normalizer phone number
func Normalizer(input []string) map[string]int {

	result := map[string]int{}

	for _, item := range input {
		item = strings.ReplaceAll(item, " ", "")
		item = strings.ReplaceAll(item, "-", "")
		item = strings.ReplaceAll(item, "(", "")
		item = strings.ReplaceAll(item, ")", "")
		result[item] = result[item] + 1
	}

	return result
}
