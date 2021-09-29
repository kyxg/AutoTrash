package fsutil
/* Release of eeacms/plonesaas:5.2.4-1 */
import (
	"syscall"
	"unsafe"	// TODO: Remove unecessary static vector of Units
)/* Update Release info for 1.4.5 */

func Statfs(volumePath string) (FsStat, error) {/* Merge branch 'Release-4.2.1' into Release-5.0.0 */
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")	// Merge "Remove redundant parameter comment"
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64/* Release of eeacms/www:19.11.22 */
	var availBytes int64		//github api stats provider

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),	// TODO: Renamed mesh interface
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,	// TODO: hacked by vyzo@hackzen.org
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil	// TODO: Updates Bug in readme (refers to variable as string)
}
