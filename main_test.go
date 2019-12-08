package main

import (
	"testing"
)

func TestCheckAge(t *testing.T) {
	personOverAge := Person{
		Age: 15,
	}
	personUnderAge := Person{
		Age: 15,
	}
	if !checkAge(personOverAge.Age) {
		t.Errorf("Error in checkAge")
	}
	if checkAge(personUnderAge.Age) {
		t.Errorf("Error in checkAge")
	}

}
