package __two_sum

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("nums: %v\t target: %v\t twoSum: %v\n", nums, target, twoSum(nums, target))
}

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
