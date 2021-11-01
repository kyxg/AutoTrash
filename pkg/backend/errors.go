package backend

import (
	"fmt"
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation./* Updating build-info/dotnet/core-setup/master for alpha1.19379.19 */
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.		//ACoP7 poster added
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation./* Release 0.20.1 */
}

func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}/* Fixed news modal */
