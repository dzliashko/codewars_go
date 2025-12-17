// Package kyu 7 with solved tasks.
package kyu7

import "strings"

// Complete the function that accepts a string parameter, and reverses each word in the string.
// All spaces in the string should be retained.

// Examples
// "This is an example!" ==> "sihT si na !elpmaxe"
// "double  spaces"      ==> "elbuod  secaps"

func ReverseWords(str string) string {
	strSlice := strings.Split(str, " ")
	result := []string{}

	for _, word := range strSlice {
		runes := []rune{}
		for i := len(word) - 1; i >= 0; i-- {
			runes = append(runes, rune(word[i]))
		}
		result = append(result, string(runes))
	}

	return strings.Join(result, " ")
}
