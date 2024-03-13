package utils

import (
	"math/rand"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Gibt einen zufälligen Wert zwischen min und max zurück
func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// Generiert einen string mit beliebiger Länge (limit)
func GenerateShort(limit int) string {
	var short string
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < limit; i++ {
		short += string(letters[RandomInt(0, len(letters))])
	}
	return short
}

func Log(message string) {
	println(">>LOG<< [ " + time.Now().Format("2006-01-02 15:04:05") + " ] : " + message)
}
