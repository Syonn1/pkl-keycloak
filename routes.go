package main

import (
	"encoding/json"
	"net/http"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &GreetRes{
		Hello: "worlds",
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	payload := new(LoginPayload)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid payload"))
		return
	}

	kpayload := &KLoginPayload{
		clientId:     "pkl-keycloak",
		username:     payload.Username,
		password:     payload.Password,
		grantType:    "password",
		clientSecret: "BmZQ6JQbSxEdjq3sh1EyEKoPefetiTcm",
	}

	kres, err := s.client.login(kpayload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	res := &LoginRes{
		AccessToken: kres.AccessToken,
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
