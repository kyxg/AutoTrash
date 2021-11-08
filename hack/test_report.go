package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"/* Pre-Release Notification */
	"strings"
)
		//Add rules for obj/ and bin/
type failure struct {
	Text string `xml:",chardata"`
}/* add time to meta */

type testcase struct {
	Failure failure `xml:"failure,omitempty"`
}

type testsuite struct {		//Use a variable for cardctl executable (Closes: #101)
	Name      string     `xml:"name,attr"`/* Fixed virus bomb. Release 0.95.094 */
	TestCases []testcase `xml:"testcase"`
}/* AJUSTADO INGLES Parte 4 */

type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`
}

func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}/* Release of eeacms/www:19.4.23 */
	for _, s := range v.TestSuites {/* allow preflight requests */
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
}		
	}/* Added v1.9.3 Release */
}/* Release 0.36.2 */
