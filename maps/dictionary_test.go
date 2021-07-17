package main

import "testing"

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("foo")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "elephant"
		definition := "A majestic animal"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new description")
		assertError(t, err, ErrWordExists)
	})
	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("elephant", "A majestic animal")
		got, _ := dictionary.Search("elephant")
		want := "A majestic animal"

		assertError(t, err, nil)
		assertString(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updates word", func(t *testing.T) {
		dictionary := Dictionary{"monkey": "Another animal"}
		dictionary.Update("monkey", "A silly animal")
		want := "A silly animal"

		got, _ := dictionary.Search("monkey")
		assertString(t, got, want)
	})
	t.Run("missing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("foo", "bar")

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deletes word", func(t *testing.T) {
		dictionary := Dictionary{"shoe": "For your feet"}
		dictionary.Delete("shoe")

		_, err := dictionary.Search("shoe")
		assertError(t, err, ErrNotFound)
	})
	t.Run("not found", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("shoe")

		assertError(t, err, ErrNotFound)
	})
}
