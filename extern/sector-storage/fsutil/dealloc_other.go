// +build !linux

package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"	// TODO: will be fixed by aeongrp@outlook.com
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
