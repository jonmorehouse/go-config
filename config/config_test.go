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

type TestSuite struct{}
var _ = Suite(&TestSuite{})

func (s *TestSuite) SetUpTest(c *C) {
	New()
}

func (s *TestSuite) TestSet(c *C) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	Set("PORT", port)
	configPort, err := Get("PORT")
	c.Assert(err, IsNil)
	c.Assert(port, Equals, configPort.(int))
}

func (s *TestSuite) TestManualSet(c *C) {
	value := []int{5,5,5,6}
	// set the value
	err := ManualSet("list", value)
	c.Assert(err, IsNil)
	// now grab the value as needed
	grabValue := Value("list")
	c.Assert(grabValue, Not(Equals), nil)
	// make sure we got the correct list back
	for i := range value {
		c.Assert(value[i], Equals, grabValue.([]int)[i])
	}
}

func (s *TestSuite) TestGet(c *C) {

	value, err := Get("NULL_VALUE")
	c.Assert(value, IsNil)
	c.Assert(err, Not(Equals), nil)
	// make sure we can get a port successfully
}

func (s *TestSuite) TestBootstrap(c *C) {

	keys := []string{"PORT", "HOST"}
	err := Bootstrap(keys)
	c.Assert(err, IsNil)
	err = Bootstrap([]string{"NO_PORT"})
	c.Assert(err, Not(Equals), nil)
}

func (s *TestSuite) TestFirst(c *C) {
	//setup individual test by setting a dudd value
	testPort := 4432
	Set("testPort", testPort)
	result := First(Get("testPort"))
	c.Assert(testPort, Equals, result)
}

func (s *TestSuite) TestValue(c *C) {
	Set("TEST_PORT", 4444)
	value := Value("TEST_PORT")
	c.Assert(value, Equals, 4444)
	c.Assert(reflect.TypeOf(value).Name(), Equals, "int")
}
