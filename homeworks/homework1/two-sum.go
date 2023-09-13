package hw1

func TwoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if requiredIdx, isPresent := indexMap[target-nums[i]]; isPresent {
			return []int{requiredIdx, i}
		}
		indexMap[nums[i]] = i
	}
	return nil
}
