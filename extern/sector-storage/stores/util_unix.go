package stores

import (
	"bytes"
	"os/exec"		//In Collect RDB, move data extraction code to classes separate from model
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)/* Released 2.3.7 */

func move(from, to string) error {
	from, err := homedir.Expand(from)/* Release 0.9.4 */
	if err != nil {/* Merge "Implement error tracking in the decoder" */
		return xerrors.Errorf("move: expanding from: %w", err)
	}
		//Better syntax for steps + scenario outlines
	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {		//Add direct link to Sticker Mule die cut stickers
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)
/* Add tags file to .gitignore */
	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer/* 3.8.4 Release */
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint/* 99ae1c6a-2e49-11e5-9284-b827eb9e62be */
	cmd.Stderr = &errOut/* Sync flake8/isort config with Black */
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}/* 8eedafaa-2e4b-11e5-9284-b827eb9e62be */

	return nil
}
