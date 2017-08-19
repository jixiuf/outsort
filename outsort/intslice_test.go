package outsort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	left := IntSlice{1, 3, 5}
	right := IntSlice{2, 4, 6}
	result := make(IntSlice, 0, 6)
	Merge(left, right, &result)
	fmt.Println(result)
	if !result.Equal(IntSlice{1, 2, 3, 4, 5, 6}) {
		t.Error("IntSlice Merge err")

	}

}

func TestMerge2(t *testing.T) {
	var left IntSlice = nil
	right := IntSlice{2, 4, 6}
	result := make(IntSlice, 0, 6)
	Merge(left, right, &result)
	fmt.Println(result)
	if !result.Equal(IntSlice{2, 4, 6}) {
		t.Error("IntSlice Merge err")

	}

}

func TestMerge3(t *testing.T) {
	left := IntSlice{2, 4, 6}
	var right IntSlice = nil
	result := make(IntSlice, 0, 6)
	Merge(left, right, &result)
	fmt.Println(result)
	if !result.Equal(IntSlice{2, 4, 6}) {
		t.Error("IntSlice Merge err")

	}

}

func TestMerge4(t *testing.T) {
	left := IntSlice{2, 4, 6}
	var right IntSlice = IntSlice{}
	result := make(IntSlice, 0, 6)
	Merge(left, right, &result)
	fmt.Println(result)
	if !result.Equal(IntSlice{2, 4, 6}) {
		t.Error("IntSlice Merge err")

	}

}
