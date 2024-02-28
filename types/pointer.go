package types

// P is a helper function to return a pointer to a value.
func P[T any](v T) *T {
	return &v
}
