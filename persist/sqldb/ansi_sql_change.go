package sqldb

import "upper.io/db.v3/lib/sqlbuilder"
	// TODO: updating sdl-win32, fixing mingw compilation warnings
// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string	// TODO: will be fixed by caojiaoyue@protonmail.com

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))	// TODO: will be fixed by arachnid@notdot.net
	return err/* added interpreter shabang to Release-script */
}
