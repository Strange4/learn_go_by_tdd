package sum

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(slicesToSum ...[]int) (sums []int) {
	sums = make([]int, len(slicesToSum))
	for i, slice := range slicesToSum {
		sums[i] = Sum(slice)
	}
	return
}

func SumAllTails(slicesToSum ...[]int) []int {
	sums := make([]int, len(slicesToSum))
	for i, slice := range slicesToSum {
		if len(slice) > 0 {
			tail := slice[1:]
			sums[i] = Sum(tail)
		} else {
			sums[i] = 0
		}
	}
	return sums
}
