package gobooks

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type BookDatabse interface {
	GetAllBooks() *[]Book
	GetBook(ID int) (*Book, error)
	AddBook(title string, author string) *Book
}

type BooksServer struct {
	BookDatabse
	http.Handler
}

func NewBooksServer(bookDatabase BookDatabse) *BooksServer {
	b := new(BooksServer)

	b.BookDatabse = bookDatabase

	router := http.NewServeMux()
	router.Handle("/books", http.HandlerFunc(b.booksHandler))
	router.Handle("/book/", http.HandlerFunc(b.getBookHandler))

	b.Handler = router
	return b
}

func (b *BooksServer) booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		b.getAllBooksHandler(w, r)
	case http.MethodPost:
		b.addBookHandler(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (b *BooksServer) getAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	sendJson(w, b.BookDatabse.GetAllBooks())
}

func (b *BooksServer) getBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookIDString := r.URL.Path[len("/book/"):]
	bookID, err := strconv.Atoi(bookIDString)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := b.BookDatabse.GetBook(bookID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sendJson(w, book)
}

type BookAddRequest struct {
	Title  string
	Author string
}

func (b *BooksServer) addBookHandler(w http.ResponseWriter, r *http.Request) {
	var br BookAddRequest

	if err := json.NewDecoder(r.Body).Decode(&br); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if br.Title == "" || br.Author == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	book := b.BookDatabse.AddBook(br.Title, br.Author)

	sendJson(w, book)
}

func sendJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}
