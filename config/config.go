package config

import (

	"errors"
	"os"
)

type Config struct {

	// map all configuration elements to the variables as needed
	elements map[string]interface{}
}

// this is the variable that will hold all of our elements as needed 
var internal Config 

func bootstrap() {

	// bootstrap our configuration environment
	if internal == (Config{elements:}) {

		//internal = new(config) 
		internal.elements = make(map[string]interface{})
	}
}

func Set(key string, defaultValue interface{}) error {

	// lets check to make sure that we have a valid internal element?
		
	// initialize return variable
	var returnError error

	// grab the current element
	element := internal.elements[key]

	// now lets 
	if element == nil {

		// now lets get the environment variable
		value := os.Getenv(key)

		if value == "" && defaultValue == nil {

			returnError = errors.New("Environment variable not found. No default passed")

		} else if value == "" { // no variable found -- use default

			// set the environment variable
			internal.elements[key] = defaultValue

			// now create an error to be returned
			returnError = errors.New("Environment variable not found. Using default value")

		} else { // we have a variable here to use!

			internal.elements[key] = value
		}

	} else {

		// this variable already exists - print something?
		returnError = errors.New("Key exists")
	}

	// we passed - no need to return
	return returnError
}

func Get(key string) (interface{}, error) {

	// get the element as needed
	return 55, nil
}


