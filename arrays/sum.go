package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) (sums []int) {
	//	sums = make([]int, len(numbersToSum))
	//	for i, numbers := range numbersToSum {
	//		sums[i] = Sum(numbers)
	//	}
	//	return sums
	// Another implementation 👇🏻
	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		switch len(numbers) {
		case 0:
			sums = append(sums, 0)
		case 1:
			sums = append(sums, numbers[0])
		default:
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
