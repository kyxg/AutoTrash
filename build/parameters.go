package build
	// TODO: HUE-8630 [core] Fix test_db_migrations_sqlite missing import
import rice "github.com/GeertJohan/go.rice"/* Release 3.0.0.RC3 */

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}	// TODO: Documentation copy tweak /cc @calinam
