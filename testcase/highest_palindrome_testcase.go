package testcase

type HighestPalindromeInput struct {
	String string
	K int
}


var HighestPalindromeTestCase = []TestCase{
	{
		Input: HighestPalindromeInput{
			String:  "3943",
			K:       1,
		},
		Expectation: "3993",
	},
	{
		Input: HighestPalindromeInput{
			String:  "932239",
			K:       1,
		},
		Expectation: "992299",
	},
	{
		Input: HighestPalindromeInput{
			String:  "12345",
			K:       2,
		},
		Expectation: "54345",
	},
	{
		Input: HighestPalindromeInput{
			String:  "124212521",
			K:       3,
		},
		Expectation: "925212529",
	},
	{
		Input: HighestPalindromeInput{
			String:  "36281924",
			K:       1,
		},
		Expectation: "-1",
	},
}