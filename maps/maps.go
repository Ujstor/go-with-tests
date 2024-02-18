package maps

type Dictionary map[string]string

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

var (
	ErrNotFound   = DictionaryError("could not find word in dictionary")
	ErrWordExists = DictionaryError("could not add word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

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