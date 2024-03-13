package main

import (
	"github.com/liontv/url-shortener/db"
)

var database *db.Database

// Erstellt Datenbank und Verbindung zu dieser
func initialize() {
	database = db.NewDatabase("\\url-shortener\\database.db")
	database.Query("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY AUTOINCREMENT, url TEXT, short TEXT, clicks INTEGER);")
}

func main() {
	initialize()
	startWebServer()
}
