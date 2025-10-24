package currency_test

import (
	"testing"

	"github.com/alexvitayu/go-testing/currency"
)

type MockConverter struct {
	lastAmount float64
	lastFrom   string
	lastTo     string
	calls      int
}

func (m *MockConverter) Convert(amount float64, from, to string) float64 {
	m.lastAmount, m.lastFrom, m.lastTo = amount, from, to
	m.calls++
	return 42.0
}

// Проверка делегирования
// Функция PriceIn правильно вызывает Convert с теми же аргументами и возвращает ее результат
func TestPriceIn_DelegatesAndReturns(t *testing.T) {
	//создаём мок-объект, который выполняет интерфейс Converter
	mock := &MockConverter{}

	got := currency.PriceIn(200.0, "BYN", "USD", mock)
	if got != 42.0 {
		t.Fatalf("expected %v, but got %v", 42.0, got)
	}
	if mock.lastAmount != 200.0 || mock.lastFrom != "BYN" || mock.lastTo != "USD" {
		t.Fatalf("args mismatch: (%v,%s->%s)", mock.lastAmount, mock.lastFrom, mock.lastTo)
	}
	if mock.calls != 1 {
		t.Fatalf("calls: got %d, want 1", mock.calls)
	}
}

// Граничные случаи
func TestPriceIn_NegativeAndZero(t *testing.T) {
	mock := &MockConverter{}

	_ = currency.PriceIn(0, "RUB", "EUR", mock)
	if mock.lastAmount != 0 || mock.lastFrom != "RUB" || mock.lastTo != "EUR" {
		t.Fatalf("args mismatch for zero: (%v,%s->%s)", mock.lastAmount, mock.lastFrom, mock.lastTo)
	}
	_ = currency.PriceIn(-7, "USD", "EUR", mock)
	if mock.lastAmount != -7 || mock.lastFrom != "USD" || mock.lastTo != "EUR" {
		t.Fatalf("args mismatch for negative: (%v,%s->%s)", mock.lastAmount, mock.lastFrom, mock.lastTo)
	}

}
