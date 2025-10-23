package safe_test

import (
	"testing"

	"github.com/alexvitayu/go-testing/safe"
)

func TestMustAt_Valid(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	index := 1
	want := 2
	got := safe.MustAt(nums, index)
	if got != want {
		t.Errorf("Ожидалось %d, получили %d", want, got)
	}
}

func TestMustAt_Panic_OutOfRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, got none")
		} else if r != "index out of range" {
			t.Errorf("unexpected panic message %v", r)
		}
	}()
	nums := []int{1, 2, 3, 4, 5}
	index := 10
	_ = safe.MustAt(nums, index)
}

func TestMustAt_Panic_NegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, got none")
		} else if r != "index out of range" {
			t.Errorf("unexpected panic message %v", r)
		}
	}()
	words := []string{"go", "cat", "the"}
	index := -1
	_ = safe.MustAt(words, index)
}
