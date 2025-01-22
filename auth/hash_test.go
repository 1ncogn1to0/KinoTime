package auth_test

import (
	"PracticeServer/auth"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		expectError bool
		compareWith string
		shouldMatch bool
	}{
		{
			name:        "Valid password hash and match",
			password:    "mysecretpassword",
			expectError: false,
			compareWith: "mysecretpassword",
			shouldMatch: true,
		},
		{
			name:        "Empty password should hash",
			password:    "",
			expectError: false,
			compareWith: "",
			shouldMatch: true,
		},
		{
			name:        "Different password should not match",
			password:    "securepassword",
			expectError: false,
			compareWith: "wrongpassword",
			shouldMatch: false,
		},
		{
			name:        "Long password",
			password:    "a_very_long_password_that_exceeds_normal_length_and_should_still_work_fine",
			expectError: false,
			compareWith: "a_very_long_password_that_exceeds_normal_length_and_should_still_work_fine",
			shouldMatch: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hashedPassword, err := auth.HashPassword(test.password)

			if (err != nil) != test.expectError {
				t.Fatalf("Unexpected error status: got %v, expected error? %v", err, test.expectError)
			}

			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(test.compareWith))
			if test.shouldMatch && err != nil {
				t.Errorf("Expected password to match but it didn't: %v", err)
			} else if !test.shouldMatch && err == nil {
				t.Errorf("Expected password NOT to match but it did")
			}
		})
	}
}
