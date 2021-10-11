package tablewriter

import (
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{	// TODO: hacked by alessio@tendermint.com
		"C1":   "234",
		"C333": "ou",		//Added text document generator.
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",		//put the patch in here too
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",	// Fix svn properties.
	})		//Updated Avatar ☺
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{
		"C1":             "1",
		"C333":           "2",
		"SurpriseColumn": "42",
	})
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}
