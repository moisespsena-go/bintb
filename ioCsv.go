package bintb

import (
	"encoding/csv"
	"fmt"
)

type CsvRecordReader struct {
	ed *Decoder
	r  *csv.Reader
}

func NewCsvReadDecode(ed *Decoder, r *csv.Reader) *CsvRecordReader {
	return &CsvRecordReader{ed: ed, r: r}
}

func (this *CsvRecordReader) Read() (rec Recorde, err error) {
	var record []string
	if record, err = this.r.Read(); err != nil {
		return
	}

	if len(record) != len(this.ed.Columns) {
		return nil, fmt.Errorf("CsvRecordReader.read: bad csv columns count. Expected %d, but get %d. %v", len(this.ed.Columns), len(record), record)
	}

	rec, err = this.ed.Decode(record...)
	return
}

type CsvRecordWriter struct {
	OnlyColumnNames bool
	enc             *Encoder
	w               *csv.Writer
}

func NewCsvRecordWriter(enc *Encoder, w *csv.Writer) *CsvRecordWriter {
	return &CsvRecordWriter{enc: enc, w: w}
}

func (this *CsvRecordWriter) Write(rec Recorde) (err error) {
	var values []string
	if values, err = this.enc.Encode([]interface{}(rec)...); err != nil {
		return
	}
	return this.w.Write(values)
}