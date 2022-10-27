package util

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hashedPassword1, err := HashPassword(password)
	if err != nil {
		t.Fatalf("ERROR HASHING PASSWORD: %v", err)
	}
	if hashedPassword1 == "" {
		t.Fatalf("HASHED PASSWORD IS EMPTY")
	}

	err = CheckPassword(password, hashedPassword1)
	if err != nil {
		t.Fatalf("PASSWORDS DON'T MATCH")
	}

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword1)

	if err != bcrypt.ErrMismatchedHashAndPassword {
		t.Fatalf("WRONG PASSWORD SHOULD RETURN ERROR")
	}

	hashedPassword2, err := HashPassword(password)
	if err != nil {
		t.Fatalf("ERROR HASHING PASSWORD: %v", err)
	}
	if hashedPassword2 == "" {
		t.Fatalf("HASHED PASSWORD IS EMPTY")
	}
	if hashedPassword2 == hashedPassword1 {
		t.Fatalf("SAME PASSWORD SHOULD RETURN DIFFERENT HASH")
	}
}
