package connect

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DB connects to MySQL with mysql driver
func DB() (db *sql.DB) {
	DATABASE := os.Getenv("MusicAPIDB")

	db, err := sql.Open("mysql", DATABASE)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return db

	// results, err := db.Query("SELECT * FROM Songs")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// for results.Next() {
	// 	var song song.Song

	// 	err = results.Scan(&song.ID, &song.Name, &song.Artist, &song.Link, &song.Genre, &song.Album, &song.Release)

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	data, err := json.MarshalIndent(song, "", "   ")
	// 	if err != nil {
	// 		log.Fatalf("JSON marshalling failed: %s", err)
	// 	}

	// 	fmt.Printf("%s\n", data)
	// }
}
