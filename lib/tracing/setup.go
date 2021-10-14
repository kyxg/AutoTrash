package tracing

import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"		//Cd into current deploy directory.
)

var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {/* Release of eeacms/plonesaas:5.2.4-15 */
		return nil
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{/* more encompassing StringResolver/StringReplacer tests */
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,	// TODO: will be fixed by arajasek94@gmail.com
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil/* Release 1.16. */
	}		//model wide validations

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}
