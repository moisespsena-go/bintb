package bintb

import (
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var reColSep = regexp.MustCompile("^#? *col-sep *: *(.)$")

type CsvStreamReaderOptions struct {
	Decoder         *Decoder
	OnlyColumnNames bool
	StreamReaderOptions
	OnStrings       func(i int, values []string) []string
	OnHeaderStrings func(values []string) []string
}

type CsvStreamReader struct {
	r            *csv.Reader
	dec          *Decoder
	br           CsvRecordReader
	last         []string
	recordeCount int
	opt          *CsvStreamReaderOptions
}

func NewCsvStreamReader(r *csv.Reader, opts ...*CsvStreamReaderOptions) *CsvStreamReader {
	opt := &CsvStreamReaderOptions{}
	for _, opt_ := range opts {
		if opt_ != nil {
			opt = opt_
		}
	}
	dec := opt.Decoder
	if dec == nil {
		dec = NewDecoder(nil)
	}
	r.TrimLeadingSpace = true

	if opt.OnStrings == nil {
		cp := *opt
		cp.OnStrings = func(i int, s []string) []string {
			return s
		}
		opt = &cp
	}
	if opt.OnHeaderStrings == nil {
		cp := *opt
		cp.OnHeaderStrings = func(s []string) []string {
			return s
		}
		opt = &cp
	}

	return &CsvStreamReader{
		dec: dec,
		r:   r,
		opt: opt,
	}
}

func (this *CsvStreamReader) RecordeCount() int {
	return this.recordeCount
}

func (this *CsvStreamReader) Open() (header *ReadHeader, it RecordsIterator, err error) {
	var columnsS []string
	if columnsS, err = this.r.Read(); err != nil {
		err = errors.Wrap(err, "read columns")
		return
	}

	columnsS = this.opt.OnHeaderStrings(columnsS)

	// detec columns separator
	if len(columnsS) == 1 {
		if matches := reColSep.FindAllStringSubmatch(columnsS[0], -1); len(matches) == 1 {
			this.r.Comma = rune(matches[0][1][0])
			this.r.FieldsPerRecord = 0
			if columnsS, err = this.r.Read(); err != nil {
				err = errors.Wrap(err, "read columns")
				return
			}
		}
	}

	if len(columnsS) == 0 {
		err = errors.New("empty input")
		return
	}
	var (
		onlyNames, hasId bool
	)
	if len(this.dec.Columns) > 0 {
		if this.opt.IdColumn != nil {
			for _, name := range columnsS {
				if name == this.opt.IdColumn.Name {
					this.dec.Columns = append([]*Column{this.opt.IdColumn}, this.dec.Columns...)
					hasId = true
					break
				}
			}
		}
		if len(columnsS) != len(this.dec.Columns) {
			err = errors.New("bad columns count")
			return
		}
		onlyNames = !strings.ContainsRune(columnsS[0], ':')
	} else if !strings.ContainsRune(columnsS[0], ':') {
		err = errors.New("require full columns spec")
		return
	}

	header = &ReadHeader{
		ByName: map[string]int{},
		HasID:  hasId,
	}

	if onlyNames {
	a:
		for _, name := range columnsS {
			for _, col := range this.dec.Columns {
				if name == col.Name {
					header.Items = append(header.Items, col)
					continue a
				}
			}

			return nil, nil, fmt.Errorf("column %q is not expected", name)
		}
	} else if len(this.dec.Columns) == 0 {
		for _, spec := range columnsS {
			if col, err := ParseColumn(spec); err != nil {
				return nil, nil, errors.Wrapf(err, "parse column %q failed", spec)
			} else {
				header.Items = append(header.Items, col)
			}
		}
	} else {
	b:
		for _, spec := range columnsS {
			if rcol, err := ParseColumn(spec); err != nil {
				return nil, nil, errors.Wrapf(err, "parse column %q failed", spec)
			} else {
				for _, col := range this.dec.Columns {
					if rcol.String() == col.String() {
						header.Items = append(header.Items, col)
						continue b
					}
				}
				return nil, nil, fmt.Errorf("column %q is not expected", spec)
			}
		}
	}

	for i, col := range header.Items {
		if _, ok := header.ByName[col.Name]; ok {
			return nil, nil, fmt.Errorf("duplicate column %q", col.Name)
		}
		header.ByName[col.Name] = i
	}

	if this.opt.OnHeader != nil {
		if err = this.opt.OnHeader(header); err != nil {
			return nil, nil, errors.Wrap(err, "onHeader callback")
		}
	}
	if len(this.dec.Columns) == 0 {
		this.dec.Columns = header.Items
	}
	it = &CsvStreamReaderIterator{this}
	return
}

type CsvStreamReaderIterator struct {
	r *CsvStreamReader
}

func (this *CsvStreamReaderIterator) Start() (state interface{}, err error) {
	s := &struct {
		r     *CsvStreamReader
		last  []string
		count int
	}{r: this.r}
	if s.last, err = this.r.r.Read(); err != nil {
		s.last = nil
		if err == io.EOF {
			err = nil
		} else {
			return
		}
	}
	this.r = nil
	return s, nil
}

func (CsvStreamReaderIterator) Done(state interface{}) (ok bool) {
	return state.(*struct {
		r     *CsvStreamReader
		last  []string
		count int
	}).last == nil
}

func (CsvStreamReaderIterator) Next(state interface{}) (rec Recorde, newState interface{}, err error) {
	s := state.(*struct {
		r     *CsvStreamReader
		last  []string
		count int
	})
	s.count++
	values := s.r.opt.OnStrings(s.count, s.last)
	if rec, err = s.r.dec.Decode(values...); err != nil {
		return
	}
	if s.r.opt.OnRecorde != nil {
		err = errors.Wrap(s.r.opt.OnRecorde(s.count, rec), "onRecorde callback")
	}
	if s.last, err = s.r.r.Read(); err != nil {
		s.last = nil
		if err != io.EOF {
			return
		}
	}
	return rec, s, nil
}
