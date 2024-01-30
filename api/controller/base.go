package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/bitmattz/go-task-manager-server/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("hot=%s port=%s user=%s dbname=%s sslmode=require password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)

	if err == nil {
		fmt.Printf("We are connected to the %s database", Dbdriver)
		server.DB.Debug().AutoMigrate(&model.Project{})
		server.Router = mux.NewRouter()
		server.initializeRoutes()
		return
	}
	fmt.Printf("Cannot conect to %s database", Dbdriver)
	log.Fatal("This is the error:", err)
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
