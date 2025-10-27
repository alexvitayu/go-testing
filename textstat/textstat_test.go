package textstat_test

import (
	"testing"

	"github.com/alexvitayu/go-testing/textstat"
	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want map[string]int
	}{
		{name: "empty", in: "", want: map[string]int{}},
		{name: "repeated", in: "word WORD WoRd", want: map[string]int{"word": 3}},
		{name: "punctuation", in: " , ! ", want: map[string]int{}},
		{name: "mixed", in: "Deal, it is time! it is opportunity.", want: map[string]int{"deal": 1, "it": 2, "is": 2, "time": 1, "opportunity": 1}},
	}
	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := textstat.WordCount(c.in)
			assert.Equal(t, c.want, got)
		})
	}
}
