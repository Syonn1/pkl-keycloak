package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) login(payload *KLoginPayload) (*KLoginRespon, error) {
	
	formData := url.Values{
		"client_id": {payload.clientId},
		"client_secret": {payload.clientSecret},
		"grant_type": {payload.grantType},
		"username": {payload.username},
		"password": {payload.password},
	}

	encodeFormData := formData.Encode()

	req, err := http.NewRequest("POST", "http://localhost:8080/realms/pkl-realm/protocol/openid-connect/token", strings.NewReader(encodeFormData))
	
	if err != nil{
		return nil, err	
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)

	if err != nil{
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Something went wrong")
	}

	kLoginRespon := &KLoginRespon{}

	json.NewDecoder(resp.Body).Decode(kLoginRespon)

	return kLoginRespon, nil
}