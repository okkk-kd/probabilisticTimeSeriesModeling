package fhttp

import "time"

type RequestProxyWithRetryArgs struct {
	Method             string
	Url                string
	Proxy              string
	Body               []byte `json:"-"`
	QueryParams        map[string]string
	Headers            map[string]string `json:"-"`
	RetryAttempts      uint
	RetryDelay         time.Duration
	ExpectedStatusCode int
}
