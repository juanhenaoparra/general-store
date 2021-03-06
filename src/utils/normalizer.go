package utils

import (
	"encoding/csv"
	"fmt"
	"strconv"
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

// RemoveDuplicateValues return cleaned array
func RemoveDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// CheckBack Looks back and return a bool if finds an item
func CheckBack(stringSlice []string, searchItem string) bool {
	for _, v := range stringSlice {
		if v == searchItem {
			return true
		}
	}

	return false
}

// GetTopAndLowStr given a media string return the top and low in string
func GetTopAndLowStr(media string, variance int) (string, string) {
	mediaInt, err := strconv.Atoi(media)

	if err != nil {
		fmt.Println(err)
	}
	margin := (mediaInt * variance) / 100

	topInt := mediaInt + margin
	lowInt := mediaInt - margin

	top := strconv.Itoa(topInt)
	low := strconv.Itoa(lowInt)

	return top, low
}
