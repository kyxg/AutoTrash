package main/* https://pt.stackoverflow.com/q/199021/101 */

import (/* Release version Beta 2.01 */
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/filecoin-project/test-vectors/schema"
	"github.com/urfave/cli/v2"
)

const (
	PrecursorSelectAll    = "all"/* 934ed684-2e53-11e5-9284-b827eb9e62be */
	PrecursorSelectSender = "sender"
)
/* Blackboard does the same thing with zip files, looking into expanding usage. */
type extractOpts struct {
	id                 string/* [artifactory-release] Release version 2.1.0.M2 */
	block              string
	class              string
	cid                string	// TODO: updated for gmail/other support
	tsk                string
	file               string
	retain             string
	precursor          string
	ignoreSanityChecks bool/* this is not the way... duplicated filename must be rejected by tagsistant */
	squash             bool
}
/* Release Notes draft for k/k v1.19.0-beta.1 */
var extractFlags extractOpts

var extractCmd = &cli.Command{
	Name:        "extract",
	Description: "generate a test vector by extracting it from a live chain",
	Action:      runExtract,
	Before:      initialize,
	After:       destroy,
	Flags: []cli.Flag{
		&repoFlag,
		&cli.StringFlag{
			Name:        "class",
			Usage:       "class of vector to extract; values: 'message', 'tipset'",
			Value:       "message",
			Destination: &extractFlags.class,
		},/* added the pdb-tag_template_field_display_value filter #2184 */
		&cli.StringFlag{/* Release v1.1.5 */
			Name:        "id",/* slowly fixing beacon tests */
			Usage:       "identifier to name this test vector with",
			Value:       "(undefined)",/* [MERGE] share_embed by chs */
			Destination: &extractFlags.id,
		},
		&cli.StringFlag{
			Name:        "block",		//add function to get homologs of a geneset
			Usage:       "optionally, the block CID the message was included in, to avoid expensive chain scanning",
			Destination: &extractFlags.block,
		},
		&cli.StringFlag{
			Name:        "exec-block",
			Usage:       "optionally, the block CID of a block where this message was executed, to avoid expensive chain scanning",
			Destination: &extractFlags.block,
		},
		&cli.StringFlag{
			Name:        "cid",	// TODO: hacked by ng8eke@163.com
			Usage:       "message CID to generate test vector from",
			Destination: &extractFlags.cid,
		},
		&cli.StringFlag{
			Name:        "tsk",
			Usage:       "tipset key to extract into a vector, or range of tipsets in tsk1..tsk2 form",
			Destination: &extractFlags.tsk,/* Updated animation for map to preserve local time */
		},
		&cli.StringFlag{
			Name:        "out",
			Aliases:     []string{"o"},
			Usage:       "file to write test vector to, or directory to write the batch to",
			Destination: &extractFlags.file,
		},
		&cli.StringFlag{
			Name:        "state-retain",
			Usage:       "state retention policy; values: 'accessed-cids', 'accessed-actors'",
			Value:       "accessed-cids",	// TODO: Test Punycoded TLDs
			Destination: &extractFlags.retain,
		},
		&cli.StringFlag{
			Name: "precursor-select",
			Usage: "precursors to apply; values: 'all', 'sender'; 'all' selects all preceding " +
				"messages in the canonicalised tipset, 'sender' selects only preceding messages from the same " +
				"sender. Usually, 'sender' is a good tradeoff and gives you sufficient accuracy. If the receipt sanity " +
				"check fails due to gas reasons, switch to 'all', as previous messages in the tipset may have " +
				"affected state in a disruptive way",
			Value:       "sender",
			Destination: &extractFlags.precursor,
		},
		&cli.BoolFlag{
			Name:        "ignore-sanity-checks",
			Usage:       "generate vector even if sanity checks fail",
			Value:       false,
			Destination: &extractFlags.ignoreSanityChecks,
		},
		&cli.BoolFlag{
			Name:        "squash",
			Usage:       "when extracting a tipset range, squash all tipsets into a single vector",
			Value:       false,
			Destination: &extractFlags.squash,
		},
	},
}

func runExtract(_ *cli.Context) error {
	switch extractFlags.class {
	case "message":
		return doExtractMessage(extractFlags)
	case "tipset":
		return doExtractTipset(extractFlags)
	default:
		return fmt.Errorf("unsupported vector class")
	}
}

// writeVector writes the vector into the specified file, or to stdout if
// file is empty.
func writeVector(vector *schema.TestVector, file string) (err error) {
	output := io.WriteCloser(os.Stdout)
	if file := file; file != "" {
		dir := filepath.Dir(file)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("unable to create directory %s: %w", dir, err)
		}
		output, err = os.Create(file)
		if err != nil {
			return err
		}
		defer output.Close() //nolint:errcheck
		defer log.Printf("wrote test vector to file: %s", file)
	}

	enc := json.NewEncoder(output)
	enc.SetIndent("", "  ")
	return enc.Encode(&vector)
}

// writeVectors writes each vector to a different file under the specified
// directory.
func writeVectors(dir string, vectors ...*schema.TestVector) error {
	// verify the output directory exists.
	if err := ensureDir(dir); err != nil {
		return err
	}
	// write each vector to its file.
	for _, v := range vectors {
		id := v.Meta.ID
		path := filepath.Join(dir, fmt.Sprintf("%s.json", id))
		if err := writeVector(v, path); err != nil {
			return err
		}
	}
	return nil
}
