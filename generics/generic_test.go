package main

import "testing"

func TestAssertFunction(t *testing.T) {
	t.Run("assering on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNoEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNoEqual(t, "hello", "Grace")

	})

}

func TestStack(t *testing.T) {
	t.Run("integer stack ", func(t *testing.T) {
		myStackOfInts := new(StackOfInts)
		//check stack empty
		AssertTrue(t, myStackOfInts.IsEmpty())
		//add a thing and check not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		//add anothe and pop it back out
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackofStrings := new(StackOfStrings)

		//chek empty
		AssertTrue(t, myStackofStrings.IsEmpty())

		//add athing chek no empty
		myStackofStrings.Push("123")
		AssertFalse(t, myStackofStrings.IsEmpty())

		//add another thing pop it out
		myStackofStrings.Push("456")
		value, _ := myStackofStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = myStackofStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, myStackofStrings.IsEmpty())
	})
	t.Run("interface stack dx is horrid", func(t *testing.T) {
		myStackOfInts := new(StackOfInts)

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()

		// get our ints from out interface{}
		reallyFirstNum, ok := firstNum.(int)
		AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}

		reallySecondNum, ok := secondNum.(int)
		AssertTrue(t, ok) // and again!

		AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
	})
}

// Generic data structures to the rescuie
func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}
func AssertNoEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didnt want %+v to equal %+v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v want true", got)
	}
}
func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v want flase", got)
	}
}
