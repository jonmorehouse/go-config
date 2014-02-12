package config

import (

	"testing"
	"reflect"
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

	// now pass in an unset varaible and ensure we pass an error back
	err = Bootstrap([]string{"NO_PORT"})

	c.Assert(err, Not(Equals), nil)
}

func (s *TestSuite) TestFirst(c *C) {

	// setup
	testPort := 4432
	Set("testPort", testPort)

	// now lets grab the first element
	result := First(Get("testPort"))

	// ensure that we successfully got the 
	c.Assert(testPort, Equals, result)
}

func (s *TestSuite) TestValue(c *C) {

	Set("TEST_PORT", 4444)

	// now lets get the value. Make sure we got the correct kind of result
	value := Value("TEST_PORT")
	c.Assert(value, Equals, 4444)
	c.Assert(reflect.TypeOf(value).Name(), Equals, "int")
}
