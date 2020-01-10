package __two_sum

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("nums: %v\t target: %v\t twoSum: %v\n", nums, target, twoSum(nums, target))
}
