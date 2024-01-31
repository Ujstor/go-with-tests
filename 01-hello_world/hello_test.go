package main

import "testing"

func TestHello (t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {

		want := "Hello, Chris"
		got := Hello("Chris", "")
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("saying hello when empty string is passed" , func (t *testing.T) {

		want := "Hello, World"
		got := Hello("", "")
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func (t *testing.T) {

		want := "Hola, Elodine"
		got := Hello("Elodine", "Spanish")
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func (t *testing.T) {

		want := "Bonjour, Elodine"
		got := Hello("Elodine", "French")
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Croatian", func (t *testing.T) {

		want := "Bok, G"
		got := Hello("G", "Cro")
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
} 