package tracing

import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)
/* rev 839947 */
var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {
		//[panel] use super+shift+<number> to launch new instance of an application
	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")/* Release 13.0.0 */

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,		//crop optmization
		ServiceName:   serviceName,		//Add solution to #9 Palindrome Number
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}	// Update Network Diagnostic Instructions

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{/* Release TomcatBoot-0.3.4 */
		DefaultSampler: trace.AlwaysSample(),
	})	// TODO: PreRelease metadata cleanup.
	return je
}
