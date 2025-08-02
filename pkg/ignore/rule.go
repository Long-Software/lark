package ignore

import (
	"regexp"
	"strings"
)

type Rule struct {
	Pattern   string
	Regex     *regexp.Regexp
	IsNegated bool
	IsDir     bool
}

func NewRule(pattern string) Rule {
	isNegated := false
	if strings.HasPrefix(pattern, "!") {
		isNegated = true
		pattern = strings.TrimPrefix(pattern, "!")
	}

	isDir := strings.HasSuffix(pattern, "/")

	regexPattern := globToRegex(pattern)
	regex := regexp.MustCompile(regexPattern)

	return Rule{
		Pattern:   pattern,
		Regex:     regex,
		IsNegated: isNegated,
		IsDir:     isDir,
	}
}

func globToRegex(pattern string) string {
	regex := regexp.QuoteMeta(pattern)
	regex = strings.ReplaceAll(regex, `\*`, ".*")
	return "^" + regex + "$"
}