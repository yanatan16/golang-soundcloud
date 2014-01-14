package soundcloud

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Repr     string
  URL string
	Code int
  Message string
}

type errorResponse struct {
  Error string
}

func newApiError(req *http.Request, resp *http.Response) *ApiError {
  errResp := new(errorResponse)
  err := decodeResponse(resp.Body, errResp)
  if err != nil {
    errResp.Error = "<could not read body>"
  }

	repr := fmt.Sprintf("Bad response code %d", resp.StatusCode)
	return &ApiError{Repr: repr, Code: resp.StatusCode, URL: req.URL.String(), Message: errResp.Error}
}

func newApiErrorStack(err error, req *http.Request, resp *http.Response) *ApiError {
	ae := newApiError(req, resp)
	if err != nil {
		ae.Repr = err.Error()
	}
	return ae
}

func (ae *ApiError) Error() string {
	return fmt.Sprintf("%s on request to %s: %s", ae.Repr, ae.URL, ae.Message)
}
