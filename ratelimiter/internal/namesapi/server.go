package namesapi

import (
	"net/http"
)

type server struct {
	mux *http.ServeMux
}

func NewServer() http.Handler {
	mux := http.NewServeMux()
	s := server{mux: mux}
	mux.HandleFunc("/names", s.GetNames)

	return server{
		mux: mux,
	}
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s server) GetNames(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("A Name"))
}
