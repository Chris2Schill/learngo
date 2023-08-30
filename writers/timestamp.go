package writers

import (
	"io"
	"time"
)

type TimestampWriter struct {
	W io.Writer
}

func (w *TimestampWriter) Write(b []byte) (int, error) {
	timeStr := time.Now().Format(time.ANSIC)

	// Thers probably a better way to do this?
	b2 := append([]byte(timeStr), " | "...)
	b2 = append(b2, b...)
	return w.W.Write(b2)
}
