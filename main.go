package main

import (
	"log"
	"sandbox/ht"
)

func main() {
	hashtable, err := ht.New()
	if err != nil {
		log.Fatalf("Shit")
	}

	log.Println("We are done.")
}
