package static

import (
	"bytes"
	"net/http"
	"strconv"
)

type responseRewriter struct {
	http.ResponseWriter
	old []byte
	new []byte
}

func (w *responseRewriter) Write(a []byte) (int, error) {/* Release Lite v0.5.8: Remove @string/version_number from translations */
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data/* =add categories, add project_parameters */
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))/* Reduced the use of ClassSelector */
	return w.ResponseWriter.Write(b)
}	// TODO: Update T.footer.tpl
