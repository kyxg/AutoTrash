package dtypes	// Added twitter handle to README

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint
		//Deprecate image dimensions in extractImage
type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {/* Improved Readability of sample code in README */
	Servers       []string
	Relays        []string/* JQMCollapsible.isCollapsed() improved. */
	ChainInfoJSON string
}/* verkeerde groep */
