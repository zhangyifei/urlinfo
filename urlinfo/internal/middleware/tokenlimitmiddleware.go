package middleware

import (
	"net/http"
	"urlinfo/urlinfo/internal/config"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type TokenlimitMiddleware struct {
	Tokenlimit *limit.TokenLimiter
}

const (
	burst   = 100
	rate    = 100
	seconds = 5
)

func newTokenLimiter(c config.Config) *limit.TokenLimiter {
	rdxKey := "urlinfo-rate"
	store := redis.New(c.Cache[0].RedisConf.Host)
	limit := limit.NewTokenLimiter(rate, burst, store, rdxKey)
	return limit
}

func NewTokenlimitMiddleware(c config.Config) *TokenlimitMiddleware {
	return &TokenlimitMiddleware{Tokenlimit: newTokenLimiter(c)}
}

func (m *TokenlimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if m.Tokenlimit.Allow() {
			next(w, r)
		} else {
			http.Error(w, "server is too busy", http.StatusTooManyRequests)
		}
	}
}
