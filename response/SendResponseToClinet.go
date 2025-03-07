package response

import "net/http"

func SendResponseToClient(
	resp *http.Response,
	method string,
	writer http.ResponseWriter,
) {
	respBody, respCode :=
		getBaseData(resp)
	writeRespHeaders(resp, writer)
	writeRespCode(respCode, method, writer)
	writeRespBody(writer, respBody)
}
