package writers

import (
	"io"
	"runtime"
)

type CallerInfoWriter struct {
	W io.Writer
}

func (w *CallerInfoWriter) Write(b []byte) (int, error) {
	pc, file, line, ok := runtime.Caller(4)
	_ = pc
	_ = line
	_ = ok

	// Thers probably a better way to do this?
	b2 := append([]byte(file), " | "...)
	b2 = append(b2, b...)
	return w.W.Write(b2)
}
