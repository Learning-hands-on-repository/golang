package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessageHelper(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessageHelper(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessageHelper(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessageHelper(t, got, want)
	})
}

func assertCorrectMessageHelper(t testing.TB, got, want string) {
	t.Helper() // so in the callstack it's easier to trace back to caller of 'assertCorrectMessageHelper'
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
