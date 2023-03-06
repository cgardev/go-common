package ptr

func ToPtr[T any](v T) *T {
	return &v
}

func ToValue[T any](p *T) T {
	return *p
}
