package slidingwindow

import (
	"net/http"

	"github.com/ehsundar/interview.git/ratelimiter/internal/ratelimiter/config"
)

type middleware struct {
	next http.Handler
	rule config.Rule
}

func NewMiddleware(next http.Handler, rule config.Rule) (http.Handler, error) {
	return &middleware{
		next: next,
		rule: rule,
	}, nil
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.next.ServeHTTP(w, r)
}
