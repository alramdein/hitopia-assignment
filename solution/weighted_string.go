package solution

func WeightedStrings(s string, queries []int) []string {
	charWeights := make(map[rune]int)
	for i := 1; i <= 26; i++ {
		charWeights[rune('a'+i-1)] = i
	}

	weightsSet := make(map[int]struct{})

	n := len(s)
	i := 0

	for i < n {
		currentChar := rune(s[i])
		weight := charWeights[currentChar]

		j := i
		for j < n && rune(s[j]) == currentChar {
			weightsSet[(j-i+1)*weight] = struct{}{}
			j++
		}
		i = j
	}

	results := make([]string, len(queries))
	for idx, query := range queries {
		if _, found := weightsSet[query]; found {
			results[idx] = "Yes"
		} else {
			results[idx] = "No"
		}
	}

	return results
}
