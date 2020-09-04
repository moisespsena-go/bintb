package bintb

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/moisespsena-go/aorm/types"
)

var (
	reDTime  = regexp.MustCompile(`^(\d+)-(\d+)-(\d+) (\d+):(\d+):(\d+)$`)
	reDTimeZ = regexp.MustCompile(`^(\d+)-(\d+)-(\d+) (\d+):(\d+):(\d+) ([+\-]\d+):(\d+)$`)
)

func (TimeZero) Zero() interface{} {
	var t time.Time
	return t
}

type ToolDateTimer interface {
	DateTimerOf(value interface{}) time.Time
}

type DateTime struct{}

func (DateTool) DateTimerOf(value interface{}) time.Time {
	return value.(time.Time)
}

func IsDateTimeTool(typ ColumnType) (ok bool) {
	if tool := ColumnTypeTool.Get(typ); tool != nil {
		_, ok = tool.(ToolDateTimer)
	}
	return
}

type CTdtime struct {
	TimeZero
	DateTool
	DateTime
}

func (CTdtime) Description() string {
	return "Date Time. (2006-12-31 15:04:05)"
}

func (CTdtime) Decode(value string) (v interface{}, err error) {
	result := reDTime.FindAllStringSubmatch(value, 1)
	if len(result) == 0 {
		return nil, fmt.Errorf("bad date time format `%s`", value)
	}
	r0 := result[0][1:]
	Y, _ := strconv.Atoi(r0[0])
	M, _ := strconv.Atoi(r0[1])
	D, _ := strconv.Atoi(r0[2])
	h, _ := strconv.Atoi(r0[3])
	m, _ := strconv.Atoi(r0[4])
	s, _ := strconv.Atoi(r0[5])
	return time.Date(Y, time.Month(M), D, h, m, s, 0, time.UTC), nil
}

func (CTdtime) Encode(value interface{}) string {
	return value.(time.Time).Format("2006-01-02 15:04:05")
}

func (CTdtime) BinRead(r io.Reader) (v interface{}, err error) {
	var b [8]byte
	if _, err = r.Read(b[:]); err != nil {
		return
	}
	return time.Unix(int64(binOrder.Uint64(b[:])), 0).In(time.UTC), nil
}

func (CTdtime) BinWrite(w io.Writer, v interface{}) (err error) {
	t := v.(time.Time)
	return bw(w, t.Unix())
}

type CTdtimeZ struct {
	TimeZero
	DateTool
	DateTime
}

func (CTdtimeZ) Zone(value interface{}) (h, m int) {
	return types.TZOfTime(value.(time.Time))
}

func (CTdtimeZ) Description() string {
	return "Date Time with Zone. (2006-12-31 15:04:05 -07:00)"
}

func (CTdtimeZ) Decode(value string) (v interface{}, err error) {
	result := reDTimeZ.FindAllStringSubmatch(value, 1)
	if len(result) == 0 {
		return nil, fmt.Errorf("bad date timez format `%s`", value)
	}
	r0 := result[0][1:]
	Y, _ := strconv.Atoi(r0[0])
	M, _ := strconv.Atoi(r0[1])
	D, _ := strconv.Atoi(r0[2])
	h, _ := strconv.Atoi(r0[3])
	m, _ := strconv.Atoi(r0[4])
	s, _ := strconv.Atoi(r0[5])
	zh, _ := strconv.Atoi(r0[6])
	zm, _ := strconv.Atoi(r0[7])
	return time.Parse("2006-01-02 15:04:05 -07:00",
		fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d %03d:%02d", Y, M, D, h, m, s, zh, zm))
}

func (CTdtimeZ) Encode(value interface{}) string {
	return value.(time.Time).Format("2006-01-02 15:04:05 -07:00")
}

func (CTdtimeZ) BinRead(r io.Reader) (v interface{}, err error) {
	var b [8]byte
	if _, err = r.Read(b[:]); err != nil {
		return
	}
	var zh, zm int8
	if err = br(r, &zh, &zm); err != nil {
		return
	}
	loc, _ := time.Parse("-0700", fmt.Sprintf("%03d%02d", zh, zm))
	return time.Unix(int64(binOrder.Uint64(b[:])), 0).In(loc.Location()), nil
}

func (CTdtimeZ) BinWrite(w io.Writer, v interface{}) (err error) {
	t := v.(time.Time)
	loc := t.Format("-0700")
	zh, _ := strconv.Atoi(loc[0:3])
	zm, _ := strconv.Atoi(loc[3:])
	return bw(w, t.Unix(), int8(zh), int8(zm))
}

const (
	CtDTime  ColumnType = "T"
	CtDTimeZ ColumnType = "Tz"
)

func init() {
	ColumnTypeTool.Set(CtDTime, CTdtime{}, "dtime", "timestamp")
	ColumnTypeTool.Set(CtDTimeZ, CTdtimeZ{}, "dtimez", "timestampz")
}
