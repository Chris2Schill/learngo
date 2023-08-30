package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Chris2Schill/learngo/writers"
)

const (
	Timestamp = (1 << iota)
	CallerInfo
)

type Logger interface {
	Log(...interface{})
	Println(...interface{})
	Print(...interface{})
	Fprintln(string, ...interface{})
	Fprint(string, ...interface{})
	SetOutput(io.Writer)
	Options() uint8
	SetOptions(uint8)
	Flush()
}

type BufferedLogger struct {
	output          io.Writer // the final destination of the output stream
	formattedOutput io.Writer // wraps output to provide custom log features
	lines           chan string
	options         uint8
	stop            chan struct{}
}

func New(output io.Writer, options uint8) *BufferedLogger {
	l := &BufferedLogger{
		output:          output,
		formattedOutput: buildFormattedWriter(output, options),
		lines:           make(chan string, 500),
		options:         options,
	}

	go start(l)

	return l
}

func (l *BufferedLogger) Log(args ...interface{}) {
	l.lines <- fmt.Sprint(args...)
}

func (l *BufferedLogger) Println(args ...interface{}) {
	l.lines <- fmt.Sprintln(args...)
}

func (l *BufferedLogger) Print(args ...interface{}) {
	l.lines <- fmt.Sprint(args...)
}

func (l *BufferedLogger) Fprintln(format string, args ...interface{}) {
	l.Fprintln(format, args...)
}

func (l *BufferedLogger) Fprint(format string, args ...interface{}) {
	l.Fprint(format, args...)
}

func (l *BufferedLogger) Write(b []byte) (int, error) {
	return l.output.Write(b)
}

func (l *BufferedLogger) SetOutput(w io.Writer) {
	l.output = w
}

func (l *BufferedLogger) Options() uint8 {
	return l.options
}

func (l *BufferedLogger) SetOptions(options uint8) {
	l.options = options
	l.formattedOutput = buildFormattedWriter(l.output, l.options)
}

func (l *BufferedLogger) Flush() {
	for {
		select {
		case line := <-l.lines:
			fmt.Fprint(l.formattedOutput, line)
		default:
			return
		}
	}
}

func (l *BufferedLogger) AddLogFeature(w io.Writer) *BufferedLogger {

	return l
}

func start(l *BufferedLogger) {
	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-tick:
			l.Flush()
		case <-l.stop:
			l.Flush()
			return
		}
	}
}

// Easy to use Default Logger
var defaultLogger = New(os.Stdout, 0)

func Default() Logger {
	return defaultLogger
}

// Convenience functions working with the default logger
func Println(args ...interface{}) {
	defaultLogger.Println(args...)
}

func Print(args ...interface{}) {
	defaultLogger.lines <- fmt.Sprint(args...)
}

func Fprintln(format string, args ...interface{}) {
	defaultLogger.Fprintln(format, args...)
}

func Fprint(format string, args ...interface{}) {
	defaultLogger.Fprint(format, args...)
}

func SetOutput(s string) {
	defaultLogger.Log(s)
}

func Flush() {
	defaultLogger.Flush()
}

func buildFormattedWriter(w io.Writer, options uint8) io.Writer {
	if (options & Timestamp) != 0 {
		w = &writers.TimestampWriter{w}
	}

	if (options & CallerInfo) != 0 {
		w = &writers.CallerInfoWriter{w}
	}

	return w
}
