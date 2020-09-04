package bintb

import (
	"encoding/json"
)

type JsonWriterConfig struct {
	Encoder           *Encoder
	OnlyNames, Arrays bool
	JosonEncoder      *json.Encoder
}

type JsonRecordWriter struct {
	*JsonWriterConfig
}

func NewJsonRecordWriter(config *JsonWriterConfig) *JsonRecordWriter {
	return &JsonRecordWriter{config}
}

func (this *JsonRecordWriter) Write(rec Recorde) (err error) {

	if this.Arrays {
		return this.JosonEncoder.Encode(rec)
	}
	var data = map[string]interface{}{}
	if this.OnlyNames {
		for i, col := range this.Encoder.Columns {
			data[col.Name] = rec[i]
		}
	} else {
		for i, col := range this.Encoder.Columns {
			data[col.String()] = rec[i]
		}
	}
	return this.JosonEncoder.Encode(data)
}
