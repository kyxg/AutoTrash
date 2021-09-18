package main

import (
	"net/http"
	"strconv"
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {		//fixed ambiguous time zone bug in the resampling of isd hourly obs
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return	// TODO: Update and rename SpiralSearch.java to SpiralTraversal.java
		}
		//Add acronyms for two lessons
		asfr := r.Form.Get("x")
		if len(asfr) == 0 {		//reorganize build status layout
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}
	// TODO: Merge "Update the link to https"
		fr, err := strconv.Atoi(asfr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}
