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
func Set(key string, defaultValue interface{}) error {

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

