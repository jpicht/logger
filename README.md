# Logger for golang

The goal of this little library is to provide a nice and easy to use logger with severity handling, flexible output and "context" handling. Currently the interface provides two convenience functions for each severity. One for logging plain strings, the other one with a Printf-like interface.

In contrast to the build-in "log" module, this library uses a logger instance, that needs to be passed around. Apart from the obvious drawback of needing to pass it around this provides the ability to inject a logger instance with additional information attached.

In the near future I plan to implement support for golang.org/x/net/context.

Currently you need to instanciate a log encoder, a log writer and the actual logger object

## Usage

```go
var (
	le logger.Encoder
	lw logger.Writer
	l  logger.Logger
)

le = logger.NewStringEncoder()
lw = logger.NewFileWriter(os.Stdout, logger.NewStringEncoder(), logger.Seperators.NewLine)
l = logger.NewLogger(w)
```


## Motivation

I was very unhappy with all loggers I could find for the go language, as there wasn't a single one with usable severity handling. I've seen a lot of go programs writing "[WARN]" etc. as part of the literal log message which I find very appalling, because it introduces the need to parse strings if you wan't to do some automatic processing and/or filtering of log messages, or want to use something like the ELK stack.
