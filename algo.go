package main

func getScore(results []Result) int {
	total := 0
	for _, result := range results {
		total += result.TotalNumRequests
	}
	return total
}
