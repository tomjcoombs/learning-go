package greetings

import(
	"errors"
	"fmt"
	"math/rand"
) 

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message
	// message:= fmt.Sprintf(randomFormat(), name)
	message:= fmt.Sprintf(randomFormat())
	return message, nil
} 

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random

func randomFormat() string {
	// A slice of message formats
	formats:= []string{
		"yo, %v what's up",
		"hey my guy %v",
		"wag wan %v",
	}

	// return a randomly selected message format by specifying a random index
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error){
	messages:= make(map[string]string)
	for _, name:= range names{
		message, err := Hello(name)
		if err != nil{
			return nil, errors.New("empty name")
		}
		messages[name] = message
	}
	return messages, nil
}