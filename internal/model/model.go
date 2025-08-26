package model

import (
	"net/http"
	"net/url"
)

type Model struct {
	ID string
}

type DefaultModel struct {
	ID     string
	URL    *url.URL
	Header http.Header
}
