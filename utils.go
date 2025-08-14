package utils

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("your name is empty")
	}

	return fmt.Sprintf(randomMessages(), name), nil
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		ms, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = ms
	}

	return messages, nil

}

func randomMessages() string {

	messages := []string{
		"Hola %v , que seas bienvenido",
		"Hola %v , como estas?",
		"Hola %v , que tal?",
		"Hola %v , hasta que te dejas ver",
		"Hola %v , hace cuanto kilos no nos vemos?",
	}

	rndIndex := rand.Intn(len(messages))

	return messages[rndIndex]

}
