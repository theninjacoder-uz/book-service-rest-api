package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"task/models"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateABook(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	// fmt.Printf("Request body: %v", body)
	if err != nil {
		models.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}
	book := models.Book{}
	err = json.Unmarshal(body, &book)

	if err != nil {
		models.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	//Validate book's value before saving
	err = book.Validate("create")
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	// Save book into db
	data, err := book.SaveABook(server.DB)
	if err != nil {
		models.ErrorResponse(c, http.StatusConflict, err)
		return
	}

	res := models.Response{
		Message: "ok",
		IsOk:    true,
		Data:    data,
	}
	res.SuccessReponse(c, http.StatusCreated)
}

// Update a book
func (server *Server) UpdateABook(c *gin.Context) {

	book := models.Book{}
	bookDto := models.BookDto{}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &bookDto)
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	book = bookDto.Book
	book.Status = bookDto.Status

	//Validate book's value before updating
	err = book.Validate("update")
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	//Extract book id from url path
	bookID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	//Update a book by id
	data, err := book.UpdateABook(server.DB, uint64(bookID))
	if err != nil {
		models.ErrorResponse(c, http.StatusNoContent, err)
		return
	}

	res := models.Response{
		Message: "ok",
		IsOk:    true,
		Data:    data,
	}
	res.SuccessReponse(c, http.StatusOK)
}

func (server *Server) GetAllBooks(c *gin.Context) {

	book := models.Book{}

	books, err := book.FindAllBooks(server.DB)

	if err != nil {
		models.ErrorResponse(c, http.StatusNoContent, err)
		return
	}

	res := models.Response{
		Message: "ok",
		IsOk:    true,
		Data:    books,
	}

	res.SuccessReponse(c, http.StatusOK)
}

func (server *Server) DeleteABook(c *gin.Context) {
	book := models.Book{}
	//Extract book id from url path
	bookID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	//delete a book by id
	_, e := book.DeleteABook(server.DB, uint64(bookID))
	if e != nil {
		models.ErrorResponse(c, http.StatusBadRequest, e)
		return
	}

	res := models.Response{
		Message: "ok",
		IsOk:    true,
		Data:    "Successfully deleted",
	}

	res.SuccessReponse(c, http.StatusOK)
}
