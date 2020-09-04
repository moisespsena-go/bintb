package bintb

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type StreamReaderOptions struct {
	IdColumn  *Column
	OnHeader  func(header *ReadHeader) (err error)
	OnRecorde func(i int, r Recorde) error
}

type BinStreamReader struct {
	decoder      *Decoder
	r            io.Reader
	opt          *StreamReaderOptions
	br           *BinRecordReader
	recordeCount int
}

func NewBinStreamReader(dec *Decoder, r io.Reader, opts ...*StreamReaderOptions) *BinStreamReader {
	r = MemCacheReader{r}
	opt := &StreamReaderOptions{}
	for _, opt_ := range opts {
		if opt_ != nil {
			opt = opt_
		}
	}
	if dec == nil {
		dec = NewDecoder(nil)
	}
	return &BinStreamReader{
		decoder: dec,
		r:       r,
		opt:     opt,
		br:      &BinRecordReader{decoder: dec, r: r},
	}
}

func (this *BinStreamReader) Open() (header *ReadHeader, it RecordsIterator, err error) {
	var countb [2]byte
	if _, err = this.r.Read(countb[:]); err != nil {
		return
	}
	count := int(binOrder.Uint16(countb[:]))
	cr := BinColumnReader{this.r}
	header = NewReadHeader()

	for i := 0; i < count; i++ {
		var col *Column
		if col, err = cr.Read(); err != nil {
			return
		}
		if err = header.Add(col); err != nil {
			return nil, nil, err
		}
	}

	if len(this.decoder.Columns) > 0 {
		if this.opt.IdColumn != nil {
			for name := range header.ByName {
				if name == this.opt.IdColumn.Name {
					this.decoder.Columns = append([]*Column{this.opt.IdColumn}, this.decoder.Columns...)
					header.HasID = true
					break
				}
			}
		}
		if len(header.Items) != len(this.decoder.Columns) {
			return nil, nil, errors.New("bad columns count")
		}
		byName := map[string]int{}
		for i, col := range this.decoder.Columns {
			byName[col.Name] = i
		}

		for x, col := range header.Items {
			if i, ok := byName[col.Name]; !ok {
				return nil, nil, fmt.Errorf("unexpected column %q", col.Name)
			} else if dcol := this.decoder.Columns[i]; col.String() != dcol.String() {
				return nil, nil, fmt.Errorf("unexpected column spec %q", col.String())
			} else {
				header.Items[x] = dcol
			}
		}
	}
	this.decoder.Columns = header.Items
	if this.opt.OnHeader != nil {
		err = this.opt.OnHeader(header)
	}
	it = &BinRecordsStreamIterator{this}
	return
}

type BinRecordsStreamIterator struct {
	r *BinStreamReader
}

func (this *BinRecordsStreamIterator) Start() (state interface{}, err error) {
	s := &struct {
		r     *BinStreamReader
		count int
		left  uint16
	}{r: this.r}
	this.r = nil
	if s.left, err = this.left(s.r); err != nil {
		return
	}
	return s, nil
}

func (BinRecordsStreamIterator) left(r *BinStreamReader) (left uint16, err error) {
	var (
		bleft [2]byte
		n int
	)

	if n, err = r.r.Read(bleft[:]); !(err == io.EOF && n == 2) && err != nil {
		return 0, errors.Wrap(err, "read chunk size")
	}
	return binOrder.Uint16(bleft[:]), nil
}

func (BinRecordsStreamIterator) Done(state interface{}) (ok bool) {
	return state.(*struct {
		r     *BinStreamReader
		count int
		left  uint16
	}).left == 0
}

func (this *BinRecordsStreamIterator) Next(state interface{}) (rec Recorde, newState interface{}, err error) {
	s := state.(*struct {
		r     *BinStreamReader
		count int
		left  uint16
	})
	if rec, err = s.r.br.Read(); err != nil {
		return
	}
	s.left--
	s.count++
	if s.r.opt.OnRecorde != nil {
		if err = s.r.opt.OnRecorde(s.count, rec); err != nil {
			return
		}
	}
	if s.left == 0 {
		if s.left, err = this.left(s.r); err != nil {
			return
		}
	}
	return rec, s, nil
}
