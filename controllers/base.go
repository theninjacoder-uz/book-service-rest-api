package controllers

import (
	"fmt"
	"log"
	"net/http"
	"task/models"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbHost, DbUser, DbPassword, DbName, DbPort string) {
	var err error

	switch Dbdriver {
	case "mysql":
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	case "postgres":
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	default:
		fmt.Printf("Invalid %s database driver", Dbdriver)
		return
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Book{})
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
