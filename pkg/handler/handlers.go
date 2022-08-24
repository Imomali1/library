package handler

import (
	"github.com/gin-gonic/gin"
	"library/pkg/repository/db"
	"library/pkg/service"
	"net/http"
	"strconv"
)

type Controller interface {
	ListBooks(ctx *gin.Context)
	GetBook(ctx *gin.Context)
	CreateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
}

type controller struct {
	service service.BookService
}

func New(s service.BookService) Controller {
	return &controller{
		service: s,
	}
}

func (h *controller) ListBooks(ctx *gin.Context) {
	books, err := h.service.ListBooks()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (h *controller) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	book, err := h.service.GetBook(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (h *controller) CreateBook(ctx *gin.Context) {
	var input db.CreateBookParams
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	book, err := h.service.CreateBook(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"book created": book})
}
func (h *controller) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = h.service.DeleteBook(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "deleted"})
}
func (h *controller) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	book, err := h.service.UpdateBook(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"book updated": book})
}
