package twoSum

func TwoSum(nums []int, target int) []int {
	begin := 0
	end := len(nums) - 1

	for {
		if begin >= end {
			break
		}
		sum := nums[begin] + nums[end]
		if sum == target {
			return []int {begin, end}
		}
		if sum > target {
			end = end - 1
			continue
		}
		if sum < target {
			begin = begin + 1
			continue
		}
	}
	return nil
}


//func UpperCase(str string) string {
//	return strings.ToUpper(str)
//}
