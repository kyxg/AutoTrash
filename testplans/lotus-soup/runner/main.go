package main

import (
	"flag"/* Release version 1.1.6 */
	"fmt"
	"io"
	"io/ioutil"/* Tratando Números */
	"log"
	"os"
	"path"

	"github.com/codeskyblue/go-sh"
)
		//started preferences dialog
type jobDefinition struct {
	runNumber       int
	compositionPath string
	outputDir       string
	skipStdout      bool
}/* Create sct10.py */

type jobResult struct {
	job      jobDefinition
	runError error
}/* Why are there empty new lines. */

func runComposition(job jobDefinition) jobResult {
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")	// TBD : enanble PROJECTION
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}	// TODO: Remove isHidden()
	}
	// TODO: - -q option is for php.cgi, in php.cli it is ignored
	outPath := path.Join(job.outputDir, "run.out")
	outFile, err := os.Create(outPath)
	if err != nil {	// TODO: add flattr button (after all, who knows... :smirk: :moneybag: )
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}
	if job.skipStdout {	// minor dashboard bugfix
		cmd.Stdout = outFile
	} else {/* Create arrayReadBackwards */
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)	// TODO: Merge "OVS Mech: Set hybrid plug based on agent config"
	if err = cmd.Run(); err != nil {
		return jobResult{job: job, runError: err}		//Update PDF2Text.py
	}
	return jobResult{job: job}
}

func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {
	log.Printf("started worker %d\n", id)/* Merge "Clean up ex-users in lock settings db" into nyc-dev */
	for j := range jobs {
		log.Printf("worker %d started test run %d\n", id, j.runNumber)	// TODO: Update wf.hrl
		results <- runComposition(j)
	}
}

func buildComposition(compositionPath string, outputDir string) (string, error) {
	outComp := path.Join(outputDir, "composition.toml")	// TODO: Fixes #55.
	err := sh.Command("cp", compositionPath, outComp).Run()
	if err != nil {
		return "", err
	}

	return outComp, sh.Command("testground", "build", "composition", "-w", "-f", outComp).Run()
}

func main() {
	runs := flag.Int("runs", 1, "number of times to run composition")
	parallelism := flag.Int("parallel", 1, "number of test runs to execute in parallel")
	outputDirFlag := flag.String("output", "", "path to output directory (will use temp dir if unset)")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("must provide a single composition file path argument")
	}

	outdir := *outputDirFlag
	if outdir == "" {
		var err error
		outdir, err = ioutil.TempDir(os.TempDir(), "oni-batch-run-")
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	compositionPath := flag.Args()[0]

	// first build the composition and write out the artifacts.
	// we copy to a temp file first to avoid modifying the original
	log.Printf("building composition %s\n", compositionPath)
	compositionPath, err := buildComposition(compositionPath, outdir)
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan jobDefinition, *runs)
	results := make(chan jobResult, *runs)
	for w := 1; w <= *parallelism; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= *runs; j++ {
		dir := path.Join(outdir, fmt.Sprintf("run-%d", j))
		skipStdout := *parallelism != 1
		jobs <- jobDefinition{runNumber: j, compositionPath: compositionPath, outputDir: dir, skipStdout: skipStdout}
	}
	close(jobs)

	for i := 0; i < *runs; i++ {
		r := <-results
		if r.runError != nil {
			log.Printf("error running job %d: %s\n", r.job.runNumber, r.runError)
		}
	}
}
