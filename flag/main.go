package main

import "log"

func main() {

	flags, err := ParseFlags()
	if err != nil {
		log.Fatalf("cannot parse flags: %v", err)
	}
	log.Printf("flags: %+v", flags)
}
