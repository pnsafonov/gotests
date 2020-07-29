package ttestify

import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "log"
    "testing"
)


// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
    suite.Suite
    VariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupSuite() {
    log.Println("SetupSuite")
}

func (suite *ExampleTestSuite) SetupTest() {
    log.Println("SetupTest")
    suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TearDownSuite() {
    log.Println("TearDownSuite")
}

func (suite *ExampleTestSuite) TearDownTest() {
    log.Println("TearDownTest")
}

func (suite *ExampleTestSuite) BeforeTest(suiteName, testName string) {
    log.Printf("BeforeTest, suiteName = %s, testName = %s", suiteName, testName)
}

func (suite *ExampleTestSuite) AfterTest(suiteName, testName string) {
    log.Printf("AfterTest, suiteName = %s, testName = %s", suiteName, testName)
}


// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample1() {
    log.Println("TestExample1")
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *ExampleTestSuite) TestExample2() {
    log.Println("TestExample2")
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *ExampleTestSuite) TestExample3() {
    log.Println("TestExample3")
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(ExampleTestSuite))
}