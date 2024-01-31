package main

import "testing"

func TestHello (t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {

		want := "Hello, Chris"
		got := Hello("Chris")
	
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
	t.Run("saying hello when empty string is passed" , func (t *testing.T) {

		want := "Hello, World"
		got := Hello("")
	
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}