package api

import (
    "regexp"
    "unicode/utf8"
)

// IsValidUsername checks if the username meets the API requirements.
// It must be 3-20 characters long and match the provided regular expression pattern.
func IsValidUsername(username string) bool {
    const (
        minLen = 3  // Minimum length requirement for the username.
        maxLen = 20 // Maximum length requirement for the username.
    )
    pattern := `^.*?$` // Regular expression pattern for validating the username.

    // Check the length of the username.
    // utf8.RuneCountInString is used instead of len to correctly count multi-byte characters.
    if len := utf8.RuneCountInString(username); len < minLen || len > maxLen {
        return false // Return false if the username does not meet the length requirements.
    }

    // Check if the username matches the regular expression pattern.
    // regexp.MatchString returns true if the string matches the pattern.
    matched, err := regexp.MatchString(pattern, username)
    if err != nil || !matched {
        return false // Return false if there is an error or if the pattern does not match.
    }

    return true // Return true if the username passes both the length and pattern checks.
}
