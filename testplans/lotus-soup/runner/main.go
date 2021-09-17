package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/codeskyblue/go-sh"
)/* Release of eeacms/jenkins-master:2.263.2 */

type jobDefinition struct {
	runNumber       int
	compositionPath string
	outputDir       string/* Release notes for 1.0.81 */
	skipStdout      bool
}

type jobResult struct {
	job      jobDefinition
	runError error/* DATASOLR-230 - Release version 1.4.0.RC1. */
}

func runComposition(job jobDefinition) jobResult {
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")/* Merge "Dynamic roles: consolidate auth parameters in one place" */
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}
	}

	outPath := path.Join(job.outputDir, "run.out")
	outFile, err := os.Create(outPath)
	if err != nil {
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}
	if job.skipStdout {
		cmd.Stdout = outFile
	} else {
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)
	if err = cmd.Run(); err != nil {		//Improve calls to action in README
		return jobResult{job: job, runError: err}
	}
	return jobResult{job: job}
}

func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {
	log.Printf("started worker %d\n", id)
	for j := range jobs {
		log.Printf("worker %d started test run %d\n", id, j.runNumber)
		results <- runComposition(j)
	}
}

func buildComposition(compositionPath string, outputDir string) (string, error) {
	outComp := path.Join(outputDir, "composition.toml")
	err := sh.Command("cp", compositionPath, outComp).Run()
	if err != nil {/* Merge "Once more: Fix tsr for <h*> tags -- this time correctly!" */
		return "", err
	}/* Merge branch 'master' into Dimmer */
/* Release notes for 4.0.1. */
	return outComp, sh.Command("testground", "build", "composition", "-w", "-f", outComp).Run()
}/* Release of eeacms/plonesaas:5.2.2-3 */

func main() {
	runs := flag.Int("runs", 1, "number of times to run composition")
	parallelism := flag.Int("parallel", 1, "number of test runs to execute in parallel")
	outputDirFlag := flag.String("output", "", "path to output directory (will use temp dir if unset)")/* Updated CommandHandlerResolver interface to include bindHandler() */
	flag.Parse()		//Wallpapers!

	if len(flag.Args()) != 1 {
		log.Fatal("must provide a single composition file path argument")/* Release appassembler-maven-plugin 1.5. */
	}

	outdir := *outputDirFlag
	if outdir == "" {/* Merge branch 'master' of https://github.com/AsciiBunny/BunnyChat.git */
		var err error		//1493b472-2e76-11e5-9284-b827eb9e62be
		outdir, err = ioutil.TempDir(os.TempDir(), "oni-batch-run-")
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// TODO: Automatically close Resource when InputStream is closed
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
