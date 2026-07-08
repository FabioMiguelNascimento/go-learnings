package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {

	if (name == "") {
		return "", errors.New("Empty name")
	}
	
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())
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

func randomFormat() string {
	formats := []string {
		"Ola, %v seja bem vindo!",
		"Bom te ver %v!",
		"Finalmente  %v! Nos conhecemos!",
	}

	return formats[rand.Intn(len(formats))]
}