package auth

import (
	"fmt"

	"github.com/LYH2263/go-auth-chain/internal/model"
)

type Handler func(model.Request) error

func RequireToken(r model.Request) error {
	if r.Token == "" {
		return fmt.Errorf("missing token")
	}
	return nil
}

func RequireRole(r model.Request) error {
	if r.Role == "" {
		return fmt.Errorf("missing role")
	}
	return nil
}
