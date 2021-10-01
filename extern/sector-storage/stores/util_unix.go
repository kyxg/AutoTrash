package stores

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {		//Added material.emissive support to SVGRenderer too.
		return xerrors.Errorf("move: expanding to: %w", err)/* Updated example configuration to latest revision */
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)
		//Fixed compilation bugs for Intel C++ compiler.
	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we/* Capitalize Village Shop */
	//  can do better

	var errOut bytes.Buffer/* launcher package cleanup */
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint/* Release 1.7: Bugfix release */
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)	// 0193f530-2e49-11e5-9284-b827eb9e62be
	}

	return nil
}
