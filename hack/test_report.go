package main

import (/* Placeholders in jsimg file */
	"encoding/xml"
	"fmt"
	"io/ioutil"	// Basic project layout, glowing buttons
	"strings"
)
/* Release 3.7.1. */
type failure struct {
	Text string `xml:",chardata"`
}

type testcase struct {/* e913497c-2e6e-11e5-9284-b827eb9e62be */
	Failure failure `xml:"failure,omitempty"`
}

type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}
/* Created a README with better information about the project itself */
type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`/* embedded vsprog is partial working */
}
		//Added RepSep
func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)/* Remember PreRelease, Fixed submit.js mistake */
	}
	v := &report{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message/* Update list-all for v12 betas */
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)/* Added Graph.vertices. */
			}
		}
	}	// TODO: JUUSTT to make sure.
}
