package util

import "slices"

// Deprecated
// IntDeduplicate() is deprecated, please use DeduplicateSlice() instead.
// int 数组去重
func IntDeduplicate(slice []int) (result []int) {
	var dupMap = make(map[int]struct{})
	for _, v := range slice {
		length := len(dupMap)
		dupMap[v] = struct{}{}
		if len(dupMap) != length {
			result = append(result, v)
		}
	}
	return result
}

// Deprecated
// IntMergeDeduplicate() is deprecated, please use MergeDeduplicateSlice() instead.
// int 数组合并+去重
func IntMergeDeduplicate(slice1, slice2 []int) (result []int) {
	slice1 = append(slice1, slice2...)
	return IntDeduplicate(slice1)
}

// Deprecated
// IntIntersect() is deprecated, please use IntersectSlice() instead.
// int 数组，slice1 和 slice2 交集
func IntIntersect(slice1, slice2 []int) (result []int) {
	m := make(map[int]struct{})
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		_, ok := m[v]
		if ok {
			result = append(result, v)
		}
	}
	return
}

// Deprecated
// IntUnion() is deprecated, please use UnionSlice() instead.
// int 数组，slice1 和 slice2 并集
func IntUnion(slice1, slice2 []int) (result []int) {
	m := make(map[int]struct{})
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		_, ok := m[v]
		if !ok {
			m[v] = struct{}{}
		}
	}
	for k := range m {
		result = append(result, k)
	}
	return
}

// Deprecated
// IntRemoveElementByIndex() is deprecated, please use RemoveSliceElementByIndex() instead.
// int 数组，根据index移除元素
// return new slice
func IntRemoveElementByIndex(slice []int, index int) (result []int) {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result = append(slice[:index], slice[index+1:]...)
	return
}

// Deprecated
// IntRemoveElement() is deprecated, please use RemoveSliceElement() instead.
// int 数组，移除元素
// If n < 0, there is no limit on the number of remove.
// return new slice
func IntRemoveElement(slice []int, elem, n int) (result []int) {
	if n == 0 {
		return append(result, slice...) // 返回输入切片的复制品
	}
	// 复制输入切片到result中
	result = append(result, slice...)
	i, j := 0, 0
	for j < len(result) {
		if result[j] != elem || n == 0 {
			result[i] = result[j]
			i++
		} else {
			n--
		}
		j++
	}
	return result[:i]
}

// Deprecated
// StringDeduplicate() is deprecated, please use DeduplicateSlice() instead.
// string 数组去重
func StringDeduplicate(slice []string) (result []string) {
	var dupMap = make(map[string]struct{})
	for _, v := range slice {
		length := len(dupMap)
		dupMap[v] = struct{}{}
		if len(dupMap) != length {
			result = append(result, v)
		}
	}
	return result
}

// Deprecated
// StringMergeDeduplicate() is deprecated, please use MergeDeduplicateSlice() instead.
// string 数组合并+去重
func StringMergeDeduplicate(slice1, slice2 []string) (result []string) {
	slice1 = append(slice1, slice2...)
	return StringDeduplicate(slice1)
}

// Deprecated
// StringRemoveElementByIndex() is deprecated, please use RemoveSliceElementByIndex() instead.
// string 数组，根据index移除元素
// return new slice
func StringRemoveElementByIndex(slice []string, index int) (result []string) {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result = append(slice[:index], slice[index+1:]...)
	return
}

// Deprecated
// string 数组，移除元素
// If n < 0, there is no limit on the number of remove.
// return new slice
func StringRemoveElement(slice []string, elem string, n int) (result []string) {
	if n == 0 {
		return append(result, slice...) // 返回输入切片的复制品
	}
	// 复制输入切片到result中
	result = append(result, slice...)
	i, j := 0, 0
	for j < len(result) {
		if result[j] != elem || n == 0 {
			result[i] = result[j]
			i++
		} else {
			n--
		}
		j++
	}
	return result[:i]
}

// Deprecated
// FilterIntSlice() is deprecated, please use FilterSlice() instead.
// 过滤数组，去除src在dst中存在的item
// src[1,2,3,4,5]   dst[2,4,6,8]	result[1,3,5]
func FilterIntSlice(src []int, dst []int) (result []int) {
	aMap := make(map[int]struct{})
	result = make([]int, 0)
	for _, v := range dst {
		aMap[v] = struct{}{}
	}
	for _, v := range src {
		if _, has := aMap[v]; !has {
			result = append(result, v)
		}
	}
	return result
}

// Deprecated
// FilterStringSlice() is deprecated, please use FilterSlice() instead.
// 过滤数组，去除src在dst中存在的item
// src["a","b","c","d","e"]   dst["b","d","f","h"]	result["a","c","e"]
func FilterStringSlice(src []string, dst []string) (result []string) {
	aMap := make(map[string]struct{})
	result = make([]string, 0)
	for _, v := range dst {
		aMap[v] = struct{}{}
	}
	for _, v := range src {
		if _, has := aMap[v]; !has {
			result = append(result, v)
		}
	}
	return result
}

// ==================new==================

// 数组去重
func DeduplicateSlice[T comparable](slice []T) (result []T) {
	var dupMap = make(map[T]struct{})
	for _, v := range slice {
		length := len(dupMap)
		dupMap[v] = struct{}{}
		if len(dupMap) != length {
			result = append(result, v)
		}
	}
	return result
}

// 数组合并+去重
func MergeDeduplicateSlice[T comparable](slice1, slice2 []T) (result []T) {
	slice1 = append(slice1, slice2...)
	return DeduplicateSlice(slice1)
}

// slice1 和 slice2 交集
func IntersectSlice[T comparable](slice1, slice2 []T) (result []T) {
	m := make(map[T]int)
	for _, v := range slice1 {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		}
	}
	for _, v := range slice2 {
		count, ok := m[v]
		if ok && count == 1 {
			m[v]++
			result = append(result, v)
		}
	}
	return
}

// string 数组，slice1 和 slice2 并集
func UnionSlice[T comparable](slice1, slice2 []T) (result []T) {
	m := make(map[T]struct{})
	for _, v := range slice1 {
		_, ok := m[v]
		if !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}
	for _, v := range slice2 {
		_, ok := m[v]
		if !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}
	return
}

// 数组根据index移除元素
// return new slice
func RemoveSliceElementByIndex[T comparable](slice []T, index int) (result []T) {
	clone := slices.Clone(slice)
	if index < 0 || index >= len(clone) {
		return clone
	}
	result = append(clone[:index], clone[index+1:]...)
	return
}

// 数组移除元素
// If n < 0, there is no limit on the number of remove.
// return new slice
func RemoveSliceElement[T comparable](slice []T, elem T, n int) (result []T) {
	if n == 0 {
		return append(result, slice...) // 返回输入切片的复制品
	}
	// 复制输入切片到result中
	result = append(result, slice...)
	i, j := 0, 0
	for j < len(result) {
		if result[j] != elem || n == 0 {
			result[i] = result[j]
			i++
		} else {
			n--
		}
		j++
	}
	return result[:i]
}

// 过滤数组，去除src在dst中存在的item
// src[1,2,3,4,5]   dst[2,4,6,8]	result[1,3,5]
// src["a","b","c","d","e"]   dst["b","d","f","h"]	result["a","c","e"]
func FilterSlice[T comparable](src, dst []T) (result []T) {
	aMap := make(map[T]struct{})
	result = make([]T, 0)
	for _, v := range dst {
		aMap[v] = struct{}{}
	}
	for _, v := range src {
		if _, has := aMap[v]; !has {
			result = append(result, v)
		}
	}
	return result
}
