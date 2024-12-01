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
