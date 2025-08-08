package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal(
			"PORT environment variable is not set",
		) 
	}

	fmt.Println("Server is running on port:", portString)
}
