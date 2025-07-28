package git

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type GitignoreRule struct {
	Pattern   string
	Regex     *regexp.Regexp
	IsNegated bool
	IsDir     bool
}

type GitignoreParser struct {
	FilePath string
}

func (g *GitignoreRule) ShouldIgnore(path string) bool {
	if g.Regex.MatchString(path) {
		if g.IsNegated {
			return false
		} else {
			return true
		}

	}
	return false
}

func (g *GitignoreParser) Parse() []GitignoreRule {
	var rules = []GitignoreRule{}
	file, err := os.Open(g.FilePath)
	if err != nil {
		return rules
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		rule := g.parsRule(line)
		if rule != nil {
			rules = append(rules, *rule)
		}
	}
	return rules
}

func (g *GitignoreParser) parsRule(pattern string) *GitignoreRule {
	rule := &GitignoreRule{
		Pattern: pattern,
	}

	// Handle negation (!)
	if strings.HasPrefix(pattern, "!") {
		rule.IsNegated = true
		pattern = pattern[1:]
	}

	// Handle directory patterns (/)
	if strings.HasSuffix(pattern, "/") {
		rule.IsDir = true
		pattern = strings.TrimSuffix(pattern, "/")
	}

	// Convert gitignore pattern to regex
	regexPattern := g.parsePattern(pattern)
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil
	}

	rule.Regex = regex
	return rule
}

func (g *GitignoreParser) parsePattern(pattern string) string {
	// Escape special regex characters except * and ?
	pattern = regexp.QuoteMeta(pattern)

	// Handle gitignore wildcards
	pattern = strings.ReplaceAll(pattern, "\\*\\*", ".*") // ** matches any path
	pattern = strings.ReplaceAll(pattern, "\\*", "[^/]*") // * matches anything except /
	pattern = strings.ReplaceAll(pattern, "\\?", "[^/]")  // ? matches single character except /

	// Handle leading slash (absolute path)
	if strings.HasPrefix(pattern, "/") {
		pattern = "^" + pattern[1:]
	} else {
		pattern = "(^|/)" + pattern
	}

	// Handle trailing patterns
	pattern = pattern + "(/.*)?$"

	return pattern
}
