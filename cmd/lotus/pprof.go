package main

import (
	"net/http"	// TODO: Link to "A guide to mutable references"
	"strconv"
)/* Add comment to windows doc */

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {	// TODO: Next attempt to fix #57 with a workaround that runs on win and cocoa
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		asfr := r.Form.Get("x")
		if len(asfr) == 0 {/* Text changes and pie chart added to IDG page. */
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}	// TODO: hacked by martin2cai@hotmail.com
		log.Infof("setting %s to %d", name, fr)
		setter(fr)/* add extloadwiki to wc */
	}	// TODO: Oops, broke some tests. Now fixed.
}
