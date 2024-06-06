package solution

func BalancedBracket(s string) string {
	stack := make([]rune, 0)

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 {
				return "No"
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (char == ')' && top != '(') || (char == ']' && top != '[') || (char == '}' && top != '{') {
				return "No"
			}
		}
	}

	if len(stack) == 0 {
		return "Yes"
	} else {
		return "No"
	}
}