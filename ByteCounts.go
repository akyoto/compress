package compress

func ByteCounts(data []byte) [256]int {
	var count [256]int

	for _, b := range data {
		count[b] += 1
	}

	return count
}
