package main

import "C"
import (
	"github.com/gin-gonic/gin"
	"library/pkg/handler"
	"library/pkg/service"
	"log"
)

var (
	s = service.New()
	h = handler.New(s)
)

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/books", h.ListBooks)
	router.GET("/books/{id}", h.GetBook)
	router.GET("/books", h.CreateBook)
	router.GET("/books/{id}", h.DeleteBook)
	router.GET("/books/{id}", h.UpdateBook)

	log.Fatal(router.Run())
}
