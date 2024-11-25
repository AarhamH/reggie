package tests

import (
	"fmt"
	"testing"
	graph "tinyreg/graph"
	parser "tinyreg/parser"
)

func TestNfa(t *testing.T) {

	var data = []struct {
		email    string
		validity bool
	}{
		{email: "valid_email@example.com", validity: true},
		{email: "john.doe@email.com", validity: true},
		{email: "user_name@email.org", validity: true},
		{email: "support@email.io", validity: true},
		{email: "contact@123.com", validity: true},
		{email: "sales@email.biz", validity: true},
		{email: "test_email@email.test", validity: true},
		{email: "random.email@email.xyz", validity: true},
		{email: "user@domain12345.com", validity: true},
		{email: "user@12345domain.com", validity: true},
		// invalid when compared against our regex
		{email: "alice.smith123@email.co.uk", validity: false},
		{email: "invalid.email@", validity: false},
		{email: ".invalid@email.com", validity: false},
		{email: "email@invalid..com", validity: false},
		{email: "user@-invalid.com", validity: false},
		{email: "user@invalid-.com", validity: false},
		{email: "user@in valid.com", validity: false},
		{email: "user@.com", validity: false},
		{email: "user@.co", validity: false},
		{email: "user@domain.c", validity: false},
		{email: "user@domain.1a", validity: false},
		{email: "user@domain.c0m", validity: false},
		{email: "user@domain..com", validity: false},
		{email: "user@.email.com", validity: false},
		{email: "user@emai.l.com", validity: false},
		{email: "user@e_mail.com", validity: false},
		{email: "user@e+mail.com", validity: false},
		{email: "user@e^mail.com", validity: false},
		{email: "user@e*mail.com", validity: false},
		{email: "user@e.mail.com", validity: false},
		{email: "user@e_mail.net", validity: false},
		{email: "user@sub.domain.com", validity: false},
		{email: "user@sub-domain.com", validity: false},
		{email: "user@sub.domain12345.com", validity: false},
		{email: "user@sub.domain-12345.com", validity: false},
		{email: "user@-sub.domain.com", validity: false},
		{email: "user@sub-.domain.com", validity: false},
		{email: "user@domain-.com", validity: false},
		{email: "user@sub.domain.c0m", validity: false},
		{email: "user@sub.domain.c", validity: false},
		{email: "user@sub.domain.1a", validity: false},
		{email: "user@sub.domain.c0m", validity: false},
		{email: "user@sub.domain..com", validity: false},
		{email: "user@sub.domain.c0m", validity: false},
		{email: "user@sub.domain..com", validity: false},
		{email: "user@sub.domain.c0m", validity: false},
	}

	ctx := parser.Parse(`[a-zA-Z][a-zA-Z0-9_.]+@[a-zA-Z0-9]+.[a-zA-Z]{2,}`)
	graph := graph.ToGraph(ctx)

	for _, instance := range data {
		t.Run(fmt.Sprintf("Test: '%s'", instance.email), func(t *testing.T) {
			result := graph.Check(instance.email, -1)
			if result != instance.validity {
				t.Logf("Expected: %t, got: %t\n", instance.validity, result)
				t.Fail()
			}
		})
	}
}
