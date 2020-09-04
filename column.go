package bintb

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type ColumnType string

type ColumnValidator = func(ctx context.Context, value interface{}) error

type Flags uint8

const (
	Flag    Flags = 0
	Required Flags = 1 << iota
	AllowBlank
	Unique = Required | 4
)

func (f Flags) Set(flag Flags) Flags    { return f | flag }
func (f Flags) Clear(flag Flags) Flags  { return f &^ flag }
func (f Flags) Toggle(flag Flags) Flags { return f ^ flag }
func (f Flags) Has(flag Flags) bool     { return f&flag != 0 }

type Column struct {
	Name       string
	typ        ColumnType
	Tool       ColumnTool
	vt         reflect.Type
	maxLength  uint16
	flags      Flags
	validators []ColumnValidator
	Data map[interface{}]interface{}
}

func NewColumn(name string, typ ColumnType, flags ...Flags) *Column {
	tool := ColumnTypeTool.Get(typ)
	vt := reflect.TypeOf(tool.Zero())
	var _flags Flags
	for _, f := range flags {
		_flags = _flags.Set(f)
	}
	// disable null for slices and string
	switch vt.Kind() {
	case reflect.Slice, reflect.String:
		_flags.Set(AllowBlank)
		_flags.Set(Required)
	}
	if _flags.Has(Unique) {
		_flags = _flags.Clear(AllowBlank).Set(Required)
	}
	return &Column{
		Name:  name,
		typ:   typ,
		Tool:  tool,
		vt:    vt,
		flags: _flags,
	}
}

func (this *Column) MaxLength() uint16 {
	return this.maxLength
}

func (this *Column) SetMaxLength(maxLength uint16) {
	this.maxLength = maxLength
}

func (this *Column) Validator(f ...ColumnValidator) {
	this.validators = append(this.validators, f...)
}

func (this *Column) Validators() []ColumnValidator {
	return this.validators
}

func (this *Column) Type() ColumnType {
	return this.typ
}

func (this *Column) ValueType() reflect.Type {
	return this.vt
}

func (this *Column) String() string {
	var r string
	if this.maxLength > 0 {
		r = "[" + fmt.Sprint(this.maxLength) + "]"
	}
	r += this.Name
	if this.flags.Has(Unique) {
		r += "+"
	} else if this.flags.Has(Required) && !this.flags.Has(AllowBlank) {
		r += "*"
	}
	r += ":" + string(this.typ)
	return r
}

func (this *Column) Required() bool {
	return this.flags.Has(Required)
}

func (this *Column) Unique() bool {
	return this.flags.Has(Unique)
}

func (this *Column) AllowBlank() bool {
	return this.flags.Has(AllowBlank)
}

func (this *Column) Validate(ctx context.Context, value interface{}) (err error) {
	if value == nil {
		return nil
	}
	if this.maxLength > 0 {
		switch vt := value.(type) {
		case []byte:
			if len(vt) > int(this.maxLength) {
				return errors.New("value maxLength exceeded.")
			}
		case string:
			if len(vt) > int(this.maxLength) {
				return errors.New("value maxLength exceeded.")
			}
		}
	}

	for _, vldr := range this.validators {
		if err = vldr(ctx, value); err != nil {
			return
		}
	}
	return
}
