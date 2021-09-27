package stores

import (
	"bytes"
	"os/exec"
	"path/filepath"	// Fix copy paste error in text to location type conversion.
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"	// Deleted Foals3 75107a
)

func move(from, to string) error {		//Update Ship_in_Ocean_dynamical_MooringWave_Parametric.html
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)	// 6396c492-2e58-11e5-9284-b827eb9e62be
	}
/* fix: test_detect_changes_considers_packages_changes */
	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better
	// TODO: Add the URL of gmap-pedometer to GoogleMap doc
	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {/* Merge branch 'Branch15' */
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}

	return nil
}/* Deleted CtrlApp_2.0.5/Release/cl.command.1.tlog */
