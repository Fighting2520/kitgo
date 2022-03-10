package jwttoken

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("123456", "jkwang", "123456", 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestParseUsernameFromToken(t *testing.T) {
	token, err := GenerateToken("123456", "jkwang", "123456", 5)
	if err != nil {
		t.Fatal(err)
	}
	username, password, err := ParseUsernameAndPasswordFromToken(token, "123456")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(username, password)
}
