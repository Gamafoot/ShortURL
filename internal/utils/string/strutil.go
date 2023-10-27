package string

import "strings"

// Strip Убирает лишние пробелы в тексте.
func Strip(text string) string {
	array := strings.Fields(text)
	cleanedText := strings.Join(array, " ")
	return cleanedText
}
