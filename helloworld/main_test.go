package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying hello to gopher", func(t *testing.T) {
		got := Hello("Gopher")
		want := "Hello, Gopher"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("")
		want := "Hello"
		assertCorrectMessage(t, got, want)
	})
}
