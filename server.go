package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/liontv/url-shortener/utils"
)

// Erstellt eine kürzere URL für eine lange
func createShort(w http.ResponseWriter, r *http.Request) {
	var url string = r.URL.Query().Get("url")
	var short string = utils.GenerateShort(5)
	var id int
	var tmp_url string
	var tmp_short string
	var clicks int

	// checkt ob der short schon in der Datenbank vorhanden ist
	res, _ := database.Query("SELECT * FROM urls WHERE short='" + short + "' LIMIT 1")

	for res.Next() == true {
		utils.Log("Der Short >>" + short + "<< existiert bereits. Es wird ein neuer erstellt. (Create)")
		short = utils.GenerateShort(5)
		res, _ = database.Query("SELECT * FROM urls WHERE short='" + short + "' LIMIT 1")
	}

	res.Scan(&id, &tmp_url, &tmp_short, &clicks)
	res.Close()

	database.Query("INSERT INTO urls (url, short, clicks) VALUES ('" + url + "', '" + short + "', 0)")
	fmt.Fprintf(w, "%s", short)
}

// Kümmert sich um die Weiterleitung
func handleShort(w http.ResponseWriter, r *http.Request) {
	var short string = r.URL.Path[1:]
	var id int
	var url string
	var tmp_short string
	var clicks int

	// wenn im browser nur / aufgerufen wird, wird favicon.ico ignoriert
	if short == "favicon.ico" {
		return
	}

	// checkt ob der short in der Datenbank vorhanden ist
	res, _ := database.Query("SELECT * FROM urls WHERE short='" + short + "' LIMIT 1")

	if res.Next() == false {
		utils.Log("Der Short >>" + short + "<< konnte nicht gefunden werden. (Redirect)")
		return
	}
	res.Scan(&id, &url, &tmp_short, &clicks)
	res.Close()

	// updated den click counter
	_, err := database.Exec("UPDATE urls SET clicks = clicks + 1 WHERE id = " + strconv.Itoa(id))
	utils.CheckErr(err)

	utils.Log("Redirecting to: " + url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func deleteShort(w http.ResponseWriter, r *http.Request) {
	var short string = r.URL.Query().Get("short")

	// checkt ob short existiert
	res, _ := database.Query("SELECT * FROM urls WHERE short = '" + short + "' LIMIT 1")

	if res.Next() == false {
		utils.Log("Der short >>" + short + "<< existiert nicht. (Delete)")
		return
	}

	res.Close()
	database.Exec("DELETE FROM urls WHERE short = '" + short + "'")
	utils.Log("Short >>" + short + "<< wurde gelöscht.")
}

func handleGetClicks(w http.ResponseWriter, r *http.Request) {
	var short string = r.URL.Query().Get("short")
	var clicks int = database.GetClicks(short)
	fmt.Fprintf(w, "%d", clicks)
}

func handleGetAll(w http.ResponseWriter, r *http.Request) {
	var json string = database.GetAll()
	fmt.Fprintf(w, "%s", json)
}

func startWebServer() {
	// system arguments
	port := flag.Int("port", 8080, "Port auf dem der Server laufen soll")
	flag.Parse()

	// api routes
	http.HandleFunc("/", handleShort)
	http.HandleFunc("/api/clicks/", handleGetClicks)
	http.HandleFunc("/api/create/", createShort)
	http.HandleFunc("/api/delete/", deleteShort)
	http.HandleFunc("/api/all/", handleGetAll)

	// startet server
	println("Server started at port " + strconv.Itoa(*port) + ".")
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
