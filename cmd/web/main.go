package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/dbdriver"
	"web3/pkg/handlers"
	"web3/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {

	// 25 Call the new run func
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	// Close DB when done
	defer db.SQL.Close()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// 25 I now need to pass the database and
// main is getting busy so I'll create a
// function that will handle sessions and
// connecting to the DB
func run() (*dbdriver.DB, error) {
	gob.Register(models.Article{})

	// 26 Add table models in session
	gob.Register(models.User{})
	gob.Register(models.Post{})

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = sessionManager

	// Connect to DB
	db, err := dbdriver.ConnectSQL("host=localhost port=5432 dbname=blog_db user=postgres password=TurtleDove")
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	// 25 Need to pass the DB pool
	repo := handlers.NewRepo(&app, db)

	handlers.NewHandlers(repo)

	// 28 Pass over AppConfig to render
	render.NewAppConfig(&app)

	return db, nil
}
