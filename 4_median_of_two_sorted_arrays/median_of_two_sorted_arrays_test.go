package __median_of_two_sorted_arrays

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	var (
		num1 = []int{1, 2, 34, 45}
		num2 = []int{23, 24, 245}
	)
	fmt.Printf("median of two sorted arrays:num1:%+v,num2:%+v is %v\n", num1, num2, findMedianSortedArrays(num1, num2))
}
