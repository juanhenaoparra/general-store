package utils

import (
	"encoding/csv"
	"strings"
)

// GetNormalizeCsv gets a text and replace the anormal char or string with another one
func GetNormalizeCsv(text string, toreplace string, withthis string) [][]string {
	// Replace 'toreplace' string with 'withthis'
	replaced := strings.ReplaceAll(text, toreplace, withthis)

	// Create a csv reader for the replaced string
	reader := csv.NewReader(strings.NewReader(replaced))
	results, _ := reader.ReadAll()

	return results
}
