package ratelimiter

import (
	"net/http"
)

type writerProxy struct {
	w     http.ResponseWriter
	dirty bool
}

func (wp *writerProxy) Header() http.Header {
	return wp.w.Header()
}

func (wp *writerProxy) Write(b []byte) (int, error) {
	wp.dirty = true
	return wp.w.Write(b)
}

func (wp *writerProxy) WriteHeader(statusCode int) {
	wp.dirty = true
	wp.w.WriteHeader(statusCode)
}
