package db

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/liontv/url-shortener/utils"
	_ "modernc.org/sqlite"
)

type Database struct {
	connection *sql.DB
}

// Stellt eine Verbindung zu einer Datenbank her.
func NewDatabase(dbPath string) *Database {
	// checkt ob Datenbank existiert, wenn nicht wird sie erstellt
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Cannot get user home directory: ", err)
		return nil
	}

	if _, err := os.Stat(homePath + dbPath); err != nil {
		os.Mkdir(homePath+dbPath[:15], os.ModePerm)
		os.Create(homePath + dbPath)
	}

	// stellt Verbindung zur Datenbank her
	connection, err := sql.Open("sqlite", homePath+dbPath)
	if err != nil {
		log.Fatal("Cannot open database: ", err)
		return nil
	}
	return &Database{connection}
}

// Schließt Datenbankverbindung
func (db *Database) Close() {
	db.connection.Close()
	println("Connection closed.")
}

// Checkt ob Datenbankverbindung besteht
func (db *Database) IsAlive() bool {
	if db.connection.Ping() != nil {
		return false
	}
	return true
}

func (db *Database) Query(query string, args ...any) (*sql.Rows, error) {
	utils.Log(query)
	return db.connection.Query(query, args...)
}

func (db *Database) Exec(query string, args ...any) (sql.Result, error) {
	utils.Log(query)
	return db.connection.Exec(query, args...)
}

// Gibt den clickcount des shorts zurück
func (db *Database) GetClicks(short string) int {
	var clicks int = 0
	res, _ := db.Query("SELECT clicks FROM urls WHERE short = '" + short + "' LIMIT 1")

	if res.Next() == false {
		utils.Log("Der short >>" + short + "<< existiert nicht. (GetClicks)")
		return -1
	}
	res.Scan(&clicks)
	res.Close()
	return clicks
}

// Gibt alle Einträge der Datenbank als json zurück
func (db *Database) GetAll() string {
	var id int
	var url string
	var short string
	var clicks int
	var json string = "{"

	res, _ := db.Query("SELECT * FROM urls")

	for res.Next() {
		res.Scan(&id, &url, &short, &clicks)
		json += "\"" + short + "\":{\"url\":\"" + url + "\",\"clicks\":" + strconv.Itoa(clicks) + "},"
	}
	res.Close()

	if json != "{" {
		json = json[:len(json)-1]
	}
	json += "}"
	return json
}
