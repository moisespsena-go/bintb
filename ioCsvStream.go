package bintb

import (
	"encoding/csv"
)

func CsvStreamWriteF(opt CsvStreamWriterOptions, w *csv.Writer, next func() (rec Recorde, err error)) (err error) {
	csw := NewCsvStreamWriter(opt, w)

	defer func() {
		if err == nil {
			err = csw.Close()
		} else {
			csw.Close()
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

		if err = csw.Write(rec); err != nil {
			return
		}
	}
	return
}

func CsvStreamFullIterate(r *csv.Reader, opt ...*CsvStreamReaderOptions) error {
	return OpenIterate(NewCsvStreamReader(r, opt...))
}
