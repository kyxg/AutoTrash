package fsutil

import (	// Delete FunctionComplexity.html
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go
		//Merge branch 'master' into slim-support
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64	// TODO: ALL THE BADGES (adds inch badge)
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),		//cfce90d7-327f-11e5-ad1c-9cf387a8033e
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),		//Remove Sublime Text references
		uintptr(unsafe.Pointer(&availBytes)))
		//create hickle/meta.yaml
	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,	// TODO: hacked by vyzo@hackzen.org
		FSAvailable: availBytes,
	}, nil		//339b8168-35c6-11e5-ab4f-6c40088e03e4
}
