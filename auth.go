package main

import (
	"log"
	"os"
)

func auth(triedUser string, triedPass string) bool {
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")
	log.Printf("DEBUG: Expected user='%s', got='%s' | Expected pass='%s', got='%s'", user, triedUser, pass, triedPass)

	if triedUser != user || triedPass != pass {
		log.Println("Failed login attempt!")
		return false
	}
	return true
}