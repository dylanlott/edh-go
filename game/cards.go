package main

import "errors"

// CardList exposes a set of methods for manipulating a list of Cards
type CardList []Card

// Shuffle is a sugar method to make Shuffling a list of Cards easier.
func (c CardList) Shuffle() ([]Card, error) {
	return []Card{}, errors.New("not impl")
}

// TODO: Implement the go Sort interface on Cards here for sorting methods
func (c CardList) Sort() error {
	return errors.New("not impl")
}
