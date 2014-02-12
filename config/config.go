package config

import (

	"errors"
	"os"
	"strconv"
)

type Config struct {

	// map all configuration elements to the variables as needed
	elements map[string]interface{}
}

// this is the variable that will hold all of our elements as needed 
var internal Config 
var initialized bool

func bootstrap() {
	
	if initialized {

		return
	}

	// set our global variable as needed
	initialized = true

	// now initailize our configuration container
	internal = Config{elements: map[string]interface{}{}}
}


func Set(key string, defaultValue interface{}) error {

	bootstrap()

	// initialize a return error as needed
	var returnError error

	// first lets check and see if the key already exists
	_, ok := internal.elements[key] // grab the element 

	// if the value exists
	if ok {

		// the value already exists - lets not override it
		return errors.New("Key already exists")

	}

	// now lets handle actually grabbing the key per necessary
	stringValue := os.Getenv(key)

	// now lets determine the type of the defaultValue so we can cast as necessary
	switch defaultValue.(type) {

	case int:

		// convert the output to integer
		intValue, err := strconv.Atoi(stringValue) 

		// if we have an issue converting the string or we have a 0 value
		if err != nil || intValue == 0 {

			// return the error passed
			returnError = err

			// used the default value
			internal.elements[key] = defaultValue

		} else { // if everything works then lets set this environment variable

			internal.elements[key] = intValue
		}
		
	case int64:

		// convert the output to integer
		intValue, err := strconv.Atoi(stringValue) 

		// if we have an issue converting the string or we have a 0 value
		if err != nil || intValue == 0 {

			// return the error passed
			returnError = err

			// used the default value
			internal.elements[key] = defaultValue

		} else { // if everything works then lets set this environment variable

			internal.elements[key] = int64(intValue)
		}

	case string:

		// no value found
		if stringValue == "" {

			returnError = errors.New("Environment variable not found. Using default.")

			// use the default value as the config value
			internal.elements[key] = defaultValue

		} else {

			// now set the string variable as needed 
			internal.elements[key] = stringValue
		}

	default: // just go ahead and use the normal type - cast as the defaultValue's type 

		// would be nice to create a list dynamically as needed as well ...
		// not sure how to handle this -- errors could be everywhere with casting to the default value's type ... 

	}

	return returnError
}

func Get(key string) (interface{}, error) {

	bootstrap()

	value, ok := internal.elements[key]

	if ok {

		return value, nil

	} else {

		return nil, errors.New("invalid key")
	}
}

