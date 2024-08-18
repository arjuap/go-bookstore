package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arjun/go-bookstore/pkg/models"
	"github.com/arjun/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBooks()
	res, _ := json.Marshal(NewBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookid"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookdetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newbook := &models.Book{}
	utils.ParseBody(r, newbook)
	b := newbook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	Id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatebook := &models.Book{}
	utils.ParseBody(r, updatebook)
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	ID, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")

	}
	bookdetails, _ := models.GetBookById(ID)
	if updatebook.Name != "" {
		bookdetails.Name = updatebook.Name
	}
	if updatebook.Author != "" {
		bookdetails.Author = updatebook.Author
	}
	if updatebook.Publication != "" {
		bookdetails.Publication = updatebook.Publication
	}
	bookdetails.UpdateBook()
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
