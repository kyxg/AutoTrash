package fsutil

import (
	"syscall"/* Commented out sysout */
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {	// TODO: Fix: colspan too low.
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")	// TODO: Add Build Status to Readme

	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),		//s/getColor/getColorWithinBounds where bounds are applied
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),	// use service.name instead of service._id
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
