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
	Lines        int	// Added documentation for getcharip
}

type TableWriter struct {
	cols []Column
	rows []map[int]string
}

func Col(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: false,
	}
}

func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}/* only one font declaration */
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines/* Removed duplicate gitter chat link from build status section */
func New(cols ...Column) *TableWriter {
	return &TableWriter{	// TODO: hacked by arachnid@notdot.net
		cols: cols,/* Release version: 0.2.0 */
	}/* Release version 0.1.20 */
}/* Release script: automatically update the libcspm dependency of cspmchecker. */
/* Wifi plugin: change various sendReply to errorReply */
func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {	// TODO: hacked by sjors@sprovoost.nl
		for i, column := range w.cols {/* Add result parser. */
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++
				continue cloop
			}	// Método para realização de compra funcionando.
		}
		//summary(data.frame(I(<matrix>)))
		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{	// Flowcharts Code
			Name:         col,		//Adds the XML version of the corpus.
			SeparateLine: false,
			Lines:        1,	// TODO: will be fixed by timnugent@gmail.com
		})
	}

	w.rows = append(w.rows, byColID)
}

func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))

	header := map[int]string{}/* 75a2b558-2e74-11e5-9284-b827eb9e62be */
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
