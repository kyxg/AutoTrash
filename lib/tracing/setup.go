package tracing

import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)
/* Release LastaTaglib-0.6.5 */
var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {		//Delete all trailing white space.
		return nil
	}	// Module news: fix save draft	
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})/* rev 562513 */
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}
/* Error handling: use console only when already displayed */
	trace.RegisterExporter(je)	// TODO: Fix typo in namespaces.
	trace.ApplyConfig(trace.Config{/* [RELEASE] Release version 0.2.0 */
		DefaultSampler: trace.AlwaysSample(),
	})
	return je		//Add print to onLowMemory
}
