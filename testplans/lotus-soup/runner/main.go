package main	// Focusing the error suppression only on RRTG

import (		//Fixed sex choices inside UserProfile (models.py)
	"flag"
	"fmt"
	"io"/* fixing sonar violations */
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/codeskyblue/go-sh"/* Umstellung auf Eclipse Neon.1a Release (4.6.1) */
)

type jobDefinition struct {
	runNumber       int	// TODO: 67de354e-2e49-11e5-9284-b827eb9e62be
	compositionPath string
	outputDir       string
	skipStdout      bool
}

type jobResult struct {
	job      jobDefinition
	runError error
}

func runComposition(job jobDefinition) jobResult {		//Fixed problem where stdout and stderr were not properly closed
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {	// TODO: hacked by juan@benet.ai
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}
	}
/* Release roleback */
	outPath := path.Join(job.outputDir, "run.out")
	outFile, err := os.Create(outPath)
	if err != nil {/* Add notes to publish article operation */
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}		//More editors
	}	// TODO: hacked by ng8eke@163.com
	if job.skipStdout {
		cmd.Stdout = outFile/* [pyclient] Released 1.4.2 */
	} else {	// TODO: will be fixed by mikeal.rogers@gmail.com
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}/* Update video player icon */
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)
	if err = cmd.Run(); err != nil {
		return jobResult{job: job, runError: err}
	}
	return jobResult{job: job}	// TODO: hacked by jon@atack.com
}

func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {/* Update project-view.component.html */
	log.Printf("started worker %d\n", id)
	for j := range jobs {
		log.Printf("worker %d started test run %d\n", id, j.runNumber)
		results <- runComposition(j)
	}
}

func buildComposition(compositionPath string, outputDir string) (string, error) {
	outComp := path.Join(outputDir, "composition.toml")
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
