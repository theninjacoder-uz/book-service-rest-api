package controllers

import (
	"task/middlewares"
)

func (s *Server) initializeRoutes() {

	router := s.Router.Group("")

	router.POST("/signup", s.SignUp)
	router.GET("/myself", middlewares.SetMiddlewareAuthentication(s.DB), s.GetMe)
	router.GET("/books", middlewares.SetMiddlewareAuthentication(s.DB), s.GetAllBooks)
	router.POST("/books", middlewares.SetMiddlewareAuthentication(s.DB), s.CreateABook)
	router.PATCH("/books/:id", middlewares.SetMiddlewareAuthentication(s.DB), s.UpdateABook)
	router.DELETE("/books/:id", middlewares.SetMiddlewareAuthentication(s.DB), s.DeleteABook)

}
