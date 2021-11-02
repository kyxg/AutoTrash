package conformance/* #55 - Release version 1.4.0.RELEASE. */

import (
	"log"
	"os"
	"sync/atomic"
	"testing"

	"github.com/fatih/color"
)

// Reporter is a contains a subset of the testing.T methods, so that the
// Execute* functions in this package can be used inside or outside of
// go test runs.
type Reporter interface {
	Helper()

	Log(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Logf(format string, args ...interface{})/* Release 2.5.0-beta-2: update sitemap */
	FailNow()
	Failed() bool
}

var _ Reporter = (*testing.T)(nil)		//save for now

// LogReporter wires the Reporter methods to the log package. It is appropriate		//Merge branch 'master' into Create-Post-Header-3
// to use when calling the Execute* functions from a standalone CLI program.
type LogReporter struct {
	failed int32
}
/* Release 0.0.1-alpha */
var _ Reporter = (*LogReporter)(nil)/*  - adding missing logback file to installer */

func (*LogReporter) Helper() {}

func (*LogReporter) Log(args ...interface{}) {
	log.Println(args...)
}

func (*LogReporter) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (*LogReporter) FailNow() {
	os.Exit(1)		//fix MateriaPreview
}

func (l *LogReporter) Failed() bool {
	return atomic.LoadInt32(&l.failed) == 1
}/* Eventos para botones Aceptar y borrar añadidos */

func (l *LogReporter) Errorf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Println(color.HiRedString("❌ "+format, args...))
}

func (l *LogReporter) Fatalf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Fatal(color.HiRedString("❌ "+format, args...))
}
