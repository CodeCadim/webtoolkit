package user_test

import (
	"testing"
	"webtoolkit/metier/user"
)

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name     string
		username string
		password string
		role     user.Role
		err      error
	}{
		{
			name:     "Correct username and password length",
			username: "username",
			password: "password",
			role:     user.Customer,
			err:      nil,
		},
		{
			name:     "Username too short ie less than 4 chars",
			username: "abc",
			password: "12345678",
			err:      user.ErrUsernameTooShort,
		},
		{
			name:     "Password too short ie less than 4 chars",
			username: "abcdefgh",
			password: "123",
			err:      user.ErrPasswordTooShort,
		},
		{
			name:     "Role undefined",
			username: "username",
			password: "password",
			role:     9999,
			err:      user.ErrUndefinedRole,
		},
	}

	for _, test := range testcases {
		_, err := user.New(test.username, test.password, test.role)
		if err != test.err {
			t.Fatalf("Waiting %v but got %v", test.err, err)
		}
	}
}