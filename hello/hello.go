package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Configure the logger, including the log entry prefix and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Igor")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// Greetings for multiple names
	names := []string{"Igor", "Lauren", "Naomi"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
