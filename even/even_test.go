package even

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Errorf("Ожидалось true для 2, но получили %v", IsEven(2))
	}
	if IsEven(3) {
		t.Errorf("Ожидалось false для 3, но получили %v", IsEven(3))
	}
	if !IsEven(0) {
		t.Errorf("Ожидалось true для 0, но получили %v", IsEven(0))
	}
	if IsEven(-3) {
		t.Errorf("Ожидалось false для -3, но получили %v", IsEven(-3))
	}
	if !IsEven(-4) {
		t.Errorf("Ожидалось true для -4, но получили %v", IsEven(-4))
	}

}
