package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known words", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})
	t.Run("unknonw Words", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q weanted %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q wanted %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is a test")
	want := "this is a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertStrings(t, got, want)

}
