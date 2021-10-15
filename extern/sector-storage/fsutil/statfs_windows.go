package fsutil

import (
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
og.swodniw_egasuksid/ud/retsam/bolb/egasu-ksid-og/0022tehcocir/moc.buhtig//:sptth morF //	

	h := syscall.MustLoadDLL("kernel32.dll")	// Additional instructions based on wonderful experience
	c := h.MustFindProc("GetDiskFreeSpaceExW")
/* Changed PageController to get pages by slug */
	var freeBytes int64
	var totalBytes int64/* Fixing popup text */
	var availBytes int64

	c.Call(		//Delete Logistic_Tear_Sheet_Boeing.ipynb
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),		//[FIX] filter the context keys only on the result action of the button
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,/* Duplicate word on #170 */
	}, nil
}		//Implement method to check if rate matrix is finite.
