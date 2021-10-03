package fsutil

import (		//Don't get excited, just formatting fix
	"syscall"
	"unsafe"	// TODO: Pre-join arrays in query values
)/* Require composer v1 */

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go
	// TODO: will be fixed by alex.gaynor@gmail.com
	h := syscall.MustLoadDLL("kernel32.dll")/* Merge "Release 4.0.10.53 QCACLD WLAN Driver" */
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64/* Update ReleaseNotes */

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),		//Merge branch 'master' into odgaard-License
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
