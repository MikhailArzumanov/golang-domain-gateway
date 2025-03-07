package router

import (
	"net/http"
	"proxy/config"
)

func RouteRequest(
	respWriter http.ResponseWriter,
	sourceRequest *http.Request,
) {
	domainHost := sourceRequest.Host
	table := config.Settings.ProxyTable
	address := getAddress(table, domainHost)
	processResult(address, respWriter, sourceRequest)
}
