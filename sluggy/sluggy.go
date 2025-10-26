package sluggy

import (
	"strings"
	"unicode"
)

// Slug нормализует строку для использования в URL.
func Slug(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	prevHyphen := false
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevHyphen = false
			continue
		}
		if !prevHyphen {
			b.WriteByte('-')
			prevHyphen = true
		}
	}
	out := b.String()
	out = strings.Trim(out, "-")
	// Схлопывание уже обеспечено логикой prevHyphen
	return out
}

func ReturnMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type OrderStorage interface {
	GetOrder(id int) (string, error)
}

type OrderService struct {
	storage OrderStorage
}

func NewOrderService(storage OrderStorage) *OrderService {
	return &OrderService{storage: storage}
}

func (s *OrderService) GetOrderName(id int) (string, error) {
	return s.storage.GetOrder(id)
}
