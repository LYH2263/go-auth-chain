package auth_test

import (
	"testing"

	"github.com/LYH2263/go-auth-chain/internal/auth"
	"github.com/LYH2263/go-auth-chain/internal/model"
)

func TestRequireToken(t *testing.T) {
	if auth.RequireToken(model.Request{}) == nil {
		t.Fatal("want err")
	}
}
