package logger

import (
	ctx "context"

	"os"
	"testing"
)

func TestContext(t *testing.T) {
	logger := NewLogger(NewFileWriter(os.Stdout, NewStringEncoder(), Separators.NewLine))
	with := logger.Context(ctx.Background())
	from := MustFromContext(with)
	if logger != from {
		t.Fatal("Logger changed while inside context")
	}
}
