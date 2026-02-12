package trim

import "strings"

func FormatText(text string) string {
	// Convert the text to uppercase and trim spaces
	processedText := strings.TrimSpace(strings.ToUpper(text))
	return processedText
}

func Trim(text string) string {
	trimmedText := strings.TrimSpace(text)
	return trimmedText
}
