package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Get a greeting message and print it.
	// message, err := greetings.Hello("Mario")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
