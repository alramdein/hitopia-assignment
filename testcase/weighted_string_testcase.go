package testcase

type WeightedStringInput struct {
	String string
	Queries []int
}

var WeightedStringTestCase = []TestCase{
	{
		Input: WeightedStringInput{
			String:  "abbcccd",
			Queries: []int{1, 3, 9, 8},
		},
		Expectation: []string{"Yes", "Yes", "Yes", "No"},
	},
	{
		Input: WeightedStringInput{
			String:  "abcdefg",
			Queries: []int{1, 2, 3, 4, 5, 6, 7},
		},
		Expectation: []string{"Yes", "Yes", "Yes", "Yes", "Yes", "Yes", "Yes"},
	},
	{
		Input: WeightedStringInput{
			String:  "xyz",
			Queries: []int{1, 2, 3, 4, 5, 6, 7},
		},
		Expectation: []string{"No", "No", "No", "No", "No", "No", "No"},
	},
	{
		Input: WeightedStringInput{
			String:  "aaaaabbbccc",
			Queries: []int{1, 3, 5, 7, 9, 11, 13},
		},
		Expectation: []string{"Yes", "Yes", "Yes", "No", "Yes", "No", "No"},
	},
	{
		Input: WeightedStringInput{
			String:  "aaabbbccc",
			Queries: []int{1, 2, 3, 4, 5, 6, 7},
		},
		Expectation: []string{"Yes", "Yes", "Yes", "Yes", "No", "Yes", "No"},
	},
}