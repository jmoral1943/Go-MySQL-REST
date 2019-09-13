package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoral1943/musicapi/model"
)

// App has a pointer to the mux router and the mysql db
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize initializes the app to start the connection with the MySQL database
func (a *App) Initialize(connection string) {

	a.DB, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	defer a.DB.Close()

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

// Run listens and server based on the port given to it 
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getSongs(w http.ResponseWriter, r *http.Request) {
	results, err := a.DB.Query("SELECT * FROM Songs")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var song model.Song

		err = results.Scan(&song.ID, &song.Name, &song.Artist, &song.Link, &song.Genre, &song.Album, &song.Release)

		if err != nil {
			panic(err.Error())
		}

		data, err := json.MarshalIndent(song, "", "   ")
		if err != nil {
			log.Fatalf("JSON marshalling failed: %s", err)
		}

		fmt.Printf("%s\n", data)
	}
}

func (a *App) getSong(w http.ResponseWriter, r *http.Request) {

}

func (a *App) createSong(w http.ResponseWriter, r *http.Request) {

}

func (a *App) updateSong(w http.ResponseWriter, r *http.Request) {

}
func (a *App) deleteSong(w http.ResponseWriter, r *http.Request) {

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/songs", a.getSongs).Methods("GET")
	a.Router.HandleFunc("/song/{id}", a.getSongs).Methods("GET")
	a.Router.HandleFunc("/song", a.createSong).Methods("POST")
	a.Router.HandleFunc("/song/{id}", a.updateSong).Methods("PUT")
	a.Router.HandleFunc("/song/{id}", a.deleteSong).Methods("DELETE")
}

// results, err := db.Query("SELECT * FROM Songs")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	for results.Next() {
// 		var song song.Song

// 		err = results.Scan(&song.ID, &song.Name, &song.Artist, &song.Link, &song.Genre, &song.Album, &song.Release)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		data, err := json.MarshalIndent(song, "", "   ")
// 		if err != nil {
// 			log.Fatalf("JSON marshalling failed: %s", err)
// 		}

// 		fmt.Printf("%s\n", data)
// 	}
