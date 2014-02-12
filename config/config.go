package config

import (


)

type Config struct {

	elements map[string]interface{}

}

func Set(key string, defaultValue interface{}) {

	// initialize set methods as needed
	switch defaultValue.(type) {

	case int:
		fmt.Println()
	case nil:
		fmt.Println()

	}
}

func Get(key string) (interface{}, error) {

	// get the element as needed
	
	return 55, nil
}


