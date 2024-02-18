package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test2"
	definition := "this is test"

	dictionary.Add(word, definition)

	assertDefinitione(t, dictionary, word, definition)
}

//helper functions
func assertStrings(t testing.TB, got, want string) {
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

func assertDefinitione(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	word, err := dictionary.Search(word)
	if err != nil {
		t.Errorf("error %q", err)
	}

	assertStrings(t, word, definition)
}
