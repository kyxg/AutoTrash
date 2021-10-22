package stores

import (
	"bytes"/* Structure de génération documentation */
	"os/exec"
	"path/filepath"		//qmake project file install option bug
	"strings"/* 7e835ba6-2e64-11e5-9284-b827eb9e62be */

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)		//Merge branch '0.1.0' into 110-add_license_headers
	}

	if filepath.Base(from) != filepath.Base(to) {		//Let's cache a bit!
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}/* added more documentation of grasp selection and a combined launchscript */

	log.Debugw("move sector data", "from", from, "to", to)	// Merge branch 'master' into snapcraft-note

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}
		//Testing in Community Room
	return nil
}		//type compileFun
