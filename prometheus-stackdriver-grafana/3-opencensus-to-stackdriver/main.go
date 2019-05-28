package main

import (
	"log"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource"
	"github.com/kr/pretty"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"google.golang.org/api/option"
)

func main() {
	sd, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: "tommyde-191320",
		// MetricPrefix helps uniquely identify your metrics.
		MetricPrefix: "demo-prefix",

		OnError: func(err error) {
			pretty.Println(err)
		},

		MonitoringClientOptions: []option.ClientOption{
			option.WithCredentialsFile("service-account.json"),
		},

		TraceClientOptions: []option.ClientOption{
			option.WithCredentialsFile("service-account.json"),
		},

		MonitoredResource: monitoredresource.Autodetect(),
	})
	if err != nil {
		log.Fatalf("Failed to create the Stackdriver exporter: %v", err)
	}
	// It is imperative to invoke flush before your main function exits
	defer sd.Flush()

	// Register it as a metrics exporter
	view.RegisterExporter(sd)
	trace.RegisterExporter(sd)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	view.SetReportingPeriod(60 * time.Second)

	//In our main, register ochttp Server views
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Fatalf("Failed to register server views for HTTP metrics: %v", err)
	}
	// Register it as a trace exporter

	//exporter := &exporter.PrintExporter{}
	// trace.RegisterExporter(exporter)
	// trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	// view.RegisterExporter(exporter)

	// view.SetReportingPeriod(1 * time.Second)

	// In our main, register ochttp Server views

	ServerRequestTest := &view.View{
		Name:        "tommyd/test",
		Description: "Server request count by HTTP method",
		TagKeys:     []tag.Key{ochttp.Method, ochttp.StatusCode, ochttp.KeyServerRoute},
		Measure:     ochttp.ServerLatency,
		Aggregation: view.Count(),
	}

	if err := view.Register(ServerRequestTest); err != nil {
		log.Fatalf("Failed to register server views for HTTP metrics: %v", err)
	}

	mux := http.NewServeMux()

	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "example.com/ProcessVideo")
		defer span.End()

		ochttp.SetRoute(ctx, "/users")
		w.Write([]byte("Hello, World!"))
	})

	originalHandlerV2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "example.com/ProcessVideo")
		defer span.End()

		ochttp.SetRoute(ctx, "/v2/users")
		w.WriteHeader(400)
		w.Write([]byte("Hello, World!"))
	})

	mux.Handle("/users", originalHandler)
	mux.Handle("/v2/users", originalHandlerV2)

	och := &ochttp.Handler{
		Handler: mux, // The handler you'd have used originally
	}

	// Now use the instrumnted handler
	if err := http.ListenAndServe(":9999", och); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
