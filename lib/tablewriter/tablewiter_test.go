package tablewriter

import (
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{/* Be paranoid and unlink build/bin before creating a new symlink */
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",/* Added testcase for insertion before time-out */
	})		//NGS fixed frame rate
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{
		"C1":             "1",
		"C333":           "2",/* Release 2.1.17 */
		"SurpriseColumn": "42",	// TODO: hacked by nicksavers@gmail.com
	})
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}/* Update howto use this library */
