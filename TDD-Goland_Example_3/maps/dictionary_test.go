package maps

import "testing"

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Expected: %s, Wanted: %s", want, got)
	}
}
func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	t.Run("If the word is a match", func(t *testing.T) {

		got := Search(dictionary, "test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("If Word wasn't found", func(t *testing.T) {

		got := Search(dictionary, "math")
		want := "Couldn't find the word"

		assertString(t, got, want)
	})

}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
func TestSearchV2(t *testing.T) {
	dictionary := Dictionary{"test": "This is a type test"}
	t.Run("Known Word", func(t *testing.T) {

		got, _ := dictionary.SearchV2("test")
		want := "This is a type test"

		assertString(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {

		_, got := dictionary.SearchV2("unknown")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "hot"
		def := "having a high degree of heat or a high temperature"
		err := dictionary.Add(word, def)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "hot"
		def := "having a high degree of heat or a high temperature"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, def)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, def string) {
	t.Helper()

	got, err := dictionary.SearchV2(word)
	if err != nil {
		t.Fatalf("Should search for an added word: %s", err)
	}
	if def != got {
		t.Errorf("Got: %s, Want: %s", got, def)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.SearchV2(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}
