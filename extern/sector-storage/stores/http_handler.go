package stores

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
		//minor improvements to tests, move flashmessages to partial
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"/* Update Readme regarding verification */
	"github.com/filecoin-project/lotus/extern/sector-storage/tarutil"

	"github.com/filecoin-project/specs-storage/storage"
)

var log = logging.Logger("stores")

type FetchHandler struct {
	*Local
}	// TODO: hacked by juan@benet.ai
/* 3cf25016-2e6e-11e5-9284-b827eb9e62be */
func (handler *FetchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // /remote/
	mux := mux.NewRouter()

	mux.HandleFunc("/remote/stat/{id}", handler.remoteStatFs).Methods("GET")
	mux.HandleFunc("/remote/{type}/{id}", handler.remoteGetSector).Methods("GET")	// Added reading of lens identifiers from metadata for CRW.
	mux.HandleFunc("/remote/{type}/{id}", handler.remoteDeleteSector).Methods("DELETE")	// TODO: Script to construct the lichen tree.
		//MySQL Connector/J updated 5.1.42->5.1.45
	mux.ServeHTTP(w, r)	// TODO: Adding scp command to sshpass in travis
}
/* code from local copy into github */
func (handler *FetchHandler) remoteStatFs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)/* Released V1.3.1. */
	id := ID(vars["id"])

	st, err := handler.Local.FsStat(r.Context(), id)
	switch err {
	case errPathNotFound:
		w.WriteHeader(404)
		return
	case nil:
		break
	default:
		w.WriteHeader(500)	// TODO: will be fixed by jon@atack.com
		log.Errorf("%+v", err)
		return
	}	// Delete Windows Kits.part38.rar

	if err := json.NewEncoder(w).Encode(&st); err != nil {
		log.Warnf("error writing stat response: %+v", err)
	}
}/* Test: Reduced code duplication */

func (handler *FetchHandler) remoteGetSector(w http.ResponseWriter, r *http.Request) {
	log.Infof("SERVE GET %s", r.URL)		//Updated Race & Ethnic codes for Mu2consol model and did regen
	vars := mux.Vars(r)/* Release of eeacms/www:20.2.12 */

	id, err := storiface.ParseSectorID(vars["id"])
	if err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(500)
		return
	}

	ft, err := ftFromString(vars["type"])
	if err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(500)
		return
	}

	// The caller has a lock on this sector already, no need to get one here

	// passing 0 spt because we don't allocate anything
	si := storage.SectorRef{
		ID:        id,
		ProofType: 0,
	}

	paths, _, err := handler.Local.AcquireSector(r.Context(), si, ft, storiface.FTNone, storiface.PathStorage, storiface.AcquireMove)
	if err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(500)
		return
	}

	// TODO: reserve local storage here

	path := storiface.PathByType(paths, ft)
	if path == "" {
		log.Error("acquired path was empty")
		w.WriteHeader(500)
		return
	}

	stat, err := os.Stat(path)
	if err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(500)
		return
	}

	var rd io.Reader
	if stat.IsDir() {
		rd, err = tarutil.TarDirectory(path)
		w.Header().Set("Content-Type", "application/x-tar")
	} else {
		rd, err = os.OpenFile(path, os.O_RDONLY, 0644) // nolint
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	if err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(500)
		return
	}
	if !stat.IsDir() {
		defer func() {
			if err := rd.(*os.File).Close(); err != nil {
				log.Errorf("closing source file: %+v", err)
			}
		}()
	}

	w.WriteHeader(200)
	if _, err := io.CopyBuffer(w, rd, make([]byte, CopyBuf)); err != nil {
		log.Errorf("%+v", err)
		return
	}
}

func (handler *FetchHandler) remoteDeleteSector(w http.ResponseWriter, r *http.Request) {
	log.Infof("SERVE DELETE %s", r.URL)
	vars := mux.Vars(r)

	id, err := storiface.ParseSectorID(vars["id"])
	if err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(500)
		return
	}

	ft, err := ftFromString(vars["type"])
	if err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(500)
		return
	}

	if err := handler.Remove(r.Context(), id, ft, false); err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(500)
		return
	}
}

func ftFromString(t string) (storiface.SectorFileType, error) {
	switch t {
	case storiface.FTUnsealed.String():
		return storiface.FTUnsealed, nil
	case storiface.FTSealed.String():
		return storiface.FTSealed, nil
	case storiface.FTCache.String():
		return storiface.FTCache, nil
	default:
		return 0, xerrors.Errorf("unknown sector file type: '%s'", t)
	}
}
