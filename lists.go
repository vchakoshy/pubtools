package pubtools

import (
	"strings"
)

// ListLib struct
type ListLib struct{}

// List returns ListLib
func List() *ListLib {
	return &ListLib{}
}

// UniqueListOfString returns unique and not empty string list
func (l *ListLib) UniqueListOfString(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			encountered[elements[v]] = true
			if len(strings.TrimSpace(elements[v])) > 0 {
				result = append(result, elements[v])
			}
		}
	}
	return result
}

// IsStringInSlice checks is string in slice
func (l *ListLib) IsStringInSlice(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
