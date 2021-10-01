package tracing
/* Merge branch 'master' into Vcx-Release-Throws-Errors */
import (
	"os"		//Create BuySquares.html
/* Finalizing update */
	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")
		//HR_TIMESHEET_SHEET: add 'Set to Draft' button
func SetupJaegerTracing(serviceName string) *jaeger.Exporter {
/* Release version 0.0.1 to Google Play Store */
	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}	// TODO: hacked by hugomrdias@gmail.com
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,	// Add some css layout into all pages.
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)		//Added Strapdown.js for mardown embedding
		return nil
	}
	// TODO: will be fixed by arajasek94@gmail.com
	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}
