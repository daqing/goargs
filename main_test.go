package main

import (
	"strings"
	"testing"
)

func TestReplace1(t *testing.T) {
	str := "ls -l :1"

	result := strings.Join(
		replacePlaceholders(str, "main.go"),
		" ")

	if result != "ls -l main.go" {
		t.Errorf("Expected: %s, got: %s", "ls -l main.go", result)
	}
}

func TestReplace2(t *testing.T) {
	str := "mv :1 :1.bak"

	result := strings.Join(
		replacePlaceholders(str, "main.go"),
		" ")

	if result != "mv main.go main.go.bak" {
		t.Fail()
	}
}

func TestReplace3(t *testing.T) {
	str := "mv :2 :1.bak"

	result := strings.Join(
		replacePlaceholders(str, "a b"),
		" ")

	if result != "mv b a.bak" {
		t.Fail()
	}
}

func TestReplace4(t *testing.T) {
	str := "mv :0 :2.bak"

	result := strings.Join(
		replacePlaceholders(str, "Frame 123.svg"),
		" ")

	expected := "mv Frame 123.svg 123.svg.bak"

	if result != expected {
		t.Errorf("Expected: %s, got: %s", expected, result)
	}
}
