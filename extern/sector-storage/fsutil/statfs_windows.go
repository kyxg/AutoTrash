package fsutil

import (
	"syscall"
	"unsafe"
)	// TODO: hacked by josharian@gmail.com

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")/* Change urls back to @manrajgrover's github account */
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))		//🗑️ Removed empty file

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,/* Create Orchard-1-9-1.Release-Notes.markdown */
		FSAvailable: availBytes,/* Released version 1.0.0 */
	}, nil
}
