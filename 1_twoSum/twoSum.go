package twoSum

import "sort"

type twoNum struct {
	num int
	index int
}

type twoNums []twoNum

func (t twoNums) Len() int {return len(t)}
func (t twoNums) Less(i, j int) bool {return t[i].num < t[j].num}
func (t twoNums) Swap(i, j int) {t[i], t[j] = t[j], t[i]}

func TwoSum(nums []int, target int) []int {
	sortNums := make(twoNums, 0)
	for idx, n := range nums {
		t := twoNum{n, idx}
		sortNums = append(sortNums, t)
	}
	sort.Sort(sortNums)
	begin := 0
	end := len(nums) - 1
	for {
		if begin >= end {
			break
		}
		//sum := nums[begin] + nums[end]
		sum := sortNums[begin].num + sortNums[end].num
		if sum == target {
			return []int {sortNums[begin].index, sortNums[end].index}
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
