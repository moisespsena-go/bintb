package bintb

type JsonStreamWriterConfig struct {
	JsonWriterConfig
	OnlyRecords bool
	IndentLevel int
}

type JsonStreamWriter struct {
	Config       *JsonStreamWriterConfig
	failed       bool
	recordsCount int64
	headersSent  bool
	recWriter    *JsonRecordWriter
}

func NewJsonStreamWriter(config *JsonStreamWriterConfig) *JsonStreamWriter {
	return &JsonStreamWriter{Config: config, recWriter: NewJsonRecordWriter(&config.JsonWriterConfig)}
}

func (this *JsonStreamWriter) Columns() []*Column {
	return this.Config.Encoder.Columns
}

func (this *JsonStreamWriter) Encoder() *Encoder {
	return this.Config.Encoder
}

func (this *JsonStreamWriter) Failed() bool {
	return this.failed
}

func (this *JsonStreamWriter) RecordsCount() int64 {
	return this.recordsCount
}

func (this *JsonStreamWriter) sendHeaders() (err error) {
	defer func() {
		if err != nil {
			this.failed = true
		}
	}()
	if !this.headersSent {
		if !this.Config.OnlyRecords {
			var columns = make([]string, len(this.Columns()))
			for i, col := range this.Columns() {
				if this.Config.OnlyNames {
					columns[i] = col.Name
				} else {
					columns[i] = col.String()
				}
			}
			if err = this.Config.JosonEncoder.Encode(columns); err != nil {
				return
			}
		}
		this.headersSent = true
	}
	return
}

func (this *JsonStreamWriter) Write(rec Recorde) (err error) {
	if err = this.sendHeaders(); err != nil {
		return
	}

	if err = this.recWriter.Write(rec); err != nil {
		return
	}

	this.recordsCount++
	return
}

func (this *JsonStreamWriter) Close() (err error) {
	if !this.failed {
		if err = this.sendHeaders(); err != nil {
			return
		}
	}
	return nil
}
