package http

import (
	"io/ioutil"
	"net/http"
)

// Service ...
type Service struct {
	client *http.Client
}

// GetInstance ...
func GetInstance() *Service {
	return &Service{
		client: &http.Client{},
	}
}

// GET ...
func (s *Service) GET(url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

// POST ...
func (s *Service) POST(url, msg string) (string, error) {
	// TODO
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
