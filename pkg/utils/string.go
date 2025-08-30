package utils

import (
	"strings"

	"github.com/jinzhu/inflection"
)

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + word[1:]
}

func Pluralize(word string) string {
	return inflection.Plural(word)
}

func stringInSlice(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
