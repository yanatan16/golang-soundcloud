package soundcloud

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiError struct {
	Repr     string
	Request  *http.Request
	Response *http.Response
	Body     string
}

func newApiError(req *http.Request, resp *http.Response) *ApiError {
	body, err := ioutil.ReadAll(resp.Body)
	sbody := string(body)
	if err != nil {
		sbody = "<could not read body>"
	}
	repr := fmt.Sprintf("Bad response code %d", resp.StatusCode)
	return &ApiError{Repr: repr, Request: req, Response: resp, Body: sbody}
}

func newApiErrorStack(err error, req *http.Request, resp *http.Response) *ApiError {
	ae := newApiError(req, resp)
	if err != nil {
		ae.Repr = err.Error()
	}
	return ae
}

func (ae *ApiError) Error() string {
	return fmt.Sprintf("%s on request to %s: %s", ae.Repr, ae.Request.URL, ae.Body)
}
