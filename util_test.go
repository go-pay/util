package util

import (
	"slices"
	"testing"
)

func TestMergeDeduplicateSlice(t *testing.T) {
	in := []int{3, 3, 5, 7, 14, 11, 13, 15, 12}  // slice1
	in2 := []int{3, 4, 5, 7, 18, 11, 22, 15, 35} // slice2
	merged := MergeDeduplicateSlice(in, in2)
	t.Logf("in: %d", in)   // in: [3 3 5 7 14 11 13 15 12]
	t.Logf("in2: %d", in2) // in2: [3 4 5 7 18 11 22 15 35]
	slices.Sort(merged)
	t.Logf("merged: %d", merged) // merged: [3 4 5 7 11 12 13 14 15 18 22 35]

	ins := []string{"abc", "hello", "fhgk", "jerry", "world", "jerry", "abc", "hello"}
	ins2 := []string{"dfsf", "hello", "qwer", "jerry", "hello", "tom", "abc", "fuck"}
	merged2 := MergeDeduplicateSlice(ins, ins2)
	t.Logf("ins: %s", ins)         // ins: [abc hello fhgk jerry world jerry abc hello]
	t.Logf("ins2: %s", ins2)       // ins2: [dfsf hello qwer jerry hello tom abc fuck]
	t.Logf("merged2: %s", merged2) // merged2: [abc hello fhgk jerry world dfsf qwer tom fuck]
	slices.Sort(merged2)
	t.Logf("merged2_sort: %s", merged2) // merged2_sort: [abc dfsf fhgk fuck hello jerry qwer tom world]
}

func TestDeduplicateSlice(t *testing.T) {
	ins := []string{"abc", "hello", "fhgk", "jerry", "world", "jerry", "abc", "hello"}
	merged := DeduplicateSlice(ins)
	t.Logf("in: %s", ins)        // in: [abc hello fhgk jerry world jerry abc hello]
	t.Logf("merged: %s", merged) // merged: [abc hello fhgk jerry world]

	ins2 := []int{1, 2, 3, 3, 4, 5, 5, 5, 6, 8, 10}
	merged2 := DeduplicateSlice(ins2)
	t.Logf("in2: %d", ins2)        // in2: [1 2 3 3 4 5 5 5 6 8 10]
	t.Logf("merged2: %d", merged2) // merged: [1 2 3 4 5 6 8 10]
}

func TestIntersectSlice(t *testing.T) {
	ins := []string{"abc", "hello", "fhgk", "jerry", "world", "jerry", "abc", "hello"}
	ins2 := []string{"dfsf", "hello", "qwer", "jerry", "hello", "tom", "abc", "fuck"}
	result := IntersectSlice(ins, ins2)
	t.Logf("ins: %s", ins)       // ins: [abc hello fhgk jerry world jerry abc hello]
	t.Logf("ins2: %s", ins2)     // ins2: [dfsf hello qwer jerry hello tom abc fuck]
	t.Logf("result: %s", result) // result: [hello jerry abc]
	slices.Sort(result)
	t.Logf("result_sort: %s", result) // result_sort: [abc hello jerry]

	in := []int{3, 3, 5, 7, 14, 11, 13, 15, 12}  // slice1
	in2 := []int{3, 4, 5, 7, 18, 11, 22, 15, 35} // slice2
	result2 := IntersectSlice(in, in2)
	t.Logf("in: %d", in)           // in: [3 3 5 7 14 11 13 15 12]
	t.Logf("in2: %d", in2)         // in2: [3 4 5 7 18 11 22 15 35]
	t.Logf("result2: %d", result2) // result: [3 5 7 11 15]
}

