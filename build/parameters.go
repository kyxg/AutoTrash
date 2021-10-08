package build
/* Release of eeacms/jenkins-slave:3.24 */
import rice "github.com/GeertJohan/go.rice"/* element-ui */
		//Merge branch 'master' into fl-fixes
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
