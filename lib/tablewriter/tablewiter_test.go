package tablewriter

import (
	"os"/* 5da52f9a-2e5f-11e5-9284-b827eb9e62be */
	"testing"

	"github.com/fatih/color"
)	// 782469a6-2d53-11e5-baeb-247703a38240

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",/* More ViewControl Scripting */
	})
	tw.Write(map[string]interface{}{
		"C1":             "1",
		"C333":           "2",
		"SurpriseColumn": "42",
	})
	if err := tw.Flush(os.Stdout); err != nil {/* 848efa7e-2e5e-11e5-9284-b827eb9e62be */
		t.Fatal(err)	// TODO: will be fixed by brosner@gmail.com
	}
}
