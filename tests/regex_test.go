package tests

import (
	"fmt"
	"testing"
	graph "tinyreg/graph"
	parser "tinyreg/parser"
)

func TestAllCaps(t *testing.T) {
	var data = []struct {
		username string
		validity bool
	}{
		{username: "JOHN", validity: true},
		{username: "J", validity: true},

		// words that will fail
		{username: "john", validity: false},
		{username: "john", validity: false},
		{username: " ", validity: false},
		{username: "", validity: false},
		{username: "JOHN1", validity: false},
	}

	ctx := parser.Parse(`[A-Z]+`)
	graph := graph.ToGraph(ctx)

	for _, instance := range data {
		t.Run(fmt.Sprintf("Test: '%s'", instance.username), func(t *testing.T) {
			result := graph.Check(instance.username, -1)
			if result != instance.validity {
				t.Logf("Expected: %t, got: %t\n", instance.validity, result)
				t.Fail()
			}
		})
	}
}

func TestAlphabet(t *testing.T) {
	var data = []struct {
		username string
		validity bool
	}{
		// true if the word contains all alphabets or is an empty string
		{username: "aBcDeFgHiJkLmnoPqRstUvwXyZ", validity: true},
		{username: "A", validity: true},
		{username: "a", validity: true},
		{username: "", validity: true},

		// words that will fail
		{username: "apple1", validity: false},
		{username: "john_", validity: false},
		{username: " ", validity: false},
	}

	ctx := parser.Parse(`[A-Za-z]*`)
	graph := graph.ToGraph(ctx)

	for _, instance := range data {
		t.Run(fmt.Sprintf("Test: '%s'", instance.username), func(t *testing.T) {
			result := graph.Check(instance.username, -1)
			if result != instance.validity {
				t.Logf("Expected: %t, got: %t\n", instance.validity, result)
				t.Fail()
			}
		})
	}
}
