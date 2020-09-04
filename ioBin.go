package bintb

import (
	"io"
)

type BinRecordReader struct {
	decoder *Decoder
	r       io.Reader
}

func NewBinRecordReader(ed *Decoder, r io.Reader) *BinRecordReader {
	return &BinRecordReader{decoder: ed, r: r}
}

func (this *BinRecordReader) Read() (rec Recorde, err error) {
	return this.decoder.BinRead(this.r)
}

type BinRecordWriter struct {
	encoder *Encoder
	w       io.Writer
}

func NewBinRecordWriter(ed *Encoder, w io.Writer) *BinRecordWriter {
	return &BinRecordWriter{encoder: ed, w: w}
}

func (this *BinRecordWriter) Write(rec Recorde) (err error) {
	return this.encoder.BinWrite(this.w, rec...)
}

type BinColumnReader struct {
	r io.Reader
}

func (this *BinColumnReader) Read() (col *Column, err error) {
	var b [1]byte
	if _, err = this.r.Read(b[:]); err != nil {
		return
	}
	var nameb = make([]byte, int(b[0]))
	if _, err = this.r.Read(nameb); err != nil {
		return
	}
	return ParseColumn(string(nameb))
}

type BinColumnWriter struct {
	w io.Writer
}

func NewBinColumnWriter(w io.Writer) *BinColumnWriter {
	return &BinColumnWriter{w: w}
}

func (this *BinColumnWriter) Write(col *Column) (err error) {
	cols := col.String()
	var buf = append([]byte{byte(len(cols))}, []byte(cols)...)
	_, err = this.w.Write(buf)
	return
}
