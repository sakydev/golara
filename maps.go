package golara

import "strings"

func RemoveEmptyElements(items map[string]string) {
	for key, value := range items {
		if strings.Trim(value, " ") == "" {
			delete(items, key)
		}
	}
}

func SearchKeyByValue(items map[string]string, search string) string {
	for key, value := range items {
		if strings.Trim(value, "") == search {
			return key
		}
	}

	return ""
}
