package mocks

import "github.com/walid-abubakar/books-api/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "All about Go!",
	},
}