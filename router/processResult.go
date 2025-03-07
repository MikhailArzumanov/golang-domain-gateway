package router

import (
	"net/http"
	"proxy/proxy"
)

func processResult(
	address string,
	respWriter http.ResponseWriter,
	sourceRequest *http.Request,
) {
	if address == "" {
		respWriter.Write([]byte("Domain is not valid."))
		respWriter.WriteHeader(500)
	} else {
		proxy.ProxyHandler(address, respWriter, sourceRequest)
	}
}
