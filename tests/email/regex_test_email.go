package tests

import (
	"fmt"
	"testing"
	graph "tinyreg/graph"
	parser "tinyreg/parser"
)

func TestEmails(t *testing.T) {
	var data = []struct {
		email    string
		validity bool
	}{
		// valid emails
		{email: "alex.smith123@company.com", validity: true},
		{email: "jane_doe@business.org", validity: true},
		{email: "mary.jane@tech.io", validity: true},
		{email: "info@service.net", validity: true},
		{email: "helpdesk@solutions.biz", validity: true},
		{email: "support@store.store", validity: true},
		{email: "contact@brand.shop", validity: true},
		{email: "user1234@domain.co", validity: true},
		{email: "feedback@website.biz", validity: true},
		{email: "team@project.team", validity: true},

		// invalid emails
		{email: "john.doe@.example.com", validity: false},
		{email: "test@domain..com", validity: false},
		{email: "user@domain..co", validity: false},
		{email: "user@sub..domain.com", validity: false},
		{email: "user@sub-domain@domain.com", validity: false},
		{email: "user@-subdomain.com", validity: false},
		{email: "user@subdomain-.com", validity: false},
		{email: ".username@domain.com", validity: false},
		{email: "user@domain@domain.com", validity: false},
		{email: "user@domain.com.", validity: false},
		{email: "user@domain.#com", validity: false},
		{email: "user@sub_domain.com", validity: false},
		{email: "user@domain/com", validity: false},
		{email: "user@domain..co.uk", validity: false},
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
