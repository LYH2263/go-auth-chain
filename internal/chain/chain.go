package chain

import (
	"github.com/LYH2263/go-auth-chain/internal/auth"
	"github.com/LYH2263/go-auth-chain/internal/model"
)

type Chain struct {
	Handlers []auth.Handler
}

func Default() *Chain {
	return &Chain{Handlers: []auth.Handler{auth.RequireToken, auth.RequireRole}}
}

func (c *Chain) Authenticate(r model.Request) (err error) {
	for _, h := range c.Handlers {
		if err = h(r); err != nil { // 赋值给外层 err，失败即返回
			return err
		}
	}
	return err
}

