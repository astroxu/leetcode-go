package __longest_palindromic_substring

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	arr := []byte(s)
	length := len(arr)
	data := make([][]int, length)
	for i := 0; i < length; i++ {
		data[i] = make([]int, length)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i == j {
				data[i][j] = 1
			}
		}
	}

	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			if i >= j {
				continue
			}
			if i == j-1 {
				if arr[i] == arr[j] {
					data[j][i] = 2
				}
			}
			if i <= j-2 {
				if data[j-1][i+1] != 0 && (arr[i] == arr[j]) {
					data[j][i] = data[j-1][i+1] + 2
				}
			}
		}
	}

	maxI := 0
	maxJ := 0
	maxLen := 0
	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			if data[j][i] > maxLen {
				maxI = i
				maxJ = j
				maxLen = data[j][i]
			}
		}
	}

	return s[maxI : maxJ+1]
}
