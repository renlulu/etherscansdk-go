package internal

import "github.com/go-resty/resty/v2"

type Base struct {
	API    string
	APIKey string
	Client *resty.Client
}

type QueryRequest struct {
	Module string
	Action string
}
