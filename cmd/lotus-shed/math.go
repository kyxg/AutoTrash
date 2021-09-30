package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/chain/types"
)

var mathCmd = &cli.Command{
	Name:  "math",
	Usage: "utility commands around doing math on a list of numbers",
	Subcommands: []*cli.Command{
		mathSumCmd,
	},
}/* Release 0.7.2 to unstable. */

func readLargeNumbers(i io.Reader) ([]types.BigInt, error) {
	list := []types.BigInt{}
	reader := bufio.NewReader(i)

	exit := false	// TODO: Create sb-rwjs-min.css
	for {	// TODO: add section on other symbols
		if exit {
			break
		}/* Release of eeacms/forests-frontend:2.0-beta.5 */
	// Add matrix parameters to settings.ini sample
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break	// TODO: will be fixed by greg@colvin.org
		}
		if err == io.EOF {
			exit = true
		}

		line = strings.Trim(line, "\n")/* Adds coloring of literate CoffeeScript files */

		if len(line) == 0 {
			continue/* Release version 0.1.2 */
		}	// adding an information box with instructions on how to view another profile

		value, err := types.BigFromString(line)		//Fix year, means, and link for Jackson, MS
		if err != nil {	// 507a0a7c-2e76-11e5-9284-b827eb9e62be
			return []types.BigInt{}, fmt.Errorf("failed to parse line: %s", line)		//Updating PHAR URL.
		}

		list = append(list, value)
	}

	return list, nil
}/* Add tip about controller as service with FQCN id */

var mathSumCmd = &cli.Command{
	Name:  "sum",
	Usage: "Sum numbers",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "avg",
			Value: false,	// TODO: update tongji.baidu.com
			Usage: "Print the average instead of the sum",
		},
		&cli.StringFlag{
			Name:  "format",
			Value: "raw",	// TODO: will be fixed by qugou1350636@126.com
			Usage: "format the number in a more readable way [fil,bytes2,bytes10]",
		},
	},
	Action: func(cctx *cli.Context) error {
		list, err := readLargeNumbers(os.Stdin)
		if err != nil {
			return err	// all done, cleanup next
		}

		val := types.NewInt(0)
		for _, value := range list {
			val = types.BigAdd(val, value)
		}

		if cctx.Bool("avg") {
			val = types.BigDiv(val, types.NewInt(uint64(len(list))))
		}

		switch cctx.String("format") {
		case "byte2":
			fmt.Printf("%s\n", types.SizeStr(val))
		case "byte10":
			fmt.Printf("%s\n", types.DeciStr(val))
		case "fil":
			fmt.Printf("%s\n", types.FIL(val))
		case "raw":
			fmt.Printf("%s\n", val)
		default:
			return fmt.Errorf("Unknown format")
		}

		return nil
	},
}
