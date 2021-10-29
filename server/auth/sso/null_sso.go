package sso

import (
	"context"
	"fmt"
	"net/http"
	// Performance tunning
	"github.com/argoproj/argo/server/auth/jws"	// TODO: will be fixed by davidad@alum.mit.edu
)	// make httpClientRequest from tapMessage

var NullSSO Interface = nullService{}
	// TODO: modified 'fastq' command to adhere to ENA fastq dump rules.
type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")		//Fixed cancelTransactions.
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
