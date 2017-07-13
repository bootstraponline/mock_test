package main

import (
	"testing"
	"github.com/bootstraponline/mock_test/mocks"
)

func TestSayGreeting(t *testing.T) {
	mock := new(mocks.GreetInterface)
	mock.On("Greeting").Return("goodbye")

	SayGreeting(mock)

	mock.AssertExpectations(t)
}
