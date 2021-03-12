package ArraysAndSlices

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, ints := range numbersToSum {
		if len(ints) == 0 {
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, Sum(ints[1:]))
	}

	return sums
}
