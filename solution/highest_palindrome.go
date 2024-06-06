package solution

func HighestPalindrome(s string, k int) string {
	chars := []rune(s)
	n := len(chars)

	// Step 1: Create the palindrome with minimum changes
	left := 0
	right := n - 1
	changes := make([]int, n)

	for left < right {
		if chars[left] != chars[right] {
			if chars[left] > chars[right] {
				chars[right] = chars[left]
			} else {
				chars[left] = chars[right]
			}
			changes[left] = 1
			k--
		}
		left++
		right--
	}

	if k < 0 {
		return "-1"
	}

	// Step 2: Maximize the palindrome value
	left = 0
	right = n - 1
	for left <= right {
		if left == right && k > 0 {
			chars[left] = '9'
		}
		if chars[left] < '9' {
			if changes[left] == 1 && k >= 1 {
				chars[left] = '9'
				chars[right] = '9'
				k--
			} else if changes[left] == 0 && k >= 2 {
				chars[left] = '9'
				chars[right] = '9'
				k -= 2
			}
		}
		left++
		right--
	}

	return string(chars)
}

// func isPalindrome(s string) bool {
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[len(s)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func max(a, b string) string {
// 	ai, _ := strconv.Atoi(a)
// 	bi, _ := strconv.Atoi(b)
// 	if ai > bi {
// 		return a
// 	}
// 	return b
// }
