// Package regex string handle string regex operations
package regex

import (
	"fmt"
	"regexp"
	"strings"
)

// Pattern is a struct that contains a Regex pattern, a validate message, and an opperation to be validated
type Pattern struct {
	Regex   string
	Message string
	Found   bool
}

// Patterns is a slice of RegexPattern
type Patterns []Pattern

// DefineRegexPatternsAndMessages gets a slice of strings with regex pattern,
// validate message and type of validation separated by ':' and returns a slice
// of RegexPattern to be used for validating strings patterns
func DefineRegexPatternsAndMessages(inputPatterns []string) Patterns {
	patterns := Patterns{}
	for _, pattern := range inputPatterns {
		patternMessage := strings.Split(pattern, ":")
		patterns = append(patterns, Pattern{Regex: patternMessage[0],
			Message: patternMessage[1],
			Found:   strings.ToLower(patternMessage[2]) == "t"})
	}
	return patterns
}

// AssertRegexPattern validate if pattern is present in string
func AssertRegexPattern(str string, pattern string) (bool, int64) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("error when trying to compile regex due to: %s", err.Error())
	}

	found := reg.FindAllString(str, -1)
	nFound := len(found)
	if nFound <= 0 {
		return false, int64(nFound)
	}
	return true, int64(nFound)
}
