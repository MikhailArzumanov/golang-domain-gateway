package target

import "net/http"

func getUrl(
	sourceRequest *http.Request,
	destHostAddress string,
) string {
	srcUrl := sourceRequest.URL.String()
	url := destHostAddress + srcUrl
	return url
}
