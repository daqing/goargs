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

func TestReplace2(t *testing.T) {
	str := "mv :1 :1.bak"
	values := []string{"main.go"}

	newStr := replacePlaceholders(str, values)

	if newStr != "mv main.go main.go.bak" {
		t.Fail()
	}
}

func TestReplace3(t *testing.T) {
	str := "mv :2 :1.bak"
	values := []string{"a", "b"}

	newStr := replacePlaceholders(str, values)

	if newStr != "mv b a.bak" {
		t.Fail()
	}
}
