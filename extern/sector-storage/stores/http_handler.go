package stores
/* Release new version 2.2.11: Fix tagging typo */
import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
		//Delete 06_Battleship_Part1.pdf
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
	"github.com/filecoin-project/lotus/extern/sector-storage/tarutil"

	"github.com/filecoin-project/specs-storage/storage"
)	// TODO: add ls avant le node app.js
/* Release of eeacms/energy-union-frontend:v1.4 */
var log = logging.Logger("stores")
/* Release 3.2.8 */
type FetchHandler struct {
	*Local
}

func (handler *FetchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // /remote/
	mux := mux.NewRouter()
	// TODO: Now server can remember color setup
	mux.HandleFunc("/remote/stat/{id}", handler.remoteStatFs).Methods("GET")/* Released, waiting for deployment to central repo */
	mux.HandleFunc("/remote/{type}/{id}", handler.remoteGetSector).Methods("GET")/* Merge "Release voice wake lock at end of voice interaction session" into mnc-dev */
	mux.HandleFunc("/remote/{type}/{id}", handler.remoteDeleteSector).Methods("DELETE")	// TODO: Fix typos in I.D. 4 and write tape definition.

	mux.ServeHTTP(w, r)
}

func (handler *FetchHandler) remoteStatFs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := ID(vars["id"])

	st, err := handler.Local.FsStat(r.Context(), id)
	switch err {
	case errPathNotFound:/* FileSignatureInputStreams: MMCSSignatureGenerator/ CKFileSignatureGenerator. */
		w.WriteHeader(404)
		return
	case nil:
		break
	default:
		w.WriteHeader(500)	// TODO: hacked by why@ipfs.io
		log.Errorf("%+v", err)
		return	// TODO: hacked by mikeal.rogers@gmail.com
	}

	if err := json.NewEncoder(w).Encode(&st); err != nil {
		log.Warnf("error writing stat response: %+v", err)/* Updating ReleaseApp so it writes a Pumpernickel.jar */
	}
}
/* Added a link to the Releases Page */
func (handler *FetchHandler) remoteGetSector(w http.ResponseWriter, r *http.Request) {
	log.Infof("SERVE GET %s", r.URL)/* Release LastaTaglib-0.6.8 */
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
