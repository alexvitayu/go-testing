package normalize_test

import (
	"testing"

	"github.com/alexvitayu/go-testing/normalize"
)

func TestClean(t *testing.T) {
	var cases = []struct {
		name string
		s    string
		want string
	}{
		{
			name: "emptyString",
			s:    "",
			want: "",
		},
		{
			name: "severalSpaces&Tabs",
			s:    "go  is 		the best\t    from 	languages\t",
			want: "go is the best from languages",
		},
		{
			name: "nonCapital&CapitalLetters",
			s:    "Nature   AsKs For HelP",
			want: "nature asks for help",
		},
		{
			name: "nonLettersSymbols",
			s:    "hello everyone. is  	there something\t new?",
			want: "hello everyone. is there something new?",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := normalize.Clean(tc.s)
			if got != tc.want {
				t.Errorf("Ожидали %v, получили %v", tc.want, got)
			}
		})
	}
}
