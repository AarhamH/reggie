package tests

import (
	"fmt"
	"testing"
	graph "reggie/graph"
	parser "reggie/parser"
)

func TestSIN(t *testing.T) {
	var data = []struct {
		card     string
		validity bool
	}{
		// valid credit card numbers
		{card: "123-12-1234", validity: true},
		{card: "322-32-4362", validity: true},
		{card: "222-11-0000", validity: true},
		{card: "111-32-1254", validity: true},

		{card: "1234-12-1234", validity: false},
		{card: "322-322-4362", validity: false},
		{card: "222-11-00000", validity: false},
		{card: "1", validity: false},
		{card: "test", validity: false},
		{card: "tes-te-test", validity: false},
	}

	ctx := parser.Parse(`[0-9]{3}-[0-9]{2}-[0-9]{4}`)
	graph := graph.ToGraph(ctx)

	for _, instance := range data {
		t.Run(fmt.Sprintf("Test: '%s'", instance.card), func(t *testing.T) {
			result := graph.Check(instance.card, -1)
			if result != instance.validity {
				t.Logf("Expected: %t, got: %t\n", instance.validity, result)
				t.Fail()
			}
		})
	}
}
