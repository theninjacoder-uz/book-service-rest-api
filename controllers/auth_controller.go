package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"task/models"
)

func (server *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		models.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		models.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("signup")
	if err != nil {
		models.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	data, err := user.SaveUser(server.DB)
	if err != nil {
		models.ErrorResponse(w, http.StatusConflict, err)
		return
	}
	res := models.Response{
		Message: "ok",
		IsOk:    true,
		Data:    data,
	}
	res.SuccessReponse(w, http.StatusCreated)
}
