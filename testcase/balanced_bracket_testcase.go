package testcase

var BalancedBracketTestCase = []TestCase{
	{
		Input: "{[()]}",
		Expectation: "Yes",
	},
	{
		Input: "{[(])}",
		Expectation: "No",
	},
	{
		Input: "{{[[(())]]}}",
		Expectation: "Yes",
	},
	{
		Input: "{ ( ( [ ] ) [ ] ) [ ] }",
		Expectation: "Yes",
	},
	{
		Input: "{ [ ( ] ) }",
		Expectation: "No",
	},
}