package main

import "testing"

func TestReplace1(t *testing.T) {
	str := "ls -l :1"
	values := []string{"main.go"}

	newStr := replacePlaceholders(str, values)

	if newStr != "ls -l main.go" {
		t.Fail()
	}
}
