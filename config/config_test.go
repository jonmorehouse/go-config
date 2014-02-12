package config

import (

	"testing"
	"fmt"
)

func TestSet(t *testing.T) {

	// now lets test the config attribute
	result := Set("HOST", "go-test-host")
	
	if result != nil {

		fmt.Println("FIAL")

	} else {


		fmt.Println(internal.elements["HOST"])

	}
	




}



