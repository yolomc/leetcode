/*
*
* 常见算法收集之各种排序
*
 */

//MergeSort 归并排序
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	m := len(nums) / 2
	a := MergeSort(nums[:m])
	b := MergeSort(nums[m:])
	return merge(a, b)
}
func merge(a []int, b []int) []int {
	aIdx := 0
	bIdx := 0
	res := make([]int, 0, len(a)+len(b))
	for aIdx < len(a) && bIdx < len(b) {
		if a[aIdx] > b[bIdx] {
			res = append(res, b[bIdx])
			bIdx++
		} else {
			res = append(res, a[aIdx])
			aIdx++
		}
	}
	res = append(res, a[aIdx:]...)
	res = append(res, b[bIdx:]...)
	return res
}

//QuickSort 快速排序
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	pIdx := partition(nums, l, r)
	quickSort(nums, l, pIdx-1)
	quickSort(nums, pIdx+1, r)
}
func partition(nums []int, l int, r int) int {
	pivot := nums[r]
	idx := l //大堆头下标
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}
	nums[idx], nums[r] = nums[r], nums[idx]
	return idx
}