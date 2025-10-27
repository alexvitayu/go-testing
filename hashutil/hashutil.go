package hashutil

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

// HashSHA256 возвращает hex‑представление SHA‑256 для строки s.
func HashSHA256(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// Double функция, которая удваивает число(простое итоговое задание)
func Double(n int) int {
	return n * 2
}

// Increment увеличивает глобальную переменную (вариант гонки данных race condition)
var counter int

func Increment(n int) int {
	for i := 0; i < n; i++ {
		counter++
	}
	return counter
}

// IncrementSafe это безопасный вариант с использованием sync.Mutex

var (
	mu          sync.Mutex
	CounterSafe int
)

func IncrementSafe(n int) int {
	mu.Lock()
	for i := 0; i < n; i++ {
		CounterSafe++
	}
	mu.Unlock()
	return CounterSafe
}
