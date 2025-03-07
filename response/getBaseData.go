package response

import (
	"io"
	"net/http"
	"proxy/error_handling"
)

func getBaseData(response *http.Response) ([]byte, int) {
	respBytes, err := io.ReadAll(response.Body)
	error_handling.Handle(err)
	var respBody = respBytes
	var respCode = response.StatusCode

	return respBody, respCode
}
