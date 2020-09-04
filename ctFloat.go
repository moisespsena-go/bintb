package bintb

import (
	"fmt"
	"io"
	"strconv"
)

type CTFloat32 struct {
	UnLimitColumnTool
}

func (CTFloat32) Zero() interface{} {
	return float32(0)
}

func (CTFloat32) Decode(value string) (v interface{}, err error) {
	var f float64
	if f, err = strconv.ParseFloat(value, 32); err != nil {
		return
	}
	return float32(f), nil
}

func (CTFloat32) Encode(value interface{}) string {
	return fmt.Sprint(value)
}

func (CTFloat32) BinRead(r io.Reader) (v interface{}, err error) {
	var i float32
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTFloat32) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

type CTFloat64 struct {
	UnLimitColumnTool
}

func (CTFloat64) Zero() interface{} {
	return float64(0)
}

func (CTFloat64) Decode(value string) (v interface{}, err error) {
	var f float64
	if f, err = strconv.ParseFloat(value, 64); err != nil {
		return
	}
	return f, nil
}

func (CTFloat64) Encode(value interface{}) string {
	return fmt.Sprint(value)
}

func (CTFloat64) BinRead(r io.Reader) (v interface{}, err error) {
	var i float64
	if err = br(r, &i); err != nil {
		return
	}
	v = i
	return
}

func (CTFloat64) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

const (
	CtFloat32 ColumnType = "f32"
	CtFloat64 ColumnType = "f64"
)

func init() {
	ColumnTypeTool.Set(CtFloat32, CTFloat32{}, "float32", "float")
	ColumnTypeTool.Set(CtFloat64, CTFloat64{}, "float64")
}
