package api

import (
	
    "regexp"
    "unicode/utf8"
    

)

// IsValidUsername checks if the username meets the API requirements.
// It must be 3-20 characters long and match the provided pattern.
func IsValidUsername(username string) bool {
    const (
        minLen = 3
        maxLen = 20
    )
    pattern := `^.*?$` 

    // Check the length
    if len := utf8.RuneCountInString(username); len < minLen || len > maxLen {
        return false
    }

    // Check the pattern
    matched, err := regexp.MatchString(pattern, username)
    if err != nil || !matched {
        return false
    }

    return true
}
