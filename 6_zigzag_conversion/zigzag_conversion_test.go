package __zigzag_conversion

import (
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {
	var (
		str = "PAYPALISHIRING"
		num = 3
	)
	fmt.Printf("zigzag conversion \"%s\" num: %v is \"%s\"\n", str, num, convert(str, num))
}
