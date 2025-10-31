package util

func GetFromSlice[T any](slice []T, idx int) (T, bool) {
	var zero T

	if idx < 0 || idx >= len(slice) {
		return zero, false
	}

	return slice[idx], true
}
