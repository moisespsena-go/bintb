package bintb

import (
	"cloud.google.com/go/civil"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"
)

var (
	reDate = regexp.MustCompile(`^(\d+)-(\d+)-(\d+)$`)
)

type ToolDater interface {
	DateOf(value interface{}) civil.Date
}

type DateTool struct {}

func (DateTool) DateOf(value interface{}) civil.Date {
	return civil.DateOf(value.(time.Time))
}

func IsDateTool(typ ColumnType) (ok bool) {
	if tool := ColumnTypeTool.Get(typ); tool != nil {
		_, ok = tool.(ToolDater)
	}
	return
}

type CTdate struct {
	UnLimitColumnTool
	TimeZero
	DateTool
}

func (CTdate) Decode(value string) (v interface{}, err error) {
	result := reDate.FindAllStringSubmatch(value, 1)
	if len(result) == 0 {
		return nil, fmt.Errorf("bad date format `%s`", value)
	}
	r0 := result[0][1:]
	Y, _ := strconv.Atoi(r0[0])
	M, _ := strconv.Atoi(r0[1])
	D, _ := strconv.Atoi(r0[2])
	return time.Date(Y, time.Month(M), D, 0, 0, 0, 0, time.UTC), nil
}

func (CTdate) Encode(value interface{}) string {
	return value.(time.Time).Format("2006-01-02")
}

func (CTdate) BinRead(r io.Reader) (v interface{}, err error) {
	var b [8]byte
	if _, err = r.Read(b[:]); err != nil {
		return
	}
	return time.Unix(int64(binOrder.Uint64(b[:])), 0).In(time.UTC), nil
}

func (CTdate) BinWrite(w io.Writer, v interface{}) (err error) {
	t := v.(time.Time)
	return bw(w, t.Unix())
}

const (
	CtDate ColumnType = "d"
)

func init() {
	ColumnTypeTool.Set(CtDate, CTdate{}, "date")
}
