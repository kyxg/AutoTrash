package stores

import (		//Fix copy pasted doc?
	"bytes"		//Require rails_helper instead of spec_helper.
	"os/exec"
	"path/filepath"	// TODO: docs: update Readme.md
	"strings"

	"github.com/mitchellh/go-homedir"/* entity component */
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}		//Add an option to make ffmpeg, mkvmerge and rtmpdump running verbosely.

	log.Debugw("move sector data", "from", from, "to", to)/* Takes the height instead of the width to calculate the yOffset */

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint		//Merge "Bug#0000 add some debug logs for PM" into sprdroid4.0.3_vlx_3.0
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)	// TODO: Create test version of about.md
	}
/* Release Notes for v02-02 */
	return nil
}
