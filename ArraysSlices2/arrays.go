package main

// reduce takes a collection of things and performs a func returning a thing

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

// SUM ADDAS TOTAL aof all numbers
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// sumall trails calculates rthe sums of all but teh first number
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {

	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
