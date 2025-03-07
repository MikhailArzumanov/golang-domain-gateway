package target

import (
	"net/http"
	"proxy/error_handling"
)

func createRequest(
	sourceRequest *http.Request,
	destHostAddress string,
) *http.Request {

	method, url, body :=
		sourceRequest.Method,
		getUrl(sourceRequest, destHostAddress),
		getBodyBfr(sourceRequest)

	theRequest, err :=
		http.NewRequest(method, url, body)
	error_handling.Handle(err)

	return theRequest
}
