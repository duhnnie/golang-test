package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
  // the log entry prefix and a flag to disable printing
  // the time, source file, and line number.
  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  // A slice of names
  names := []string {
  	"Monica",
  	"Ramona",
  	"Chusma",
  	"Cooper",
  	"David",
  }

  // Request greeting messages for the names.
  messages, err := greetings.Hellos(names)

  // Checking for errors
  if err != nil {
		log.Fatal(err)  	
  }

  // Iterating through the map, and printing its values.
  for _, message := range messages {
  	fmt.Println(message)
  }
}
