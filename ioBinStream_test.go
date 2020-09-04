package bintb

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestBinStreamReadWrite(t *testing.T) {
	var (
		columns = MustParseColumns("name:s", "year:i32", "ok:b")
		recs    = []Recorde{{"first user", int32(2020), true}, {"second user", int32(2021), false}}
		wantB   = []byte{
			// columns
			0, 3, 6, 110, 97, 109, 101, 58, 115, 8, 121, 101, 97, 114, 58, 105, 51, 50, 4, 111, 107, 58, 98,
			0, 1, // has one
			0, 0, 10, 102, 105, 114, 115, 116, 32, 117, 115, 101, 114, 0, 0, 0, 7, 228, 0, 1,
			0, 1, // has one
			0, 0, 11, 115, 101, 99, 111, 110, 100, 32, 117, 115, 101, 114, 0, 0, 0, 7, 229, 0, 0,
			0, 0, // no has
		}
	)

	var w bytes.Buffer

	if err := BinStreamWriteF(&w, NewEncoder(columns), NextRecords(recs), 20); err != nil {
		t.Errorf("BinStreamWriteF() error = %v", err)
		return
	}

	if gotB := w.Bytes(); string(gotB) != string(wantB) {
		t.Errorf("BinColumnWriter.Write() = %v, want %v", gotB, wantB)
		return
	}

	dec := NewDecoder(nil)

	if err := OpenIterate(NewBinStreamReader(dec, &w, &StreamReaderOptions{
		OnRecorde: func(i int, rec Recorde) (err error) {
			if !reflect.DeepEqual(rec, recs[i-1]) {
				return fmt.Errorf("%v, want %v", rec, recs[i])
			}
			return
		},
	})); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(dec.Columns, columns) {
		t.Errorf("NewBinStreamReader.Header() = %v, want %v", dec.Columns, columns)
	}
}

func TestBinStreamReadWriteRequiredColumns(t *testing.T) {
	var (
		columns = MustParseColumns("name*:s", "year*:i32", "ok*:b")
		recs    = []Recorde{{"first user", int32(2020), true}, {"second user", int32(2021), false}}
		wantB   = []byte{
			// columns
			0, 3, 7, 110, 97, 109, 101, 43, 58, 115, 9, 121, 101, 97, 114, 43, 58, 105, 51, 50, 5, 111, 107, 43, 58, 98,
			0, 1, // has one
			0, 10, 102, 105, 114, 115, 116, 32, 117, 115, 101, 114, 0, 0, 7, 228, 1,
			0, 1, // has one
			0, 11, 115, 101, 99, 111, 110, 100, 32, 117, 115, 101, 114, 0, 0, 7, 229, 0,
			0, 0, // no has
		}
	)

	var w bytes.Buffer

	if err := BinStreamWriteF(&w, NewEncoder(columns), NextRecords(recs), 20); err != nil {
		t.Errorf("BinStreamWriteF() error = %v", err)
		return
	}

	if gotB := w.Bytes(); string(gotB) != string(wantB) {
		t.Errorf("BinColumnWriter.Write() = %v, want %v", gotB, wantB)
		return
	}

	dec := NewDecoder(nil)

	if err := OpenIterate(NewBinStreamReader(dec, &w, &StreamReaderOptions{
		OnRecorde: func(i int, rec Recorde) (err error) {
			if !reflect.DeepEqual(rec, recs[i-1]) {
				return fmt.Errorf("%v, want %v", rec, recs[i-1])
			}
			return
		},
	})); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(dec.Columns, columns) {
		t.Errorf("NewBinStreamReader.Header() = %v, want %v", dec.Columns, columns)
	}
}
