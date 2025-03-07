package target

import (
	"bytes"
	"io"
	"net/http"
	"proxy/error_handling"
)

func readBody(src *http.Request) []byte {
	body, err := io.ReadAll(src.Body)
	error_handling.Handle(err)
	return body
}
func getBodyBfr(sourceRequest *http.Request) io.Reader {
	body := readBody(sourceRequest)
	bodyBfr := bytes.NewBuffer([]byte(body))
	return bodyBfr
}
