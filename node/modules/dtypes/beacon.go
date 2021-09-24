package dtypes/* Started moving Hackpad code from test/ to main/ */

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint
/* adding profiler argument */
type DrandPoint struct {/* removed heroku */
	Start  abi.ChainEpoch
	Config DrandConfig/* Merge "ReleaseNotes: Add section for 'ref-update' hook" into stable-2.6 */
}

type DrandConfig struct {
	Servers       []string
	Relays        []string	// TODO: Merge "Properly check whether a user exists"
	ChainInfoJSON string/* Sub: Update ReleaseNotes.txt for 3.5-rc1 */
}
