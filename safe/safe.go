package safe

func MustAt[T any](values []T, i int) T {
	if i < 0 || i >= len(values) {
		panic("index out of range")
	}
	return values[i]
}
