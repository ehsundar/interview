package ratelimiter

import (
	"log/slog"
	"net/http"

	"github.com/ehsundar/interview.git/ratelimiter/internal/ratelimiter/config"
	"github.com/ehsundar/interview.git/ratelimiter/internal/ratelimiter/fixedwindow"
	"github.com/ehsundar/interview.git/ratelimiter/internal/ratelimiter/slidingwindow"
)

func ApplyMiddleware(next http.Handler, rules map[string]config.Rule) http.HandlerFunc {
	mws := []http.Handler{}

	for ruleName, rule := range rules {
		var err error
		var middleware http.Handler

		switch rule.Type {
		case "fixed-window":
			middleware, err = fixedwindow.NewMiddleware(next, rule)
		case "sliding-window":
			middleware, err = slidingwindow.NewMiddleware(next, rule)
		default:
			slog.Warn("unknown rule type: %s", rule.Type)
			continue
		}

		if err != nil {
			slog.Error("failed to create middleware",
				"error", err,
				"rule_name", ruleName,
			)
			continue
		}

		mws = append(mws, middleware)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		wp := &writerProxy{w: w}

		for _, mw := range mws {
			mw.ServeHTTP(wp, r)
			if wp.dirty {
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}
