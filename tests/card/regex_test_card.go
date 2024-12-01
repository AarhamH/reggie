package tests

import (
	"fmt"
	"testing"
	graph "tinyreg/graph"
	parser "tinyreg/parser"
)

func TestCreditCardNumber(t *testing.T) {
	var data = []struct {
		card     string
		validity bool
	}{
		// valid credit card numbers
		{card: "4123456789012343", validity: true},
		{card: "4234567890123452", validity: true},
		{card: "4345678901234561", validity: true},
		{card: "4901234567890123", validity: true},

		// card numbers that will fail
		{card: "4123456789012343", validity: true},
		{card: "4234567890123452", validity: true},
		{card: "4345678901234561", validity: true},
		{card: "4901234567890123", validity: true},
	}

	ctx := parser.Parse(`4[0-9]{12}[0-9]{3}`)
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
