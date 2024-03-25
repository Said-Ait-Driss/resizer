package service

import (
	"net/http"
	"resizer/internal/http/request"
)

func Fetcher(reqBody request.RequestBody) (*http.Response, error) {
	resp, err := http.Get(reqBody.ImageURL)
	return resp, err
}
