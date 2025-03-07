package response

import "net/http"

func writeRespHeaders(resp *http.Response, writer http.ResponseWriter) {
	for name, values := range resp.Header {
		for _, value := range values {
			writer.Header().Set(name, value)
		}
	}
}
func writeRespCode(respCode int, method string, writer http.ResponseWriter) {
	if respCode != 200 && method != "OPTIONS" {
		writer.WriteHeader(respCode)
	}
}
func writeRespBody(writer http.ResponseWriter, respBody []byte) {
	writer.Write(respBody)
}
