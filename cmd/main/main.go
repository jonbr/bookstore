package main

import (
	"github.com/gorilla/mux"
	"github.com/jonbr/bookstore/pkg/config"
)

func main() {
	a := &config.App{}

	// Initialize Database
	a.ConnectToDB("root:my-secret-pw@tcp(127.0.0.1:63890)/bookstore?parseTime=true")
	a.Migrate()

	// Initialize the router
	a.Router = mux.NewRouter().StrictSlash(true)
	a.InitializeRoutes()

	// Initialize logger
	a.InitializeLogger()

	//log.Println("Starting Server on port :9010")
}
