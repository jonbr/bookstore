package config

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jonbr/bookstore/pkg/models"
	"github.com/jonbr/bookstore/pkg/routes"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Log    *logrus.Logger
}

func (a *App) ConnectToDB(connectionString string) {
	var err error
	a.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func (a *App) Migrate() {
	if err := a.DB.AutoMigrate(
		&models.Book{},
	); err != nil {
		log.Fatal(err)
		panic("Database migration failed")
	}
}

func (a *App) InitializeRoutes() {
	routes.RegisterBookStoreRoutes(a.Router)
	http.Handle("/", a.Router)
	log.Fatal(http.ListenAndServe("localhost:9010", a.Router))
	log.Println("Starting Server on port :9010")
}

func (a *App) InitializeLogger() {
	a.Log = logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	a.Log.SetFormatter(&logrus.JSONFormatter{})
	//Log.SetFormatter(&logrus.TextFormatter{})
	/*if Environment == "PROD" {
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		Log.SetFormatter(&logrus.TextFormatter{})
	}*/

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	a.Log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	a.Log.SetLevel(logrus.InfoLevel)
}
