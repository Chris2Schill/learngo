package writers

import (
	"io"
)

type BoomBabyWriter struct {
	W io.Writer
}

func (w *BoomBabyWriter) Write(b []byte) (n int, err error) {
	s := []byte("Boombaby! ")
	b2 := append(s, b...)
	n, err = w.W.Write(b2)
	return
}
