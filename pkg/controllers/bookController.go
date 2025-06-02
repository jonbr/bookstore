package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jonbr/bookstore/pkg/models"
	"github.com/jonbr/bookstore/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// parse input data
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)

	// db action
	createBook.CreateBook()

	// return response to client
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createBook)
	fmt.Println("controller.CreateBook")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("controller.GetBook")
}
