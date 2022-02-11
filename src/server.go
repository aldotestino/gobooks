package gobooks

import (
	models "app/src/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

type BookDatabase interface {
	GetAllBooks() *[]models.Book
	GetBook(ID int64) (*models.Book, error)
	AddBook(title string, author string) *models.Book
}

type BooksServer struct {
	BookDatabase
	http.Handler
}

func NewBooksServer(bookDatabase BookDatabase) *BooksServer {
	b := new(BooksServer)

	b.BookDatabase = bookDatabase

	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./public")))
	router.Handle("/books", http.HandlerFunc(b.booksHandler))
	router.Handle("/book/", http.HandlerFunc(b.getBookHandler))

	b.Handler = cors.AllowAll().Handler(router)
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
	sendJson(w, b.BookDatabase.GetAllBooks())
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

	book, err := b.BookDatabase.GetBook(int64(bookID))

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

	book := b.BookDatabase.AddBook(br.Title, br.Author)

	sendJson(w, book)
}

func sendJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}
