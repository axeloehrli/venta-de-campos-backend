package token

import (
	"testing"
	"time"

	"github.com/axeloehrli/venta-de-campos-backend/util"
	"github.com/golang-jwt/jwt/v4"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	if err != nil {
		t.Fatalf("there was an error")
	}

	username := util.RandomNombreUsuario()
	duration := time.Minute

	token, err := maker.CreateToken(username, duration)

	if err != nil {
		t.Fatalf("error creating token")
	}
	if token == "" {
		t.Fatalf("token is empty")
	}
	payload, err := maker.VerifyToken(token)

	if err != nil {
		t.Fatalf("error verifying token")
	}
	if (Payload{}) == *payload {
		t.Fatalf("empty payload")
	}
	if payload.Username != username {
		t.Fatal("different username")
	}
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	if err != nil {
		t.Fatalf("there was an error %v", err)
	}

	token, err := maker.CreateToken(util.RandomNombreUsuario(), -time.Minute)
	if err != nil {
		t.Fatalf("error creating token")
	}
	if token == "" {
		t.Fatalf("token is empty")
	}
	_, err = maker.VerifyToken(token)
	if err == nil {
		t.Fatalf("error verifying token: %v", err)
	}
	if err != ErrExpiredToken {
		t.Fatalf("unknown err: %v", err)
	}

}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomNombreUsuario(), time.Minute)
	if err != nil {
		t.Fatalf("err creating payload: %v", err)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		t.Fatalf("error signing token: %v", err)
	}
	maker, err := NewJWTMaker(util.RandomString(32))
	if err != nil {
		t.Fatalf("error jwtmaker: %v", err)
	}

	_, err = maker.VerifyToken(token)

	if err == nil {
		t.Fatalf("error should not be nil: %v", err)
	}

	if err != ErrInvalidToken {
		t.Fatalf("unknown err: %v", err)
	}
}
