package bintb

import (
	"fmt"
	"io"
)

type RecordReader interface {
	Read() (rec Recorde, err error)
}

type RecordWriter interface {
	Write(rec Recorde) (err error)
}

type ColumnReader interface {
	Read() (col *Column, err error)
}

type ColumnWriter interface {
	Write(col *Column) (err error)
}

type ReadHeader struct {
	Items  []*Column
	ByName map[string]int
	HasID  bool
}

func NewReadHeader() *ReadHeader {
	return &ReadHeader{ByName: map[string]int{}}
}

func (this *ReadHeader) Add(col ...*Column) error {
	for _, col := range col {
		if _, ok := this.ByName[col.Name]; ok {
			return fmt.Errorf("column %q duplicated", col.Name)
		}
		this.ByName[col.Name] = len(this.Items)
		this.Items = append(this.Items, col)
	}
	return nil
}

func (this *ReadHeader) Get(name string, recorde Recorde) interface{} {
	return recorde[this.ByName[name]]
}


type MemCacheReader struct {
	r io.Reader
}

func NewMemCacheReader(r io.Reader) *MemCacheReader {
	return &MemCacheReader{r: r}
}

func (this MemCacheReader) Read(p []byte) (n int, err error) {
	n, err = this.r.Read(p)
	if err == nil && len(p) != n {
		var n2 int
		n2, err = this.Read(p[n:])
		n += n2
	}
	return
}