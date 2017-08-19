package outsort

import "sort"

type IntSlice []int32

func (l IntSlice) Equal(other IntSlice) bool {
	if len(l) != len(other) {
		return false
	}
	for key, value := range l {
		if value != other[key] {
			return false
		}
	}
	return true
}

func (l IntSlice) Sort() {
	sort.Sort(l)
}

// 实现sort 接口
func (l IntSlice) Len() int {
	return len(l)
}
func (l IntSlice) Less(i, j int) bool {
	return l[i] < l[j]
}
func (l IntSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// 二路归并,请注意，调用之前left,right 需要已经各自有序
// 此处仅做归并
// 请确保在调用前result 的cap>=len(left)+left(right)
// 之所以将result传进来，而不直接return过去,是为了重用result原来的内存，避免新的内存分配
func Merge(left, right *IntSlice, result *IntSlice) { //
	if len(*right) == 0 || len(*left) == 0 {
		return
	}

	l, r := 0, 0
	for l < len(*left) && r < len(*right) {
		if (*left)[l] < (*right)[r] {
			*result = append(*result, (*left)[l])
			l++
		} else {
			*result = append(*result, (*right)[r])
			r++
		}
	}

	// 后面不再直接把剩余的元素直接append的结尾，因为
	// 后续的元素还需要与文件中未读的元素进行排序
	// *result = append(*result, *left[l:]...)
	// *result = append(*result, (*right)[r:]...)

	// 把未排序的部分，移动到前列(即清除掉已经合并到result的部分)
	copy(*left, (*left)[l:])
	*left = (*left)[0:len((*left)[l:])]
	copy(*right, (*right)[r:])
	*right = (*right)[0:len((*right)[r:])]

	return

}
