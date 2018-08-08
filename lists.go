package pubtools

import (
	"strings"
)

// GetListUniqueString returns unique and not empty string list
func GetListUniqueString(elements []string) []string {
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

// GetListStringInSlice checks is string in slice
func GetListStringInSlice(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
