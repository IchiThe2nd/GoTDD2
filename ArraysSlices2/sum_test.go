package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

func TestReduce(t *testing.T) {
	t.Run("multiplicatyion of elements", func(t *testing.T) {
		multipy := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multipy, 1), 6)
	})
	t.Run("concantaate strings", func(t *testing.T) {
		concantate := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concantate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {

	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}
	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}
