package main_test

import (
	"testing"

	go_specs_greet "github.com/Popoliku/go-specs-greet"
)

func TestGreeterServer(t *testing.T) {
	go_specs_greet.GreetSpecification(t, nil)
}
