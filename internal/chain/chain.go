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
		if err := h(r); err != nil { // BUG: := 遮蔽外层 err
			// 本意返回，却只打日志风格忽略
			_ = err
		}
	}
	return err // 恒为 nil
}

