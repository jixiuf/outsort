package outsort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	left := IntSlice{1, 3, 5}
	right := IntSlice{2, 4, 6}
	result := make(IntSlice, 0, 6)
	Merge(&left, &right, &result)
	fmt.Println(result)
	if !result.Equal(IntSlice{1, 2, 3, 4, 5}) {
		t.Error("IntSlice Merge err should be IntSlice{1, 2, 3, 4, 5}")
	}
	if right[0] != 6 {
		t.Error("TestMerge right[0] should be 6")
	}
	if len(left) != 0 {
		t.Error("TestMerge left.len should be 0")
	}
	if cap(left) != 3 {
		t.Error("TestMerge left.cap should be 3")
	}

}

func TestMerge11(t *testing.T) {
	left := IntSlice{1, 3, 6}
	right := IntSlice{2, 4, 5}
	result := make(IntSlice, 0, 8)
	Merge(&left, &right, &result)
	fmt.Println(result)
	if !result.Equal(IntSlice{1, 2, 3, 4, 5}) {
		t.Error("IntSlice Merge err should be IntSlice{1, 2, 3, 4, 5}")
	}
	if !left.Equal(IntSlice{6}) {
		t.Error("IntSlice Merge err should be IntSlice{6}")
	}

}

func TestMerge2(t *testing.T) {
	var left IntSlice = nil
	right := IntSlice{2, 4, 6}
	result := make(IntSlice, 0, 6)
	Merge(&left, &right, &result)
	if !result.Equal(IntSlice{}) {
		t.Error("IntSlice Merge should by IntSlice{}")

	}

}

func TestMerge3(t *testing.T) {
	left := IntSlice{2, 4, 6}
	var right IntSlice = nil
	result := make(IntSlice, 0, 6)
	Merge(&left, &right, &result)
	fmt.Println(result)
	if !result.Equal(IntSlice{}) {
		t.Error("IntSlice Merge err")

	}

}

func TestMerge4(t *testing.T) {
	left := IntSlice{2, 4, 6}
	var right IntSlice = IntSlice{}
	result := make(IntSlice, 0, 6)
	Merge(&left, &right, &result)
	fmt.Println(result)
	if !result.Equal(IntSlice{}) {
		t.Error("IntSlice Merge err")

	}

}

// func TestCopy(t *testing.T) {
// 	// 测试 copy 时 有重叠的情况是否能正确copy
// 	var arr = IntSlice{1, 2, 3, 4, 5}
// 	copy(arr, arr[1:])
// 	if !arr.Equal(IntSlice{2, 3, 4, 5, 5}) {
// 		t.Error("TestCopy error")

// 	}
// 	fmt.Println(arr)

// }
