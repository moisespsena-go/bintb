package bintb

import (
	"bytes"
	"io"

	"github.com/moisespsena-go/aorm/types"
)

type TimeZero struct{ UnLimitColumnTool }

type CTtime struct{ UnLimitColumnTool }

func (CTtime) Zero() interface{} {
	var t types.Time
	return t
}

func (CTtime) Description() string {
	return "Time. (15:04:05)"
}

func (CTtime) Decode(value string) (v interface{}, err error) {
	var t types.Time
	if err = t.Scan(value); err != nil {
		return
	}
	return t, nil
}

func (CTtime) Encode(value interface{}) string {
	return value.(types.Time).String()
}

func (CTtime) BinRead(r io.Reader) (v interface{}, err error) {
	var h, m, s int8
	if err = br(r, &h, &m, &s); err != nil {
		return
	}
	return types.NewTime(int(h), int(m), int(s)), nil
}

func (CTtime) BinWrite(w io.Writer, v interface{}) (err error) {
	t := v.(types.Time)
	var buf bytes.Buffer
	bw(&buf, int8(t.Hour()), int8(t.Minute()), int8(t.Second()))
	_, err = w.Write(buf.Bytes())
	return
}

type ToolZonedTimer interface {
	Zone(value interface{}) (h, m int)
}

func IsZonedTimeTool(typ ColumnType) (ok bool) {
	if tool := ColumnTypeTool.Get(typ); tool != nil {
		_, ok = tool.(ToolZonedTimer)
	}
	return
}


type CTtimeZ struct{ UnLimitColumnTool }

func (CTtimeZ) Zone(value interface{}) (h, m int) {
	return value.(types.TimeZ).Zone()
}

func (CTtimeZ) Zero() interface{} {
	var t types.TimeZ
	return t
}

func (CTtimeZ) Description() string {
	return "Time with Zone: (15:04:05 -07:00)"
}

func (CTtimeZ) Decode(value string) (v interface{}, err error) {
	var t types.TimeZ
	if err = t.Scan(value); err != nil {
		return
	}
	return t, nil
}

func (CTtimeZ) Encode(value interface{}) string {
	return value.(types.TimeZ).String()
}

func (CTtimeZ) BinRead(r io.Reader) (v interface{}, err error) {
	var h, m, s, zh, zm int8
	if err = br(r, &h, &m, &s, &zh, &zm); err != nil {
		return
	}
	return types.NewTimeZ(int(h), int(m), int(s), int(zh), int(zm)), nil
}

func (CTtimeZ) BinWrite(w io.Writer, v interface{}) (err error) {
	t := v.(types.TimeZ)
	var buf bytes.Buffer
	zh, zm := t.Zone()
	bw(&buf, int8(t.Hour()), int8(t.Minute()), int8(t.Second()), int8(zh), int8(zm))
	_, err = w.Write(buf.Bytes())
	return
}

const (
	CtTime  ColumnType = "t"
	CtTimeZ ColumnType = "tz"
)

func init() {
	ColumnTypeTool.Set(CtTime, CTtime{}, "time")
	ColumnTypeTool.Set(CtTimeZ, CTtimeZ{}, "timez")
}
