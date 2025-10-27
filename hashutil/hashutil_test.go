package hashutil_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alexvitayu/go-testing/hashutil"
	"github.com/stretchr/testify/require"
)

func TestHashSHA256(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want string
	}{
		{name: "ascii", in: "hexlet's my hope", want: "661411f83ed71de148bcf282591abbcb5385d955d57d56b9d1e1a8c849b7384a"},
		{name: "unicode", in: "вариант №2", want: "28b0fd5a9309f50e29615d9671be5f8f45473fd93a8aa532ab5d2799ad8802b6"},
		{name: "emptyStr", in: "", want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}
	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := hashutil.HashSHA256(c.in)
			require.Len(t, got, 64, "len should be exactly 64 characters")
			require.Equal(t, c.want, got)
		})
	}
}

func TestDouble(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		want int
	}{
		{"positive", 5, 10},
		{"negative", -7, -14},
		{"zero", 0, 0},
	}
	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := hashutil.Double(c.in)
			if got != c.want {
				t.Errorf("expected %v but got %v", c.want, got)
			}
		})
	}
}

// TestIncrement показывает race condition, введите go test hashutil_test.go -race TestIncrement
func TestIncrement(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		t.Parallel()
		fmt.Println(hashutil.Increment(3))
	})
	t.Run("second", func(t *testing.T) {
		t.Parallel()
		fmt.Println(hashutil.Increment(2))
	})
	t.Run("third", func(t *testing.T) {
		t.Parallel()
		fmt.Println(hashutil.Increment(1))
	})
}

// TestIncrementSafe показывает, как выполняются тесты параллельно без гонки данных
func TestIncrementSafe(t *testing.T) {
	hashutil.CounterSafe = 0
	t.Run("first", func(t *testing.T) {
		t.Parallel()
		fmt.Println(hashutil.IncrementSafe(3))
	})
	t.Run("second", func(t *testing.T) {
		t.Parallel()
		fmt.Println(hashutil.IncrementSafe(2))
	})
	t.Run("third", func(t *testing.T) {
		t.Parallel()
		fmt.Println(hashutil.IncrementSafe(1))
	})
}

var counter int

func Increment() {
	counter++ // небезопасный доступ
}

var (
	mu          sync.Mutex
	counterSafe int
)

func IncrementSafe() {
	mu.Lock()     // берём «замок»
	counterSafe++ // изменяем ресурс
	mu.Unlock()   // отпускаем замок
}

func TestIncrement_Race(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10; i++ {
		t.Run("race", func(t *testing.T) {
			t.Parallel()
			Increment()
		})
	}
}

func TestIncrementSafe_NoRace(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10; i++ {
		t.Run("safe", func(t *testing.T) {
			t.Parallel()
			IncrementSafe()
		})
	}
}
