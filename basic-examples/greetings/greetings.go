package greetings;

import  ( 
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name")
	}

	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf(getRandomMessage(), name);

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomMessage() string {
	strings := []string{
		"Hi %v, nice to see you",
		"What's up %v?",
		"Nice to meet you %v",
	}

	return strings[rand.Intn(len(strings))]
}