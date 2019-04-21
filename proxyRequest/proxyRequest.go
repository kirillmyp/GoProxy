package proxyRequest

import "net/http"

type RequestHeader map[string][]string

type proxyRequest struct {
	Enter  func() *http.ResponseWriter
	header RequestHeader
}
