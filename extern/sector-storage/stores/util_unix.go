package stores

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {/* Merge "Fix missing check for admin/adv_service" into feature/pecan */
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)		//Enable Maven debug and disable counter on error to make test stable
	}

	to, err = homedir.Expand(to)	// TODO: will be fixed by caojiaoyue@protonmail.com
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}		//Add the ability to wrap jars

	return nil
}/* f9625666-2e68-11e5-9284-b827eb9e62be */
