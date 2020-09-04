package bintb

import (
	"context"
	"fmt"
	"io"

	"github.com/go-errors/errors"
)

type Decoder struct {
	Columns []*Column
	Context context.Context
}

func NewDecoder(columns []*Column, ctx ...context.Context) *Decoder {
	var ctx_ = context.Background()
	for _, ctx := range ctx {
		if ctx != nil {
			ctx_ = ctx
		}
	}
	return &Decoder{Columns: columns, Context: ctx_}
}

func (this *Decoder) Decode(values ...string) (rec []interface{}, err error) {
	rec = make([]interface{}, len(values))
	for i, value := range values[0:len(this.Columns)] {
		c := this.Columns[i]
		if value == "" {
			if !c.Required() {
				continue
			} else if c.AllowBlank() {
				rec[i] = c.Tool.Zero()
				continue
			}
			return nil, fmt.Errorf("Decoder.Decode: values[%d as %q] empty value", i, c.Name)
		} else if err = c.Validate(this.Context, value); err != nil {
			return nil, errors.WrapPrefix(err, fmt.Sprintf("Decoder.Decode: column[%d as %q] validation", i, c.Name), 1)
		}
		if rec[i], err = c.Tool.Decode(value); err != nil {
			return nil, errors.WrapPrefix(err, fmt.Sprintf("Decoder.Decode: values[%d as %q]", i, c.Name), 1)
		}
	}
	return
}

func (this *Decoder) BinRead(r io.Reader) (rec []interface{}, err error) {
	rec = make([]interface{}, len(this.Columns))
	var value interface{}
	for i, c := range this.Columns {
		if !c.Required() {
			var null bool
			if err = br(r, &null); err != nil {
				return nil, errors.WrapPrefix(err, fmt.Sprintf("Decoder.BinRead: values[%d as %q] null byte", i, c.Name), 1)
			} else if null {
				continue
			}
		}
		if value, err = c.Tool.BinRead(r); err != nil {
			return nil, errors.WrapPrefix(err, fmt.Sprintf("Decoder.BinRead: column[%d as %q]", i, c.Name), 1)
		} else if err = c.Validate(this.Context, value); err != nil {
			return nil, errors.WrapPrefix(err, fmt.Sprintf("Decoder.BinRead: column[%d as %q] validation", i, c.Name), 1)
		}
		rec[i] = value
	}
	return
}

func (this *Decoder) BinReadLoopN(r io.Reader, n int, onRec func(rec Recorde) error) error {
	for i := 0; i < n; i++ {
		rec, err := this.BinRead(r)
		if err != nil {
			return errors.WrapPrefix(err, fmt.Sprintf("Decoder.BinReadLoopN: record %d", i), 1)
		}
		if err = onRec(rec); err != nil {
			return err
		}
	}
	return nil
}
