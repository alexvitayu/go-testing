package tempfiledemo_test

import (
	"errors"
	"os"
	"testing"

	"github.com/alexvitayu/go-testing/tempfiledemo"
)

func TestWriteLinesToTemp(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		want string
	}{
		{name: "singleWord", in: []string{"hello"}, want: "hello"},
		{name: "multiLines", in: []string{"go", "is", "reliable"}, want: "go\nis\nreliable"},
		{name: "emptyFile", in: nil, want: ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testTempFile, err := tempfiledemo.WriteLinesToTemp("example-*.txt", tc.in)
			assertError(t, err, "unexpected error")
			data, err := os.ReadFile(testTempFile)
			assertError(t, err, "fail reading file")
			assertString(t, data, tc.want)
			if err := os.Remove(testTempFile); err != nil {
				t.Fatalf("fail to remove: %vv", err)
			}
			if _, err := os.Stat(testTempFile); !errors.Is(err, os.ErrNotExist) {
				t.Fatalf("expected to be removed, but got error: %v", err)
			}

		})
	}
}

func assertString(t testing.TB, data []byte, want string) {
	t.Helper()
	if string(data) != want {
		t.Errorf("expected\n %v, got\n %v", want, string(data))
	}
}

func assertError(t testing.TB, err error, msg string) {
	t.Helper()
	if err != nil {
		t.Fatalf("%v: %v", msg, err)
	}
}
