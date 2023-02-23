package tracer

import (
	"context"
	logexporter "github.com/DerryRenaldy/learnFiber/pkg/tracer/exporter"
	"github.com/DerryRenaldy/logger/logger"
	"time"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type tracerImpl struct {
	tp     trace.TracerProvider
	tr     trace.Tracer
	enable bool
}

var DefaultTracer *tracerImpl

// New creates a new tracer
func New(enable bool, projectID string, TracerName string, logger logger.ILogger) *tracerImpl {
	if !enable {
		lExporter := logexporter.New(logger)
		tp := sdktrace.NewTracerProvider(
			// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
			// In a production application, use sdktrace. ProbabilitySampler with a desired probability.
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(lExporter),
		)
		otel.SetTracerProvider(tp)
		tr := otel.Tracer(TracerName)
		DefaultTracer = &tracerImpl{
			tp:     tp,
			tr:     tr,
			enable: enable,
		}
		return DefaultTracer
	}
	// Create log exporter
	lExporter := logexporter.New(logger)

	// Create Google Cloud Trace exporter to be able to retrieve the collected spans
	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		logger.Fatalf("Error creating new exporter: %v", err)
	}

	ctx := context.Background()

	// Identify your application using resource detection
	res, err := resource.New(ctx,
		// Use the GCP resource detector to detect information about the GCP platform
		resource.WithDetectors(gcp.NewDetector()),
		// Keep the default detectors
		resource.WithTelemetrySDK(),
		// Add your own custom attributes to identify your application
		resource.WithAttributes(
			semconv.ServiceNameKey.String(TracerName),
		),
	)
	if err != nil {
		logger.Fatalf("resource.New: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
		// In a production application, use sdktrace. ProbabilitySampler with a desired probability.
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithBatcher(lExporter),
		sdktrace.WithResource(res),
	)
	defer tp.ForceFlush(ctx)
	otel.SetTracerProvider(tp)
	tr := otel.Tracer(TracerName)
	DefaultTracer = &tracerImpl{
		tp:     tp,
		tr:     tr,
		enable: enable,
	}
	return DefaultTracer
}

func Shutdown() {
	if DefaultTracer == nil || !DefaultTracer.enable {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = DefaultTracer.tp.(*sdktrace.TracerProvider).Shutdown(ctx)
}

// StartSpan creates a span and context from exist context
func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	if DefaultTracer == nil {
		tp := trace.NewNoopTracerProvider()
		DefaultTracer = &tracerImpl{
			tp:     tp,
			tr:     tp.Tracer(""),
			enable: false,
		}
	}
	return DefaultTracer.tr.Start(ctx, name)
}

// SetSpanAttributes Sets span attributes for input key-value pairs
func SetSpanAttributes(span trace.Span, input map[string]string) {
	for key, value := range input {
		span.SetAttributes(attribute.String(key, value))
	}
}
