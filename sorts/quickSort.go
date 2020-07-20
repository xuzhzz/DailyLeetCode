package sorts

func quickSort(nums []int) []int {
	end := len(nums) - 1
	if end < 0 {
		return nil
	}
	start := 0
	pivot := nums[end]
	for end > start {
		for end > start && nums[start] < pivot {
			start++
		}
		nums[end] = nums[start]
		for end > start && nums[end] > pivot {
			end--
		}
		nums[start] = nums[end]
	}
	nums[end] = pivot
	res := append(quickSort(nums[0:end]), nums[end])
	return append(res, quickSort(nums[end+1:])...)
}
