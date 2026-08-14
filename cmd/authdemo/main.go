package main

import (
	"fmt"

	"github.com/LYH2263/go-auth-chain/internal/chain"
	"github.com/LYH2263/go-auth-chain/internal/model"
)

func main() {
	err := chain.Default().Authenticate(model.Request{Token: ""})
	fmt.Println(err)
}
