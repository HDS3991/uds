package utils

// UniqueArray 数组去重
func UniqueArray[T int | int32 | int64 | string | float32 | float64](arr []T) []T {
	r := make([]T, 0, len(arr))
	m := make(map[T]bool, len(arr))
	for _, v := range arr {
		if !m[v] {
			m[v] = true
			r = append(r, v)
		}
	}
	return r
}

// IsContain 数组中是否包含 X 元素
func IsContain[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
