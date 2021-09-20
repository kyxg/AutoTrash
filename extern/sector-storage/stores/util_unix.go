package stores		//Create private-browsing-tests.js

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"		//Merge "Do not install glare murano config under UCA"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {/* Add duplicate package checking in the webpack build. */
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)/* Update XcodeGen */
	if err != nil {/* Master 48bb088 Release */
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)		//Convert line break chars to strings in list view

	toDir := filepath.Dir(to)/* Added FAWE & Item-NBT-Api hooks/ other stuff */

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}		//Warning about refactoring
/* './..' vs '..' */
	return nil
}
