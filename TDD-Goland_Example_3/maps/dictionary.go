package maps

//Search function searches the dictionary for a specific key, This was my original implementation of the function
func Search(dictionary map[string]string, word string) string {
	for k, v := range dictionary {
		if k == word {
			return v
		} else {
			return "Couldn't find the word"
		}
	}

	return dictionary[word]
}

// Textbook Implementation of the function
type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) SearchV2(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Add function adds a word and definition to an existing Dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.SearchV2(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil

}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.SearchV2(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
