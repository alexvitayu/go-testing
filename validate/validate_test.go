package validate

import "testing"

func TestValidateName_NonEmpty(t *testing.T) {
	err := ValidateName("Andy")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
}

func TestValidateName_Empty(t *testing.T) {
	err := ValidateName("")
	if err == nil {
		t.Fatalf("expected error, got nil")
	} else if err.Error() != ErrEmptyName.Error() {
		t.Errorf("unexpected error, %v", err)
	}
}
