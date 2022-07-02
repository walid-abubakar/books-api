package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/walid-abubakar/books-api/pkg/mocks"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

  	// Iterate over all the mock books
	for index, book := range mocks.Books {
		if book.Id == id {
      mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:] ... )
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}