package sorts

func swap(nums []int, idx1, idx2 int) {
	tmp := nums[idx1]
	nums[idx1] = nums[idx2]
	nums[idx2] = tmp
}

func adjustHeap(nums []int, idx, numsLen int) {
	for {
		maxIdx := 2*idx + 1
		if maxIdx >= numsLen {
			break
		} else if 2*idx+2 < numsLen && nums[2*idx+1] < nums[2*idx+2] {
			maxIdx = 2*idx + 2
		}
		if nums[idx] > nums[maxIdx] {
			break
		}
		swap(nums, maxIdx, idx)
		idx = maxIdx
	}
}

func heapSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// init
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
	// adjust
	for last := len(nums) - 1; last >= 1; last-- {
		swap(nums, 0, last)
		adjustHeap(nums, 0, last)
	}
	return nums
}
