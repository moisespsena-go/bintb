package bintb

import "encoding/csv"

type CsvStreamWriterOptions struct {
	Encoder         *Encoder
	OnlyColumnNames bool
	OnlyRecords     bool

	BeforeRecordeWrite,
	AfterRecordWrite func(w *CsvStreamWriter, rec Recorde) (err error)
}

type CsvStreamWriter struct {
	OnlyColumnNames bool
	OnlyRecords     bool
	enc             *Encoder
	w               *csv.Writer
	failed          bool
	recordsCount    int64
	headersSent     bool
	last            Recorde

	BeforeRecordeWrite,
	AfterRecordWrite func(w *CsvStreamWriter, rec Recorde) (err error)
}

func NewCsvStreamWriter(opt CsvStreamWriterOptions, w *csv.Writer) *CsvStreamWriter {
	return &CsvStreamWriter{
		w:                  w,
		enc:                opt.Encoder,
		OnlyRecords:        opt.OnlyRecords,
		OnlyColumnNames:    opt.OnlyColumnNames,
		AfterRecordWrite:   opt.AfterRecordWrite,
		BeforeRecordeWrite: opt.BeforeRecordeWrite,
	}
}

func (this *CsvStreamWriter) CsvWriter() *csv.Writer {
	return this.w
}

func (this *CsvStreamWriter) Last() Recorde {
	return this.last
}

func (this *CsvStreamWriter) Columns() []*Column {
	return this.enc.Columns
}

func (this *CsvStreamWriter) Encoder() *Encoder {
	return this.enc
}

func (this *CsvStreamWriter) Failed() bool {
	return this.failed
}

func (this *CsvStreamWriter) RecordsCount() int64 {
	return this.recordsCount
}

func (this *CsvStreamWriter) sendHeaders() (err error) {
	defer func() {
		if err != nil {
			this.failed = true
		}
	}()
	if !this.headersSent {
		if !this.OnlyRecords {
			var columns = make([]string, len(this.enc.Columns))
			for i, col := range this.enc.Columns {
				if this.OnlyColumnNames {
					columns[i] = col.Name
				} else {
					columns[i] = col.String()
				}
			}

			err = this.w.Write(columns)
		}
		this.headersSent = true
	}
	return
}

func (this *CsvStreamWriter) Write(rec Recorde) (err error) {
	if err = this.sendHeaders(); err != nil {
		return
	}

	if this.BeforeRecordeWrite != nil {
		if err = this.BeforeRecordeWrite(this, rec); err != nil {
			return
		}
	}

	var values = make([]string, len(this.enc.Columns))
	for i, col := range this.enc.Columns {
		values[i] = col.Tool.Encode(rec[i])
	}

	if err = this.w.Write(values); err != nil {
		return
	}
	this.w.Flush()

	if this.AfterRecordWrite != nil {
		if err = this.AfterRecordWrite(this, rec); err != nil {
			return
		}
	}

	this.recordsCount++
	this.last = rec
	return
}

func (this *CsvStreamWriter) Close() (err error) {
	if !this.failed {
		if err = this.sendHeaders(); err != nil {
			return
		}
		this.w.Flush()
	}
	return nil
}
