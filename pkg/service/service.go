package service

import (
	"context"
	"library/pkg/repository/db"
)

type BookService interface {
	ListBooks() ([]db.Book, error)
	GetBook(id int64) (db.Book, error)
	CreateBook(db.CreateBookParams) (db.Book, error)
	DeleteBook(id int64) error
	UpdateBook(id int64) (db.Book, error)
}

type bookService struct {
	books []db.Book
}

func New() BookService {
	return &bookService{}
}

func (s *bookService) ListBooks() ([]db.Book, error){

}
func (s *bookService) GetBook(id int64) (db.Book, error){

}
func (s *bookService) CreateBook(params db.CreateBookParams) (db.Book, error) {

}

func (s *bookService) DeleteBook(id int64) error {

}
func (s *bookService) UpdateBook(id int64) (db.Book, error)
{

}
