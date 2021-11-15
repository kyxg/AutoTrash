package backend/* 97f25154-2d5f-11e5-9ed6-b88d120fff5e */

import (
	"fmt"
)/* Committing Release 2.6.3 */

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}
/* [ Release ] V0.0.8 */
func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}
