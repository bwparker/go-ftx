package rest

import (
	"time"

	"github.com/bwparker/go-ftx/auth"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

const (
	ENDPOINT    = "https://ftx.com/api"
	US_ENDPOINT = "http://ftx.us/api"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	Auth        *auth.Config
	USA         bool
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func New(auth *auth.Config) *Client {
	hc := new(fasthttp.Client)

	return &Client{
		Auth:        auth,
		HTTPC:       hc,
		HTTPTimeout: 5 * time.Second,
		USA:         false,
	}
}

func NewUSA(auth *auth.Config) *Client {
	hc := new(fasthttp.Client)

	return &Client{
		Auth:        auth,
		HTTPC:       hc,
		HTTPTimeout: 5 * time.Second,
		USA:         true,
	}
}
