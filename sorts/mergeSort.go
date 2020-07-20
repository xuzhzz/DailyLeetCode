package sorts

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	res := make([]int, 0)
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// merge
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
