package config

import (

	"testing"
	//"fmt"
	. "launchpad.net/gocheck"
	"os"
	"strconv"

)

// bootstrap / setup go check suite
func Test(t *testing.T) { TestingT(t) }

// set up suite struct as needed
type TestSuite struct{}
var _ = Suite(&TestSuite{})

func (s *TestSuite) TestSet(c *C) {

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// set the 
	Set("PORT", port)

	configPort, err := Get("PORT")

	c.Assert(err, IsNil)

	c.Assert(port, Equals, configPort.(int))
}

func (s *TestSuite) TestGet(c *C) {

	// how do we handle a non-existent config?
	value, err := Get("NULL_VALUE")
	c.Assert(value, IsNil)
	c.Assert(err, Not(Equals), nil)

	// make sure we can get a port successfully
}

func (s *TestSuite) TestBootstrap(c *C) {

	keys := []string{"PORT", "HOST"}

	err := Bootstrap(keys)

	c.Assert(err, IsNil)


}




