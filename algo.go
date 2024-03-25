package main

func getScore(results []Result) int {
	total := 0
	for _, result := range results {
		total += result.Count
	}
	return total
}
