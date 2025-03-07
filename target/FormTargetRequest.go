package target

import (
	"net/http"
)

func FormTargetRequest(
	sourceRequest *http.Request,
	destHostAddress string,
) *http.Request {
	destRequest := createRequest(sourceRequest, destHostAddress)
	setRequestHeaders(sourceRequest, destRequest)
	return destRequest
}
