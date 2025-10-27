package textstat

import (
	"strings"
	"unicode"
)

// WordCount считает частоты слов (буквы/цифры), регистр игнорируется.
func WordCount(s string) map[string]int {
	toKey := func(r rune) bool {
		return !unicode.IsDigit(r) && !unicode.IsLetter(r)
	}
	tokens := strings.FieldsFunc(s, toKey)
	freq := make(map[string]int, len(tokens))
	for _, v := range tokens {
		k := strings.ToLower(v)
		if k == "" {
			continue
		}
		freq[k]++
	}
	return freq
}
