package bintb

import (
	"fmt"
	"io"
	"strconv"
)

type CTuint8 struct{UnLimitColumnTool}

func (CTuint8) Zero() interface{} {
	return uint8(0)
}

func (CTuint8) Decode(value string) (v interface{}, err error) {
	var i uint64
	if i, err = strconv.ParseUint(value, 10, 8); err != nil {
		return
	}
	return uint8(i), nil
}

func (CTuint8) Encode(value interface{}) string {
	return fmt.Sprint(value.(uint8))
}

func (CTuint8) BinRead(r io.Reader) (v interface{}, err error) {
	var i uint8
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTuint8) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

type CTuint16 struct{UnLimitColumnTool}

func (CTuint16) Zero() interface{} {
	return uint16(0)
}

func (CTuint16) Decode(value string) (v interface{}, err error) {
	var i uint64
	if i, err = strconv.ParseUint(value, 10, 16); err != nil {
		return
	}
	return uint16(i), nil
}

func (CTuint16) Encode(value interface{}) string {
	return fmt.Sprint(value.(uint16))
}

func (CTuint16) BinRead(r io.Reader) (v interface{}, err error) {
	var i uint16
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTuint16) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

type CTuint32 struct{UnLimitColumnTool}

func (CTuint32) Zero() interface{} {
	return uint32(0)
}

func (CTuint32) Decode(value string) (v interface{}, err error) {
	var i uint64
	if i, err = strconv.ParseUint(value, 10, 32); err != nil {
		return
	}
	return uint32(i), nil
}

func (CTuint32) Encode(value interface{}) string {
	return fmt.Sprint(value.(uint32))
}

func (CTuint32) BinRead(r io.Reader) (v interface{}, err error) {
	var i uint32
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTuint32) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

type CTuint64 struct{UnLimitColumnTool}

func (CTuint64) Zero() interface{} {
	return uint64(0)
}

func (CTuint64) Decode(value string) (v interface{}, err error) {
	var i uint64
	if i, err = strconv.ParseUint(value, 10, 64); err != nil {
		return
	}
	return i, nil
}

func (CTuint64) Encode(value interface{}) string {
	return fmt.Sprint(value.(uint64))
}

func (CTuint64) BinRead(r io.Reader) (v interface{}, err error) {
	var i uint64
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTuint64) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

const (
	CtUint8  ColumnType = "u8"
	CtUint16 ColumnType = "u16"
	CtUint32 ColumnType = "u32"
	CtUint64 ColumnType = "u64"
)

func init() {
	ColumnTypeTool.Set(CtUint8, CTuint8{}, "uint8")
	ColumnTypeTool.Set(CtUint16, CTuint16{}, "uint16")
	ColumnTypeTool.Set(CtUint32, CTuint32{}, "uint32")
	ColumnTypeTool.Set(CtUint64, CTuint64{}, "uint64")
}
