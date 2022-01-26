package gobooks

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const dsn = "postgres://postgres:postgres@db:5432/gobooksdb?sslmode=disable"

type PGBookDatabase struct {
	store *bun.DB
}

func NewPGBookDatabse() *PGBookDatabase {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	return &PGBookDatabase{store: db}
}

func (db *PGBookDatabase) Migrate() {
	_, err := db.store.NewCreateTable().Model((*Book)(nil)).Exec(context.Background())

	if err != nil {
		panic(err)
	}
}

func (db *PGBookDatabase) GetAllBooks() *[]Book {
	var books []Book
	db.store.NewSelect().Table("books").Scan(context.Background(), &books)
	if books == nil {
		return &[]Book{}
	}
	return &books
}

func (db *PGBookDatabase) GetBook(ID int64) (*Book, error) {
	var book Book
	err := db.store.NewSelect().Table("books").Where("id = ?", ID).Scan(context.Background(), &book)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ID %d doesn't exist", ID))
	}
	return &book, nil
}

func (db *PGBookDatabase) AddBook(title string, author string) *Book {
	book := Book{
		Title:  title,
		Author: author,
	}
	_, err := db.store.NewInsert().Model(&book).Exec(context.Background())

	if err != nil {
		fmt.Printf("Something went wrong %+v", err)
	}

	return &book
}
