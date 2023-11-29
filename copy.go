package util

func CopySlice[T any](src []T) (dst []T) {
	dst = make([]T, len(src))
	copy(dst, src)
	return
}

func CopyMap[K comparable, V any](src map[K]V) (dst map[K]V) {
	dst = make(map[K]V, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return
}
