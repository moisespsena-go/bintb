package bintb

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/go-errors/errors"
)

type Encoder struct {
	Columns []*Column
	Context context.Context
}

func NewEncoder(columns []*Column, ctx ...context.Context) *Encoder {
	var ctx_ = context.Background()
	for _, ctx := range ctx {
		if ctx != nil {
			ctx_ = ctx
		}
	}
	return &Encoder{Columns: columns, Context: ctx_}
}

func (this *Encoder) Encode(values ...interface{}) (rec []string, err error) {
	rec = make([]string, len(values))
	for i, value := range values[0:len(this.Columns)] {
		c := this.Columns[i]
		if value == nil {
			if !c.Required() || c.AllowBlank() {
				rec[i] = ""
				continue
			}
			return nil, fmt.Errorf("Decoder.Encode: values[%d as %q] nil value", i, c.Name)
		} else if err = c.Validate(this.Context, value); err != nil {
			return nil, errors.WrapPrefix(err, fmt.Sprintf("Decoder.Encode: column[%d as %q] validation", i, c.Name), 1)
		}
		rec[i] = c.Tool.Encode(value)
	}
	return
}

func (this *Encoder) BinWrite(w io.Writer, values ...interface{}) (err error) {
	var wb bytes.Buffer
	for i, value := range values[0:len(this.Columns)] {
		c := this.Columns[i]
		if !c.Required() {
			if err = bw(&wb, value == nil); err != nil {
				return errors.WrapPrefix(err, fmt.Sprintf("Decoder.BinWrite: values[%d as %q] null byte", i, c.Name), 1)
			} else if value == nil {
				continue
			}
		} else if value == nil {
			value = c.Tool.Zero()
		}
		if err = c.Validate(this.Context, value); err != nil {
			return errors.WrapPrefix(err, fmt.Sprintf("Decoder.BinWrite: column[%d as %q] validation", i, c.Name), 1)
		}
		if err = c.Tool.BinWrite(&wb, value); err != nil {
			return errors.WrapPrefix(err, fmt.Sprintf("Decoder.BinWrite: values[%d as %q]", i, c.Name), 1)
		}
	}
	_, err = w.Write(wb.Bytes())
	return
}
