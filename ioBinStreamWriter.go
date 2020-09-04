package bintb

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
)

type BinStreamWriter struct {
	encoder        *Encoder
	w              io.Writer
	failed         bool
	recordsCount   int64
	headersSent    bool
	buf            bytes.Buffer
	bufRecordCount uint16
	MaxBufferSize  uint32
}

func NewBinStreamWriter(encoder *Encoder, w io.Writer, maxBufSize ...uint32) *BinStreamWriter {
	if len(maxBufSize) == 0 {
		maxBufSize = append(maxBufSize, 0)
	}
	return &BinStreamWriter{encoder: encoder, w: w, MaxBufferSize: maxBufSize[0]}
}

func (this *BinStreamWriter) Columns() []*Column {
	return this.encoder.Columns
}

func (this *BinStreamWriter) Encoder() *Encoder {
	return this.encoder
}

func (this *BinStreamWriter) Failed() bool {
	return this.failed
}

func (this *BinStreamWriter) RecordsCount() int64 {
	return this.recordsCount
}

func (this *BinStreamWriter) write(p []byte) error {
	_, err := this.w.Write(p)
	return err
}

func (this *BinStreamWriter) Flush() (err error) {
	if !this.headersSent {
		var b [2]byte
		binOrder.PutUint16(b[:], uint16(len(this.encoder.Columns)))
		buf := bytes.NewBuffer(b[:])
		bcw := NewBinColumnWriter(buf)
		for _, c := range this.encoder.Columns {
			bcw.Write(c)
		}
		if err = this.write(buf.Bytes()); err != nil {
			return errors.Wrap(err, "send headers")
		}
		this.headersSent = true
	}
	if this.buf.Len() > 0 {
		b := this.buf.Bytes()
		b = append(b, 0, 0)
		copy(b[2:], b)
		binOrder.PutUint16(b, this.bufRecordCount)
		err = this.write(b)
		this.buf.Reset()
		this.bufRecordCount = 0
	}
	return
}

func (this *BinStreamWriter) Write(rec Recorde) (err error) {
	var buf bytes.Buffer
	if err = this.encoder.BinWrite(&buf, rec...); err != nil {
		return
	}
	if this.MaxBufferSize > 0 && (this.buf.Len()+buf.Len()) > int(this.MaxBufferSize) {
		err = this.Flush()
		this.buf = buf
	} else {
		this.buf.Write(buf.Bytes())
	}
	this.recordsCount++
	this.bufRecordCount++
	return nil
}

func (this *BinStreamWriter) Close() error {
	if !this.failed {
		if err := this.Flush(); err != nil {
			return err
		}
		// done
		err := this.write([]byte{0, 0})
		return err
	}
	return nil
}
