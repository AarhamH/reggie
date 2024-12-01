package tests

import (
	"fmt"
	"testing"
	graph "reggie/graph"
	parser "reggie/parser"
)

func TestUsername(t *testing.T) {
	var data = []struct {
		username string
		validity bool
	}{
		// valid usernames (letters, numbers, and underscores)
		{username: "user_name123", validity: true},
		{username: "john_doe", validity: true},
		{username: "testUser", validity: true},
		{username: "valid_123", validity: true},
		{username: "user1", validity: true},
		{username: "_underscore", validity: true},
		{username: "1st_user", validity: true},
		// invalid usernames
		{username: "user@name", validity: false},
		{username: "123!abc", validity: false},
		{username: "user--name", validity: false},
		{username: "user name", validity: false},
		{username: "-username", validity: false},
	}

	ctx := parser.Parse(`[a-zA-Z0-9_]+`)
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
