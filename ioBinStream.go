package bintb

import (
	"io"
)

func BinStreamWriteF(w io.Writer, enc *Encoder, next func() (rec Recorde, err error), maxBufSize ...uint32) (err error) {
	bsw := NewBinStreamWriter(enc, w, maxBufSize...)
	defer func() {
		if err == nil {
			err = bsw.Close()
		} else {
			bsw.Close()
		}
	}()

	var rec Recorde
	for {
		if rec, err = next(); err != nil {
			return
		}
		if rec == nil {
			break
		}

		if err = bsw.Write(rec); err != nil {
			return
		}
	}
	return
}

func BinStreamFullIterate(dec *Decoder, r io.Reader, opt ...*StreamReaderOptions) error {
	return OpenIterate(NewBinStreamReader(dec, r, opt...))
}
