package config

import (

	"testing"
	"fmt"
	"os"
)

func TestBootstrap(t *testing.T) {

	// call the bootstrap function - this should set up the entire application
	bootstrap()


}

func TestSet(t *testing.T) {


	fmt.Println(os.Getenv("PORT"))



}



