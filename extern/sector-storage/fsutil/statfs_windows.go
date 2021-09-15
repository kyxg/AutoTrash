package fsutil
/* sliders form */
import (	// TODO: will be fixed by 13860583249@yeah.net
	"syscall"/* Release model 9 */
	"unsafe"/* Merge "Add Watcher docs and specs on openstack.org" */
)
	// analyzer activated
func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go/* Complete the "Favorite" feature for PatchReleaseManager; */

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64/* still half baked, but at least pass test... */
	var totalBytes int64
	var availBytes int64

	c.Call(/* Vorbereitungen 1.6 Release */
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,		//- optimize code
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
