package chain_test

import (
	"testing"

	"github.com/LYH2263/go-auth-chain/internal/chain"
	"github.com/LYH2263/go-auth-chain/internal/model"
)

func TestMissingTokenFails(t *testing.T) {
	err := chain.Default().Authenticate(model.Request{})
	if err == nil {
		t.Fatal("want error")
	}
}

func TestOK(t *testing.T) {
	err := chain.Default().Authenticate(model.Request{Token: "t", Role: "admin"})
	if err != nil {
		t.Fatal(err)
	}
}
