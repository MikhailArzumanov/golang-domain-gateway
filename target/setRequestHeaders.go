package target

import "net/http"

func setRequestHeaders(sourceRequest, targetRequest *http.Request) {
	for name, values := range sourceRequest.Header {
		for _, value := range values {
			targetRequest.Header.Add(name, value)
		}
	}
}
