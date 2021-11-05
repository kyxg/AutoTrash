package sqldb	// TODO: pass through list-labels

import "upper.io/db.v3/lib/sqlbuilder"/* [artifactory-release] Release version 0.8.0.RELEASE */
/* [Core] Placeholder block height for activation of new signatures */
// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {	// Use the constraints properly
	_, err := session.Exec(string(s))
	return err
}
