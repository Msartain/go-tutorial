package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat())
	return message, nil
}

// Hellos returns a map that connects/assocaites each name with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// fmt.Println("map", map)
	// create map to associate names with messages
	messages := make(map[string]string)
	//Loop through the received slice of names, calling the 
	// Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with 
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// returns one of a set of greeting messages at random
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// returns formats[random interger less than the length of formats]
	return formats[rand.Intn(len(formats))]
}
