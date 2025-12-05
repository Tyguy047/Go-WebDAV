package main

import (
	"log"
	"os"
)

func checkFolder() {
	path := "./data"

	// Check if the folder exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the folder with 0755 permissions (rwxr-xr-x)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
		log.Println("Created data directory")
	}
}

func checkForUser() {
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")

	if user == "" || pass == "" {
        log.Fatal("No username or password set! Please set 'USERNAME' and 'PASSWORD' as environment variables or update them in the docker-config.yml! On some systems, Mac OS and most Linux distros, 'USERNAME' is the username of the user profile you are logged into while running this binary. You may still need to run 'export USERNAME'.")
    }

	if user == "username" || pass == "password" {
		log.Println("WARNING: You are using the default username and/or password. For better security please edit your 'docker-config.yml' file!")
	} else {
		log.Println("Username and Password check complete!")
	}
}