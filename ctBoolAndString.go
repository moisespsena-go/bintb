package bintb

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

type CTbool struct{ UnLimitColumnTool }

func (CTbool) Zero() interface{} {
	return false
}

func (CTbool) Description() string {
	return "Boolean. (TRUE values: y, yes, t, true, 1; FALSE values: n, no, f, false, 0)"
}

func (CTbool) Decode(value string) (v interface{}, err error) {
	switch strings.ToLower(value) {
	case "y", "yes", "t", "true", "1":
		return true, nil
	case "n", "no", "f", "false", "0":
		return false, nil
	default:
		err = errors.New("bad boolean value")
		return
	}
}

func (CTbool) Encode(value interface{}) string {
	if value.(bool) {
		return "true"
	}
	return "false"
}

func (CTbool) BinRead(r io.Reader) (v interface{}, err error) {
	var i bool
	if err = br(r, &i); err != nil {
		return
	}
	return i, nil
}

func (CTbool) BinWrite(w io.Writer, v interface{}) (err error) {
	return bw(w, v)
}

type CTstring struct{ LimitColumnTool }

func (CTstring) Description() string {
	return "The text value"
}

func (CTstring) Zero() interface{} {
	return ""
}

func (CTstring) Decode(value string) (v interface{}, err error) {
	return value, nil
}

func (CTstring) Encode(value interface{}) string {
	switch vt := value.(type) {
	case string:
		return vt
	case fmt.Stringer:
		return vt.String()
	case []byte:
		return base64.URLEncoding.EncodeToString(vt)
	default:
		return fmt.Sprint(vt)
	}
}

func (CTstring) BinRead(r io.Reader) (v interface{}, err error) {
	var b [2]byte
	if _, err = r.Read(b[:]); err != nil {
		return
	}
	l := binOrder.Uint16(b[:])
	data := make([]byte, l)
	if _, err = r.Read(data); err != nil {
		return
	}
	return string(data), nil
}

func (CTstring) BinWrite(w io.Writer, v interface{}) (err error) {
	var b [2]byte
	binOrder.PutUint16(b[:], uint16(len(v.(string))))
	_, err = w.Write(append(b[:], []byte(v.(string))...))
	return
}

type CTbyte struct{ UnLimitColumnTool }

func (CTbyte) Zero() interface{} {
	return '0'
}
func (CTbyte) TypeName() string {
	return "byte"
}

func (CTbyte) Description() string {
	return "The binary char value"
}

func (CTbyte) Decode(value string) (v interface{}, err error) {
	var b []byte
	if b, err = hex.DecodeString(value); err != nil {
		return
	}
	return b[0], nil
}

func (CTbyte) Encode(value interface{}) string {
	return hex.EncodeToString([]byte{value.(byte)})
}

func (CTbyte) BinRead(r io.Reader) (v interface{}, err error) {
	var b [1]byte
	if _, err = r.Read(b[:]); err != nil {
		return
	}
	return b[0], nil
}

func (CTbyte) BinWrite(w io.Writer, v interface{}) (err error) {
	_, err = w.Write([]byte{v.(byte)})
	return
}

type CTbytes struct{ LimitColumnTool }

func (CTbytes) TypeName() string {
	return "[]byte"
}

func (CTbytes) Description() string {
	return "The binary sequence value"
}

func (CTbytes) Zero() interface{} {
	return []byte{}
}

func (CTbytes) Decode(value string) (v interface{}, err error) {
	var b []byte
	b, err = hex.DecodeString(value)
	return b, err
}

func (CTbytes) Encode(value interface{}) string {
	return hex.EncodeToString(value.([]byte))
}

func (CTbytes) BinRead(r io.Reader) (v interface{}, err error) {
	var b [2]byte
	if _, err = r.Read(b[:]); err != nil {
		return
	}
	l := binOrder.Uint16(b[:])
	data := make([]byte, l)
	if _, err = r.Read(data); err != nil {
		return
	}
	return data, nil
}

func (CTbytes) BinWrite(w io.Writer, v interface{}) (err error) {
	var b [2]byte
	binOrder.PutUint16(b[:], uint16(len(v.([]byte))))
	_, err = w.Write(append(b[:], v.([]byte)...))
	return
}

const (
	CtBool   ColumnType = "b"
	CtString ColumnType = "s"
	CtByte   ColumnType = "B"
	CtBytes  ColumnType = "Bs"
)

func init() {
	ColumnTypeTool.Set(CtBool, CTbool{}, "bool")
	ColumnTypeTool.Set(CtString, CTstring{}, "string", "str")
	ColumnTypeTool.Set(CtByte, CTbyte{}, "byte")
	ColumnTypeTool.Set(CtBytes, CTbytes{}, "bytes")
}
