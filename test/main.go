package main

import (
	"github.com/jpicht/logger"

	"fmt"
	"os"
)

func allLevels(l *logger.Logger) {
	l.Tracef("test: %T", *l)
	l.Debug("test")
	l.Info("test")
	l.Notice("test")
	l.Warning("test")
	l.Error("test")
	l.Critical("test")
	l.Alert("test")
	l.Emergency("test")
}

func main() {
	var w logger.Writer
	fmt.Println("Severity levels:")
	for i := 0; i <= int(logger.TRACE); i++ {
		fmt.Printf("  %d: %s\n", i, logger.Severity(i).String())
	}
	fmt.Println()

	fmt.Println("Writer with empty message:")
	w = logger.NewFileWriter(os.Stdout)
	w.Write(logger.Message{})
	fmt.Println()

	fmt.Println("Logger:")
	l := logger.NewLogger(w)
	allLevels(l)
	fmt.Println()

	fmt.Println("Logger with level filter:")
	wf := logger.NewSeverityFilter(logger.ERROR, w)
	l = logger.NewLogger(wf)
	allLevels(l)
	fmt.Println()

	fmt.Println("Logger with prefix:")
	l = logger.NewLogger(w).WithPrefix("--PREFIX-- ")
	allLevels(l)
	fmt.Println()

	fmt.Println("Logger with data:")
	l = logger.NewLogger(w).WithData("foo", "bar")
	allLevels(l)
	fmt.Println()
}
