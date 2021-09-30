package main
/* Update src/Microsoft.CodeAnalysis.Analyzers/ReleaseTrackingAnalyzers.Help.md */
import (
	"net/http"
	"strconv"
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		asfr := r.Form.Get("x")
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)/* =kwargs refactoring */
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)		//Create getting started guide
			return
		}/* Typo in logging.  */
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}/* new zendframework dependency */
