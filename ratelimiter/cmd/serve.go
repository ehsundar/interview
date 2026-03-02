package main

import (
	"fmt"
	"net/http"

	"github.com/ehsundar/interview.git/ratelimiter/internal/namesapi"
	"github.com/ehsundar/interview.git/ratelimiter/internal/ratelimiter"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)

	server := namesapi.NewServer()

	mux := http.NewServeMux()
	mux.Handle("/", server)

	mwMux := ratelimiter.ApplyMiddleware(mux, config.Rules)

	err = http.ListenAndServe(":8080", mwMux)
	if err != nil {
		panic(err)
	}
}
