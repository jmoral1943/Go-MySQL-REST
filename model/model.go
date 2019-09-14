package model

import (
	"database/sql"
	"log"
	// "fmt"
)

// Song is the Json
type Song struct {
	ID      string `json`
	Name    string `json`
	Artist  string `json`
	Link    string `json`
	Genre   string `json`
	Album   string `json`
	Release string `json`
}

// GetSongs will get all the songs from the database
func GetSongs(db *sql.DB) ([]Song, error) {
	results, err := db.Query("SELECT * FROM Songs")
	if err != nil {
		panic(err.Error())
	}

	songs := []Song{}
	for results.Next() {
		var s Song

		err = results.Scan(&s.ID, &s.Name, &s.Artist, &s.Link, &s.Genre, &s.Album, &s.Release)

		if err != nil {
			panic(err.Error())
		}
		songs = append(songs, s)
		
	}
	return songs, nil
}

// GetSong queries one song from the database
func (s *Song) GetSong(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM Songs WHERE song_id=?", s.ID).Scan(&s.ID, &s.Name, &s.Artist, &s.Link, &s.Genre, &s.Album, &s.Release)
}

// CreateSong creates a new song in the Database
func (s *Song) CreateSong(db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO Songs( song_name, song_artist, song_link, song_genre, song_album, song_releasedate) VALUES( ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(s.Name, s.Artist, s.Link, s.Genre, s.Album, s.Release)
	if err != nil {
		log.Fatal(err)
	}
	
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}


	return lastID, nil
}

// UpdateSong updates a song
func (s *Song) UpdateSong(db *sql.DB) error {
	
	if s.Name != "" {
		stmt, err := db.Prepare("UPDATE Songs SET song_name=?  WHERE song_id=?")
		if err != nil {
			log.Fatal(err)
		} 
		_, err = stmt.Exec(s.Name, s.ID)
		if err != nil {
			log.Fatal(err)
		} 
	}
	if s.Artist != "" {
		stmt, err := db.Prepare("UPDATE Songs SET song_artist=?  WHERE song_id=?")
		if err != nil {
			log.Fatal(err)
		} 
		_, err = stmt.Exec(s.Artist, s.ID)
		if err != nil {
			log.Fatal(err)
		} 
	}
	if s.Link != "" {
		stmt, err := db.Prepare("UPDATE Songs SET song_link=?  WHERE song_id=?")
		if err != nil {
			log.Fatal(err)
		} 
		_, err = stmt.Exec(s.Link, s.ID)
		if err != nil {
			log.Fatal(err)
		} 
	}
	if s.Genre != "" {
		stmt, err := db.Prepare("UPDATE Songs SET song_Genre=?  WHERE song_id=?")
		if err != nil {
			log.Fatal(err)
		} 
		_, err = stmt.Exec(s.Genre, s.ID)
		if err != nil {
			log.Fatal(err)
		} 
	}
	if s.Album != "" {
		stmt, err := db.Prepare("UPDATE Songs SET song_album=?  WHERE song_id=?")
		if err != nil {
			log.Fatal(err)
		} 
		_, err = stmt.Exec(s.Album, s.ID)
		if err != nil {
			log.Fatal(err)
		} 
	}
	if s.Release != "" {
		stmt, err := db.Prepare("UPDATE Songs SET song_release=?  WHERE song_id=?")
		if err != nil {
			log.Fatal(err)
		} 
		_, err = stmt.Exec(s.Release, s.ID)
		if err != nil {
			log.Fatal(err)
		} 
	}
	
	return nil
}

// DeleteSong deletes a song
func (s *Song) DeleteSong(db *sql.DB) error {
	results := db.QueryRow("SELECT * FROM Songs WHERE song_id=?", s.ID).Scan(&s.ID, &s.Name, &s.Artist, &s.Link, &s.Genre, &s.Album, &s.Release)
	stmt, err := db.Prepare("DELETE FROM Songs WHERE song_id=?")
	if err != nil {
		log.Fatal(err)
	} 
	_, err = stmt.Exec(s.ID)
	if err != nil {
		log.Fatal(err)
	}

	return results
}