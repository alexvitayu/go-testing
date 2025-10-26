package sluggy_test

import (
	"fmt"
	"testing"

	"github.com/alexvitayu/go-testing/sluggy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSlug(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want string
	}{
		{"withSpacesAndPunctuation", "Go is, the Best", "go-is-the-best"},
		{"withRepeatedSpaces", "today  is   windy    day", "today-is-windy-day"},
		{"withUnicodeLettersAndToLower", "СегодНя Всё УДАСТСЯ", "сегодня-всё-удастся"},
		{"emptyString", "", ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sluggy.Slug(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}

// практика(простое задание)
func TestReturnMin(t *testing.T) {
	testCases := []struct {
		name string
		a, b int
		want int
	}{
		{"basic", 4, 5, 4},
		{"negative", -5, -10, -10},
		{"equal", 6, 6, 6},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sluggy.ReturnMin(tc.a, tc.b)
			assert.Equal(t, tc.want, got)
		})
	}
}

// практика(сложное задание)

// настраиваем мок для хранилища
type MockOrderService struct {
	mock.Mock
}

// метод, который имплементирует контракт интерфейса OrderStorage
func (m *MockOrderService) GetOrder(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestNewOrderService_Correct(t *testing.T) {
	m := new(MockOrderService) // инициализируем моковое хранилище для данного теста

	//настраиваем поведение метода нашего хранилища: при GetOrder(1) вернуть "OrderA", nil
	m.On("GetOrder", 1).Return("OrderA", nil)
	//настраиваем поведение метода нашего хранилища: при GetOrder(99) вернуть "" и ошибку
	m.On("GetOrder", 99).Return("", fmt.Errorf("not found"))

	service := sluggy.NewOrderService(m) // нициализируем наш сервис, который принимает моковое хранилище

	t.Run("successfull_call", func(t *testing.T) {
		order, err := service.GetOrderName(1) // метод нашего сервиса извлекает из мокового хранилища товар по id=1
		require.NoError(t, err)
		assert.Equal(t, "OrderA", order) // утверждаем, что наш сервис правильно извлёк товар по id=1
		m.AssertCalled(t, "GetOrder", 1) //утверждаем, что метод GetOrder действительно вызвался
		m.AssertNumberOfCalls(t, "GetOrder", 1)
	})
	t.Run("not_found", func(t *testing.T) {
		_, err := service.GetOrderName(99) // метод нашего сервиса извлекает из мокового хранилища товар по id=99
		require.EqualError(t, err, "not found")
		m.AssertCalled(t, "GetOrder", 99)
	})

}
