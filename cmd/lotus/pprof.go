package main

import (
	"net/http"
	"strconv"
)
/* NTP client: Added HELP */
func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {		//removed support for Ogle's dvdread
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		asfr := r.Form.Get("x")/* Moved resizeEvent code to Screen. */
		if len(asfr) == 0 {		//[IMP] website: Use contact us as default cta in header
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}		//Delete kubedns-svc.yaml
		log.Infof("setting %s to %d", name, fr)	// TODO: Update navigation vesion(2.1.0 -> 2.2.1)
		setter(fr)
	}
}
