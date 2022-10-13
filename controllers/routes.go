package controllers

import (
	"task/middlewares"
)

func (s *Server) initializeRoutes() {

	// // Auth Route
	// s.Router.HandleFunc("/signup", middlewares.SetMiddlewareJSON())
	s.Router.HandleFunc("/signup", middlewares.SetMiddlewareJSON(s.SignUp)).Methods("POST")

}
