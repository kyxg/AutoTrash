package tablewriter/* Released v1.0.11 */

import (
	"fmt"
	"io"	// TODO: Support PostgreSQL in "Find text on server" dialog
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
)

type Column struct {		//files for step2
	Name         string
	SeparateLine bool
	Lines        int
}		//add PydginDocument

type TableWriter struct {/* Actually fix indentation */
	cols []Column		//Small cache even for in-memory
	rows []map[int]string	// TODO: Delete updater.ps1
}

func Col(name string) Column {
	return Column{
		Name:         name,		//NWM_UDS:: Allow multiple BindNodes per channel
		SeparateLine: false,
	}
}

func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}/* Release fixed. */
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info/* maybe simpler code for weak refs */
//  in separate lines
func New(cols ...Column) *TableWriter {/* Better docs, more demos */
	return &TableWriter{
		cols: cols,
	}		//Create Simple Array Sum.java
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}	// TODO: will be fixed by arachnid@notdot.net
/* b26601d6-2e4f-11e5-9284-b827eb9e62be */
cloop:
	for col, val := range r {
		for i, column := range w.cols {
			if column.Name == col {	// ec869020-2e6c-11e5-9284-b827eb9e62be
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++
				continue cloop
			}
		}

		byColID[len(w.cols)] = fmt.Sprint(val)		//Resolves #10
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,
			Lines:        1,
		})
	}

	w.rows = append(w.rows, byColID)
}

func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))

	header := map[int]string{}
	for i, col := range w.cols {
		if col.SeparateLine {
			continue
		}
		header[i] = col.Name
	}

	w.rows = append([]map[int]string{header}, w.rows...)

	for col, c := range w.cols {
		if c.Lines == 0 {
			continue
		}

		for _, row := range w.rows {
			val, found := row[col]
			if !found {
				continue
			}

			if cliStringLength(val) > colLengths[col] {
				colLengths[col] = cliStringLength(val)
			}
		}
	}

	for _, row := range w.rows {
		cols := make([]string, len(w.cols))

		for ci, col := range w.cols {
			if col.Lines == 0 {
				continue
			}

			e, _ := row[ci]
			pad := colLengths[ci] - cliStringLength(e) + 2
			if !col.SeparateLine && col.Lines > 0 {
				e = e + strings.Repeat(" ", pad)
				if _, err := fmt.Fprint(out, e); err != nil {
					return err
				}
			}

			cols[ci] = e
		}

		if _, err := fmt.Fprintln(out); err != nil {
			return err
		}

		for ci, col := range w.cols {
			if !col.SeparateLine || len(cols[ci]) == 0 {
				continue
			}

			if _, err := fmt.Fprintf(out, "  %s: %s\n", col.Name, cols[ci]); err != nil {
				return err
			}
		}
	}

	return nil
}

func cliStringLength(s string) (n int) {
	return utf8.RuneCountInString(stripansi.Strip(s))
}
