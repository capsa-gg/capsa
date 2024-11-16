package util

// PtrTo is a util that can be used to create a pointer for a raw value.
// This prevents the need to create an intermediate variable.
// Example: ptrToInt := util.PtrTo(42)
func PtrTo[T any](value T) *T {
	return &value
}
