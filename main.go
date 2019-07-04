package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	port       = os.Getenv("PORT")
	host       = os.Getenv("SELF_HOST")
	targetHost = os.Getenv("TARGET_HOST")
)

func main() {
	var (
		target   = fmt.Sprintf("%s/subscription", targetHost)
		callback = fmt.Sprintf("%s:%s/notify", host, port)
	)
	if id, err := subscribe(target, callback); err != nil {
		log.Fatalf("Subscription failed at: %s. Error: %v\n", target, err)
	} else {
		log.Printf("Subscription created. ID: %s\n", id)
	}
}
