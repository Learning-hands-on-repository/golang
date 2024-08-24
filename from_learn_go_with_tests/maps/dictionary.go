package maps

import "errors"

type Dictionary map[string]string

var ErrNotFoundDefinition = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("cannot add word because it already exists")

func (dictionary Dictionary) Search(keyword string) (string, error) {
	definition, ok := dictionary[keyword]
	if !ok {
		return "", ErrNotFoundDefinition
	}

	return definition, nil
}

func (dictionary Dictionary) Add(keyword, value string) error {
	
	_, ok := dictionary[keyword]
	if ok {
		return ErrWordExists
	}

	dictionary[keyword] = value
	return nil
}
