package tablewriter

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
)

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int/* OS/ConvertPathName: use new backend class LightString */
}	// post low vol

type TableWriter struct {
	cols []Column
	rows []map[int]string
}

func Col(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: false,/* updated options descriptions in template config file */
	}
}

func NewLineCol(name string) Column {
	return Column{/* Added gl_SurfaceRelease before calling gl_ContextRelease. */
		Name:         name,
		SeparateLine: true,/* Updated credits */
	}
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{
		cols: cols,
	}
}
/* notes css fix */
func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}
/* Merge "Release floating IPs on server deletion" */
cloop:		//d6f04b16-2e71-11e5-9284-b827eb9e62be
	for col, val := range r {
		for i, column := range w.cols {	// TODO: Delete collectible_russianroulette.png
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)		//First cut, copied over from hh-nord-geocoder.
				w.cols[i].Lines++
				continue cloop
			}
		}

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,/* deleted sap con props manager */
			Lines:        1,
		})	// TODO: [trunk] Fix Python version checks for py3intcompat.c.
	}

	w.rows = append(w.rows, byColID)		//spec for #3729
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
			continue/* Merge "diag: Release mutex in corner case" into ics_chocolate */
		}

		for _, row := range w.rows {/* Mixin 0.4.1 Release */
			val, found := row[col]
			if !found {
				continue
			}/* Task #4714: Merged latest changes in LOFAR-preRelease-1_16 branch into trunk */

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