func TestUnionSlice(t *testing.T) {
	ins := []string{"abc", "hello", "fhgk", "jerry", "world", "jerry", "abc", "hello"}
	ins2 := []string{"dfsf", "hello", "qwer", "jerry", "hello", "tom", "abc", "fuck"}
	result := UnionSlice(ins, ins2)
	t.Logf("ins: %s", ins)       // ins: [abc hello fhgk jerry world jerry abc hello]
	t.Logf("ins2: %s", ins2)     // ins2: [dfsf hello qwer jerry hello tom abc fuck]
	t.Logf("result: %s", result) // result: [abc hello fhgk jerry world dfsf qwer tom fuck]

	in := []int{1, 2, 3, 4}  // slice1
	in2 := []int{2, 4, 6, 8} // slice2
	result2 := UnionSlice(in, in2)
	t.Logf("in: %d", in)           // in: [1 2 3 4]
	t.Logf("in2: %d", in2)         // in2: [2 4 6 8]
	t.Logf("result2: %d", result2) // result2: [1 2 3 4 6 8]
}

func TestRemoveSliceElementByIndex(t *testing.T) {
	in := []int{1, 2, 3, 4} // slice1
	result := RemoveSliceElementByIndex(in, 1)
	t.Logf("in: %d", in)         // result: [1 2 3 4]
	t.Logf("result: %d", result) // result: [1 3 4]

	ins := []string{"abc", "hello", "fhgk", "jerry", "hello"}
	result2 := RemoveSliceElementByIndex(ins, 3)
	t.Logf("ins: %s", ins)         // result: [abc hello fhgk jerry hello]
	t.Logf("result2: %s", result2) // result: [abc hello fhgk hello]
}

func TestRemoveSliceElement(t *testing.T) {
	in := []int{1, 2, 3, 4, 4, 5, 4, 4} // slice1
	result := RemoveSliceElement(in, 4, 2)
	t.Logf("in: %d", in)         // result: [1 2 3 4 4 5 4 4]
	t.Logf("result: %d", result) // result: [1 2 3 5 4 4]

	ins := []string{"abc", "hello", "abc", "abc", "hello", "abc"}
	result2 := RemoveSliceElement(ins, "abc", 3)
	t.Logf("ins: %s", ins)         // ins: [abc hello abc abc hello abc]
	t.Logf("result2: %s", result2) // result: [hello hello abc]
}

func TestFilterSlice(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // slice1
	in2 := []int{1, 3, 5, 7, 9}
	result := FilterSlice(in, in2)
	t.Logf("in: %d", in)         // in: [1 2 3 4 5 6 7 8 9 10]
	t.Logf("in2: %d", in2)       // in2: [1 3 5 7 9]
	t.Logf("result: %d", result) // result: [2 4 6 8 10]

	ins := []string{"abc", "hello", "world", "hello", "jerry"}
	ins2 := []string{"hello", "jerry"}
	result2 := FilterSlice(ins, ins2)
	t.Logf("ins: %s", ins)         // ins: [abc hello world hello jerry]
	t.Logf("ins2: %s", ins2)       // ins2: [hello jerry]
	t.Logf("result2: %s", result2) // result: [abc world]
}

func TestCopySlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := CopySlice[int](a)
	t.Logf("a:%p, a: %d", &a, a) // a: [1 2 3 4 5]
	t.Logf("b:%p, b: %d", &b, b) // b: [1 2 3 4 5]

	a[0] = 100
	a[2] = 300
	t.Logf("a:%p, a: %d", &a, a) // a: [100 2 300 4 5]
	t.Logf("b:%p, b: %d", &b, b) // b: [1 2 3 4 5]
}

func TestCopyMap(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3}
	b := CopyMap[string, int](a)
	t.Logf("a:%p, a: %v", &a, a) // a: map[a:1 b:2 c:3]
	t.Logf("b:%p, b: %v", &b, b) // b: map[a:1 b:2 c:3]

	a["a"] = 100
	a["c"] = 300
	t.Logf("a:%p, a: %v", &a, a) // a: map[a:100 b:2 c:300]
	t.Logf("b:%p, b: %v", &b, b) // b: map[a:1 b:2 c:3]
}
