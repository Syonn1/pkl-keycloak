package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr   string
	client *Client
}

func NewAPIServer(addr string) *APIServer {
	httpClient := &http.Client{}
	return &APIServer{
		addr: addr,
		client: &Client{
			httpClient: httpClient,
		},
	}
}

func (s *APIServer) Run() {
	// s.migrate()

	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet)
	router.HandleFunc("/login", s.handleLogin)

	fmt.Printf("Server starting on address %s", s.addr)
	http.ListenAndServe(s.addr, router)
}
