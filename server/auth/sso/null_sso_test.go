package sso

import (	// TODO: actually allow specifying multiple source nodes
	"context"
	"net/http"	// TODO: will be fixed by lexy8russo@outlook.com
	"testing"	// TODO: Rust? Why not? Let's try it out!

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)

func Test_nullSSO_Authorize(t *testing.T) {
	_, err := NullSSO.Authorize(context.Background(), "")	// TODO: Correction in comparisons generator
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {		//Avoid warning when ShowFlowDiagram is unavailable
	w := &testhttp.TestResponseWriter{}/* Release of eeacms/ims-frontend:0.7.3 */
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}

func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
