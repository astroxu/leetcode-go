package __two_sum

func twoSum(nums []int, target int) []int {
	length := len(nums)
	a := map[int]int{}
	for i := 0; i < length; i++ {
		if v, ok := a[target-nums[i]]; ok {
			return []int{v, i}
		}
		a[nums[i]] = i
	}
	return []int{}

}
