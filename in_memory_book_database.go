package gobooks

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

type InMemoryBookDatabase struct {
	store []Book
}

func NewInMemoryBookDatabase() *InMemoryBookDatabase {
	return &InMemoryBookDatabase{store: []Book{}}
}

func (db *InMemoryBookDatabase) GetAllBooks() *[]Book {
	return &db.store
}

func (db *InMemoryBookDatabase) GetBook(ID int) (*Book, error) {
	if ID > 0 && ID <= len(db.store) {
		return &db.store[ID-1], nil
	} else if ID > len(db.store) {
		return &Book{}, errors.New(fmt.Sprintf("id %d doesn't exist", ID))
	} else {
		return &Book{}, errors.New("ID must be a positive integer")
	}
}

func (db *InMemoryBookDatabase) AddBook(title string, author string) *Book {
	newBookId := len(db.store) + 1
	newBook := Book{
		ID:     newBookId,
		Title:  title,
		Author: author,
	}

	db.store = append(db.store, newBook)
	return &newBook
}
