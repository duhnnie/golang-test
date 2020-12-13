package greetings;

import  ( "fmt"
"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name")
	}

	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi %v, welcome!", name);

	return message, nil
}
