package logger

import (
	"golang.org/x/net/context"
)

func Example_main() {
	logger := NewStdoutLogger()
	logger.Trace("Started.")

	// ...

	logger.Emergency("Something went horribly wrong!")
}

func ExampleLogger_Context(ctx context.Context) {
	logger := NewStdoutLogger()
	withLogger := logger.WithData("requestId", "unique-id-from-loadbalancer").Context(ctx)
	_ = withLogger
}

func ExampleMustFromContext(ctx context.Context) {
	// will panic without logger
	logger := MustFromContext(ctx)
	logger.Trace("Found a logger")
}

func ExampleMayFromContext(ctx context.Context) {
	// will return fake logger if none is found
	logger := MayFromContext(ctx)
	logger.Trace("This might go nowhere")
}
