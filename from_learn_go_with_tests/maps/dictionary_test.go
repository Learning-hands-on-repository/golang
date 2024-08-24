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
		_, err := dictionary.Search("random string")
		want := ErrNotFoundDefinition

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add new word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"test": "this is just a test"}

		// Act
		dictionary.Add("newKey", "newValue")

		// Assert
		want := "newValue"
		got, err := dictionary.Search("newKey")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"existingKey": "old value"}

		// Act
		err := dictionary.Add("existingKey", "new value")

		// Assert
		want := ErrWordExists
		assertError(t, err, want)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
