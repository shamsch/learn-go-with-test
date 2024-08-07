package main

import "testing"

// test file needs to be postfixed xyz_test.go
// test function needs to prefixed by TestXYZ(t *testing.T)
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("support Spanish", func (t *testing.T){
		got := Hello("Remon", "Spanish")
		want := "Hola, Remon"

		assertCorrectMessage(t, got, want)
	})
	t.Run("support French", func (t *testing.T){
		got := Hello("Remon", "French")
		want := "Bonjour, Remon"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // when test fails, it will not give line number of this assertion function because we have declared it as a helper like this
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
