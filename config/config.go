package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	elements map[string]interface{}
}

// this is the variable that will hold all of our elements as needed 
var internal * Config 

// private methods
func New() {
	if internal != nil {
		return
	}
	// now initailize our configuration container
	internal = &Config{elements: map[string]interface{}{}}
}

// getters 
func First(args ...interface{}) interface{} {
	return args[0]
}

func Value(key string) interface{} {

	value, err := Get(key)

	if err != nil {
		return err
	} else {

		return value
	}
}

func Get(key string) (interface{}, error) {

	value, ok := internal.elements[key]

	if ok {

		return value, nil

	} else {

		return nil, errors.New("invalid key")
	}
}

// setters 
func Set(key string, value interface{}) error {
	// check if the key exists in the map already
	_, found := internal.elements[key]
	if found {
		return errors.New("Duplicate key.")
	} else {
		internal.elements[key] = value
		return nil
	}
}

func Bootstrap(keys []string) error {

	var value string

	for i := range keys {
		
		value = os.Getenv(keys[i])

		if value != "" {
				
			// attempt to convert to number ... 
			intValue, ok := strconv.Atoi(value)

			// we got a valid integer - lets save it
			if ok == nil {

				internal.elements[keys[i]] = intValue	

			} else { // no integer we can safely store - use the string value

				internal.elements[keys[i]] = value
			}

		} else { // key wasn't found. Lets find a way to pass it back as an error message

			return errors.New(keys[i])
		}
	}

	return nil
}

