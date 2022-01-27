package main

import (
	"fmt"
	"log"

	"github.com/vmerlyn/GoLang/greetings"
)

const (
	PI = 3.14
)

func main() {
	log.SetPrefix("greetngs:")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
