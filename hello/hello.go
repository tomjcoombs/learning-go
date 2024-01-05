package main 

import(
	"fmt"
	"log"
	"example.com/greetings"
) 


func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names 
	// names := []string{"timmy", "tiny", "tony", ""}

	// messages, err := greetings.Hellos(names)
	message, err := greetings.Hello("Tommy")

	// if there are no errors
	if err != nil {
		log.Fatal(err)
	}

	// otherwise print the message
	fmt.Println(message)
}