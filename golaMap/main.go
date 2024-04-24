package golaMap

import "strings"

func Flatten[T comparable](items map[T]T) []T {
	var result []T

	for key, value := range items {
		result = append(result, key, value)
	}

	return result
}

func Exists[T comparable, X comparable](items map[T]X, searchKey T) bool {
	for key, _ := range items {
		if key == searchKey {
			return true
		}
	}

	return false
}

func Only[T comparable, X comparable](items map[T]X, searchKey T) map[T]X {
	result := make(map[T]X)

	for key, value := range items {
		if key == searchKey {
			result[key] = value

			return result
		}
	}

	return nil
}

func Except[T comparable, X comparable](items map[T]X, searchKey T) map[T]X {
	result := make(map[T]X)

	for key, value := range items {
		if key == searchKey {
			continue
		}

		result[key] = value
	}

	return result
}

func RemoveEmptyElementsStr[T comparable](items map[T]string) {
	for key, value := range items {
		if strings.TrimSpace(value) == "" {
			delete(items, key)
		}
	}
}

func RemoveEmptyElementsInt[T comparable](items map[T]int) {
	for key, value := range items {
		if value == 0 {
			delete(items, key)
		}
	}
}

func SearchKeyByValue(items map[string]string, search string) string {
	for key, value := range items {
		if strings.TrimSpace(value) == search {
			return key
		}
	}

	return ""
}
