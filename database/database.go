package database

import "main.go/model"

var books []model.Book

func Init() {
	books = append(books, model.Book{ID: "1", Isbn: "5656", Title: "Book one", Author: &model.Author{FirstName : "John", LastName : "Wick"}})
	books = append(books, model.Book{ID: "2", Isbn: "5657", Title: "Book one", Author: &model.Author{FirstName : "John", LastName : "Wick"}})
	books = append(books, model.Book{ID: "3", Isbn: "5658", Title: "Book one", Author: &model.Author{FirstName : "John", LastName : "Wick"}})
	books = append(books, model.Book{ID: "4", Isbn: "5659", Title: "Book one", Author: &model.Author{FirstName : "John", LastName : "Wick"}})
}

func GetBooks() (totalBooks []model.Book) {
	return books
}

func AppendBook(book model.Book) () {
	books = append(books, book)
}

func DeleteBook(id string) () {
	for index, item := range books {
		if id == item.ID {
			books = append(books[:index], books[index + 1:]...)
			break;
		}
	}
}