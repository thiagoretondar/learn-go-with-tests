package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	sums = make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
