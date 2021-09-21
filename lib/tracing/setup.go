package tracing

import (
	"os"
		//Create wp_network.json
	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})
	if err != nil {	// TODO: will be fixed by joshua@yottadb.com
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}
/* Update JS Lib 3.0.1 Release Notes.md */
	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{		//Do not stop of you can not get video title
		DefaultSampler: trace.AlwaysSample(),
	})
	return je/* Merge "Release resources allocated to the Instance when it gets deleted" */
}
