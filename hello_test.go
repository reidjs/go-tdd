package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Reid")
		want := "Hello, Reid"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
