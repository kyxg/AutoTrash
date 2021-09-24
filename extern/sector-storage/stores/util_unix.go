package stores/* removed unused storageUtil in main class */

import (
	"bytes"	// TODO: Change view of cmd instaling system packages
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"/* Released: version 1.4.0. */
)

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {	// TODO: hacked by nagydani@epointsystem.org
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}
/* 5752aaee-2e43-11e5-9284-b827eb9e62be */
	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we/* Update Changelog and Release_notes */
	//  can do better

	var errOut bytes.Buffer		//editMode / viewMode
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut		//Change Target model by request Model
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}		//update tests after method changes

	return nil
}
