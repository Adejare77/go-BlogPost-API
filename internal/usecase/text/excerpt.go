package text

import "strings"


func ToExcerpt(content string) string {
	words := strings.Fields(content)

	if len(words) > 30 {
		words = append(words[:30], " ...")
	}

	return strings.Join(words, " ")
}
