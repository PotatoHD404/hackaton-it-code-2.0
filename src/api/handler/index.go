package function

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"net/http"
	"os"
)

// Book struct (Model)
type Book struct {
	bun.BaseModel `bun:"books"`
	ID            string `bun:"id,pk,notnull" json:"id"`
	Isbn          string `json:"isbn"`
	Title         string `json:"title"`
	Author        *User  `json:"author"`
}

// User struct
type User struct {
	bun.BaseModel `bun:"users"`
	ID            int64  `json:"id"`
	Name          string `bun:"name" json:"name"`
	Surname       string `bun:"surname" json:"surname"`
}

var handler http.Handler
var db *bun.DB


//var books []Book

func NewHttpHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	//r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	//r.HandleFunc("/api/books", CreateBook).Methods("POST")
	//r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	//r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")
	return r
}

//https://github.com/uptrace/bun/blob/master/example/fixture/main.go

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	var authors []User
	db.RegisterModel((*User)(nil))
	if err := db.NewSelect().Model(&authors).Scan(ctx); err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(authors); err != nil {
		panic(err)
	}
}

func NewDB() *bun.DB {
	dsn := os.Getenv("POSTGRESQL")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func InitProject() {
	//books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One",
	//	User: &User{name: "John", surname: "Doe", ID: 1}})
	//books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two",
	//	User: &User{name: "Steve", surname: "Smith", ID: 2}})
	handler = NewHttpHandler()
	db = NewDB()
}

//goland:noinspection GoUnusedExportedFunction
func Handler(w http.ResponseWriter, r *http.Request) {
	if handler == nil || db == nil {
		InitProject()
	}
	w.Header().Set("Content-Type", "application/json")
	handler.ServeHTTP(w, r)
}

// Init books var as a slice Book struct

//// GetBooks Get all books
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(books)
//}
//
//// GetBook Get single book
//func GetBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r) // Gets params
//	// Loop through books and find one with the id from the params
//	for _, item := range books {
//		if item.ID == params["id"] {
//			_ = json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	_ = json.NewEncoder(w).Encode(&Book{})
//}
//
//// CreateBook Add new book
//func CreateBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var book Book
//	_ = json.NewDecoder(r.Body).Decode(&book)
//	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
//	books = append(books, book)
//	_ = json.NewEncoder(w).Encode(book)
//}
//
//// UpdateBook Update book
//func UpdateBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for index, item := range books {
//		if item.ID == params["id"] {
//			books = append(books[:index], books[index+1:]...)
//			var book Book
//			_ = json.NewDecoder(r.Body).Decode(&book)
//			book.ID = params["id"]
//			books = append(books, book)
//			_ = json.NewEncoder(w).Encode(book)
//			return
//		}
//	}
//}
//
//// DeleteBook Delete book
//func DeleteBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for index, item := range books {
//		if item.ID == params["id"] {
//			books = append(books[:index], books[index+1:]...)
//			break
//		}
//	}
//	_ = json.NewEncoder(w).Encode(books)
//}

//Request sample
//{
//"isbn":"4545454",
//"title":"Book Three",
//"author":{"firstname":"Harry", "lastname":"White"}
//}
