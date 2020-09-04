package bintb

import (
	"io"
	"strconv"
)

type CTint8 struct{ UnLimitColumnTool }

func (CTint8) Zero() interface{} {
	return int8(0)
}

func (CTint8) Decode(value string) (v interface{}, err error) {
	var i int
	if i, err = strconv.Atoi(value); err != nil {
		return
	}
	return int8(i), nil
}

func (CTint8) Encode(value interface{}) string {
	return strconv.Itoa(int(value.(int8)))
}

func (CTint8) BinRead(r io.Reader) (v interface{}, err error) {
	var i int8
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTint8) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

type CTint16 struct{ UnLimitColumnTool }

func (CTint16) Zero() interface{} {
	return int16(0)
}

func (CTint16) Decode(value string) (v interface{}, err error) {
	var i int
	if i, err = strconv.Atoi(value); err != nil {
		return
	}
	return int16(i), nil
}

func (CTint16) Encode(value interface{}) string {
	return strconv.Itoa(int(value.(int16)))
}

func (CTint16) BinRead(r io.Reader) (v interface{}, err error) {
	var i int16
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTint16) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

type CTint32 struct{ UnLimitColumnTool }

func (CTint32) Zero() interface{} {
	return int32(0)
}

func (CTint32) Decode(value string) (v interface{}, err error) {
	var i int
	if i, err = strconv.Atoi(value); err != nil {
		return
	}
	return int32(i), nil
}

func (CTint32) Encode(value interface{}) string {
	return strconv.Itoa(int(value.(int32)))
}

func (CTint32) BinRead(r io.Reader) (v interface{}, err error) {
	var i int32
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTint32) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

type CTint64 struct{ UnLimitColumnTool }

func (CTint64) Zero() interface{} {
	return int64(0)
}

func (CTint64) Decode(value string) (v interface{}, err error) {
	var i int
	if i, err = strconv.Atoi(value); err != nil {
		return
	}
	return int64(i), nil
}

func (CTint64) Encode(value interface{}) string {
	return strconv.Itoa(int(value.(int64)))
}

func (CTint64) BinRead(r io.Reader) (v interface{}, err error) {
	var i int64
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTint64) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

const (
	CtInt8  ColumnType = "i8"
	CtInt16 ColumnType = "i16"
	CtInt32 ColumnType = "i32"
	CtInt64 ColumnType = "i64"
)

func init() {
	ColumnTypeTool.Set(CtInt8, CTint8{}, "int8")
	ColumnTypeTool.Set(CtInt16, CTint16{}, "int16")
	ColumnTypeTool.Set(CtInt32, CTint32{}, "int32")
	ColumnTypeTool.Set(CtInt64, CTint64{}, "int64")
}
