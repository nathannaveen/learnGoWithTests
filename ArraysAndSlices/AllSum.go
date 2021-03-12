package ArraysAndSlices

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, ints := range numbersToSum {
		sums = append(sums, Sum(ints))
	}

	return sums
}
