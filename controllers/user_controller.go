package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"task/models"

	"github.com/gin-gonic/gin"
)

func (server *Server) SignUp(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		models.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		models.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("signup")
	if err != nil {
		models.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	data, err := user.SaveUser(server.DB)
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

func (server *Server) GetMe(c *gin.Context) {

	key := c.Request.Header.Get("Key")
	user := models.User{}

	data, err := user.GetUserInfo(server.DB, key)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	res := models.Response{
		Message: "ok",
		IsOk:    true,
		Data:    data,
	}
	res.SuccessReponse(c, http.StatusOK)
}

// clear all records
func (server *Server) CleanUp(c *gin.Context) {
	user := models.User{}
	book := models.Book{}
	err := user.DeleteAllUsers(server.DB)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	err = book.DeleteAllBooks(server.DB)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
}
