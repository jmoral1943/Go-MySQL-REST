package main

import (
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoral1943/musicapi/app"
)

func main() {

	a := app.App{}
	a.Initialize(os.Getenv("MusicAPIDB"))

	a.Run(":8080")
}