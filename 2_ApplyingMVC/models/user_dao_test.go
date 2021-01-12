package models

/*
- Golang DOES NOT have Asserts by design. There are no asserts because, as long as you start using Asserts, when one of
	the asserts fail, the next ones ARE NOT evaluated.
- With no Asserts, Go will provide more than one error in a test

- But this implies not a good code to maintain. That is way libraries appear --> "stretchr/testify/assert"
- This library allows to continue testing although an assert fails

- In order for Go to know that this is a test 2 things need to be declared:
	- The name of the class must finalize in "_test"
	- The name of the test function must start with "Test"

- HAVE AS MANY TEST CASES AS NUMBER OF RETURNS YOU HAVE

- 3 steps must be complied in order to do a proper test case:
	1. Initialization (Optional) --> Needed when some type of data or something is created in order to do the test
	2. Execution --> Call the function to test
	3. Validation --> Validate that the output of the method is the expected for each case
 */

import (
	"net/http"
	"github.com/stretchr/testify/assert" // LIBRARY FOR ASSERTS IN GO
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Execution
	user, err := GetUser(0)

	//Validation
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error for user_id=0")
	assert.Equal(t, http.StatusNotFound, err.StatusCode, "We were expecting 404 when user is not found")
}

func TestGetUserNoError(t *testing.T){
	//Execution
	user, err := GetUser(123)

	//Validation
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Pablo", user.FirstName)
	assert.EqualValues(t, "Nocedal", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}

