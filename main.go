package main

import (
	"fmt"
	"hitopia-alif/solution"
	"hitopia-alif/testcase"
)


func main() {
	executeWeightedStringSolution()
	executeBalancedBracketSolution()
	executeHighestPalindromeSolution()
}

func executeWeightedStringSolution() {
	fmt.Println("#Question No.1 Weighted String")
	fmt.Println()
	for _, tc := range testcase.WeightedStringTestCase {
		input := tc.Input.(testcase.WeightedStringInput)
		expectation := tc.Expectation.([]string)

		output := solution.WeightedStrings(input.String, input.Queries)

		fmt.Println("string: ", input.String)
		fmt.Println("queries: ", input.Queries)
		fmt.Println("output: ", output)
		fmt.Println("expectation: ", expectation)

		if isEqual(output, expectation) {
			fmt.Println("Test case passed.")
		} else {
			fmt.Println("Test case failed.")
		}

		fmt.Println()
	}
}

func executeBalancedBracketSolution() {
	fmt.Println("#Question No.2 Balanced Bracket")
	fmt.Println()
	for _, tc := range testcase.BalancedBracketTestCase {
		input := tc.Input.(string)
		expectation := tc.Expectation.(string)

		output := solution.BalancedBracket(input)

		fmt.Println("Input: ", input)
		fmt.Println("Output: ", output)
		fmt.Println("Expectation: ", expectation)

		if output == expectation {
			fmt.Println("Test case passed.")
		} else {
			fmt.Println("Test case failed.")
		}

		fmt.Println()
	}
}

func executeHighestPalindromeSolution() {
	fmt.Println("#Question No.3 Highest Palindrome")
	fmt.Println()
	for _, tc := range testcase.HighestPalindromeTestCase {
		input := tc.Input.(testcase.HighestPalindromeInput)
		expectation := tc.Expectation.(string)

		output := solution.HighestPalindrome(input.String, input.K)

		fmt.Println("string: ", input.String)
		fmt.Println("k: ", input.K)
		fmt.Println("output: ", output)
		fmt.Println("expectation: ", expectation)

		if output ==  expectation {
			fmt.Println("Test case passed.")
		} else {
			fmt.Println("Test case failed.")
		}

		fmt.Println()
	}
}

func isEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}