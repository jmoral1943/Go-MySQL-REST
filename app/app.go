package app

import (
	"encoding/json"
	"strconv"
	"database/sql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jmoral1943/musicapi/model"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
)

// App has a pointer to the mux router and the mysql db
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize initializes the app to start the connection with the MySQL database
func (a *App) Initialize(connection string) {

	var err error
	a.DB, err = sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

// Run listens and server based on the port given to it 
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := model.GetSongs(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, songs)
}

func (a *App) getSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	s := model.Song{ID: id}
	if err := s.GetSong(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Song not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, s)
}


func (a *App) createSong(w http.ResponseWriter, r *http.Request) {
	var s model.Song
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	id, err := s.CreateSong(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	s.ID = strconv.FormatInt(id, 10)
	respondWithJSON(w, http.StatusCreated, s)
}

func (a *App) updateSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]

	var s model.Song
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	s.ID = id;
	if err := s.UpdateSong(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := s.GetSong(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Song not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, s)
}
func (a *App) deleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]

	s := model.Song{ID: id}

	if err := s.DeleteSong(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Song not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, s)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/songs", a.getSongs).Methods("GET")
	a.Router.HandleFunc("/song/{id}", a.getSong).Methods("GET")
	a.Router.HandleFunc("/song", a.createSong).Methods("POST")
	a.Router.HandleFunc("/song/{id}", a.updateSong).Methods("PATCH")
	a.Router.HandleFunc("/song/{id}", a.deleteSong).Methods("DELETE")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "   ")

	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
