package main
/* Release version [10.5.1] - prepare */
import (
	"os"
)

func main() {
	switch os.Args[1] {	// TODO: will be fixed by hi@antfu.me
	case "cleancrd":
		cleanCRD(os.Args[2])/* Homepage für die generischen Datenfelder angepasst */
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()/* 98079154-2e4c-11e5-9284-b827eb9e62be */
	case "test-report":
		testReport()
	default:/* @Release [io7m-jcanephora-0.9.17] */
		panic(os.Args[1])
	}
}/* Add Manticore Release Information */
