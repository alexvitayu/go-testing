package greet_test

import (
	"errors"
	"testing"

	"github.com/alexvitayu/go-testing/greet"
)

//Базовые случаи: "Go", "World".
//Пустая строка: ожидаем ошибку (want = true).
//Юникод: например, "Гофер" возвращает "Hello, Гофер".
//Сохранение пробелов: вход " Go " возвращает "Hello, Go " (два пробела подряд допустимы, т.к. функция не триммит).

//assertString(t testing.TB, got, want string)
//assertError(t testing.TB, got error, want bool)

var ErrEmptyStr = errors.New("name cannot be empty")

func TestHello(t *testing.T) {
	t.Run("basicCases", func(t *testing.T) {
		got, err := greet.Hello("Go")
		assertString(t, got, "Hello, Go")
		assertError(t, err, false)

		got, err = greet.Hello("World")
		assertString(t, got, "Hello, World")
		assertError(t, err, false)

	})
	t.Run("emptyString", func(t *testing.T) {
		got, err := greet.Hello("")
		assertString(t, got, "")
		assertError(t, err, true)
	})
	t.Run("unicodeCase", func(t *testing.T) {
		got, err := greet.Hello("Гофер")
		assertString(t, got, "Hello, Гофер")
		assertError(t, err, false)
	})
	t.Run("withSpaces", func(t *testing.T) {
		got, err := greet.Hello(" Go ")
		assertString(t, got, "Hello,  Go ")
		assertError(t, err, false)
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("ожидали %v, получили %v", want, got)
	}
}

func assertError(t testing.TB, got error, want bool) {
	t.Helper()
	if got != nil && want == false {
		t.Fatalf("unexpected error: %v", got)
	}
	if got == nil && want == true {
		t.Fatalf("got nil instead of error")
	}
	if want && got.Error() != ErrEmptyStr.Error() {
		t.Errorf("unexpected error: %v", got)
	}
}
