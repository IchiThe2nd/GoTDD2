package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numberstoSum ...[]int) []int {
	lengthOfNumbers := len(numberstoSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numberstoSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numberstoSum ...[]int) []int {

	var sums []int
	for _, numbers := range numberstoSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
