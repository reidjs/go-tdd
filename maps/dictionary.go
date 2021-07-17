package main

const (
	ErrNotFound   = DictionaryErr("Could not find the word")
	ErrWordExists = DictionaryErr("This word already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	default:
		d[word] = definition
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	default:
		delete(d, word)
	}

	return nil
}
