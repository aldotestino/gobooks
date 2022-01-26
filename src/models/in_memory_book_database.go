package gobooks

import (
	"errors"
	"fmt"
)

type InMemoryBookDatabase struct {
	store []Book
}

func NewInMemoryBookDatabase() *InMemoryBookDatabase {
	return &InMemoryBookDatabase{store: []Book{}}
}

func (db *InMemoryBookDatabase) GetAllBooks() *[]Book {
	return &db.store
}

func (db *InMemoryBookDatabase) GetBook(ID int64) (*Book, error) {
	if ID > 0 && ID <= int64(len(db.store)) {
		return &db.store[ID-1], nil
	} else if ID > int64(len(db.store)) {
		return nil, errors.New(fmt.Sprintf("ID %d doesn't exist", ID))
	} else {
		return nil, errors.New("ID must be a positive integer")
	}
}

func (db *InMemoryBookDatabase) AddBook(title string, author string) *Book {
	var newBookID int64
	newBookID = int64(len(db.store) + 1)
	newBook := Book{
		ID:     newBookID,
		Title:  title,
		Author: author,
	}

	db.store = append(db.store, newBook)
	return &newBook
}
