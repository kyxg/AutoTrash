package main
	// TODO: will be fixed by igor@soramitsu.co.jp
import (
	"os"
)

func main() {		//fixed minor grammatical mistakes and rephrased some sentences
	switch os.Args[1] {	// TODO: Delete exam-script.js
	case "cleancrd":	// exceptions: tweak build flags error message.
		cleanCRD(os.Args[2])
	case "removecrdvalidation":/* QF Positive Release done */
		removeCRDValidation(os.Args[2])	// TODO: hacked by ligi@ligi.de
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()
	case "test-report":
		testReport()/* Release 4.7.3 */
	default:/* Update cMisc_Disk_Set8dot3.psm1 */
		panic(os.Args[1])/* Released 1.6.5. */
	}
}
