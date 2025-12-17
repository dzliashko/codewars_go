// Package kyu 7 with solved tasks.
package kyu7

// import "strings"

// DisemvowelTrolls removes vowels from a string.
// func Disemvowel(comment string) string {
// 	vowels := "aeiouAEIOU"
// 	result := []string{}
// 	for _, word := range strings.Split(comment, " ") {
// 		for _, c := range vowels {
// 			word = strings.ReplaceAll(word, string(c), "")
// 		}
// 		result = append(result, word)
// 	}
// 	return strings.Join(result, " ")
// }

// Best practice
// import "regexp"

// func Disemvowel(comment string) string {
// 	return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
// }

// Another best practice
import "strings"

func Disemvowel(comment string) string {
	for _, c := range "aiueoAIUEO" {
		comment = strings.ReplaceAll(comment, string(c), "")
	}
	return comment
}
