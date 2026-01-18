package main

import (
	"log"
	"os"
)

func auth(triedUser string, triedPass string) bool {
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")
	// log.Println(user + " " + pass) // This line is used for debug

	if triedUser != user || triedPass != pass {
		log.Println("Failed login attempt!")
		return false
	}
	return true
}