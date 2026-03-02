package fixedwindow

import (
	"log/slog"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/ehsundar/interview.git/ratelimiter/internal/ratelimiter/config"
)

type middleware struct {
	next http.Handler

	currentWindow int64
	requestCount  int64
	lock          sync.Mutex

	limit     int64
	window    int64
	pathMatch *regexp.Regexp
}

func NewMiddleware(next http.Handler, c config.Rule) (http.Handler, error) {
	pathMatch, err := regexp.Compile(c.PathMatch)
	if err != nil {
		return nil, err
	}

	return &middleware{
		next:      next,
		limit:     c.Limit,
		window:    c.FixedWindow.Window,
		pathMatch: pathMatch,
	}, nil
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matches := m.pathMatch.MatchString(r.URL.Path)
	if matches {
		allowed := m.allowRequest()
		if !allowed {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
	}

	m.next.ServeHTTP(w, r)
}

func (m *middleware) allowRequest() bool {
	now := time.Now().UnixMilli()
	currentWindow := now / m.window
	slog.Info("fixed window limiter",
		"now", now,
		"window", m.window,
		"currentWindow", currentWindow,
		"requestCount", m.requestCount,
		"limit", m.limit,
	)

	m.lock.Lock()
	defer m.lock.Unlock()

	if currentWindow == m.currentWindow {
		if m.requestCount < m.limit {
			m.requestCount++
			return true
		}

		return false
	}

	m.currentWindow = currentWindow
	m.requestCount = 1
	return true
}
