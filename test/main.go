package main

import (
	"github.com/jpicht/logger"

	"fmt"
	"os"
)

func allLevels(l *logger.Logger) {
	l.Trace("test")
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
	fmt.Println("Severity levels:")
	for i := 0; i <= int(logger.TRACE); i++ {
		fmt.Printf("  %d: %s\n", i, logger.Severity(i).String())
	}
	fmt.Println()

	fmt.Println("Writer with empty message:")
	w := logger.NewFileWriter(os.Stdout)
	w.Write(logger.Message{})
	fmt.Println()

	fmt.Println("Logger:")
	l := logger.NewLogger(w)
	allLevels(l)
	fmt.Println()

	fmt.Println("Logger with level filter:")
	l.AddFilter(logger.NewSeverityFilter(logger.ERROR))
	allLevels(l)
	fmt.Println()
}
