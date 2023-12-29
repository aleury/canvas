package model_test

import (
	"canvas/model"
	"testing"

	"github.com/matryer/is"
)

func TestEmail_IsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		address string
		valid   bool
	}{
		{"john.doe@example.com", true},
		{"jane_doe@example.co.uk", true},
		{"support@stackoverflow.com", true},
		{"user.name+tag+sorting@example.com", true},
		{"sales@mycompany123.com", true},
		{"contact_us@domain.io", true},
		{"feedback@subdomain.mywebsite.com", true},
		{"firstname.lastname@mydomain.net", true},
		{"custom123@my-personal-website.me", true},
		{"me@example.com", true},
		{"@example.com", false},
		{"me@.example.com", false},
		{"me@.com", false},
		{"me@", false},
		{"me", false},
		{"@", false},
		{"", false},
	}
	t.Run("reports valid email addresses", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.address, func(t *testing.T) {
				is := is.New(t)
				email := model.Email(test.address)
				is.Equal(test.valid, email.IsValid())
			})
		}
	})
}
