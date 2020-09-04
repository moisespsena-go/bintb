package bintb

import (
	"bytes"
	"context"
	"reflect"
	"testing"
)

func TestBinRecordReader_Read(t *testing.T) {
	tests := []struct {
		name    string
		columns []*Column
		b       []byte
		wantRec Recorde
		wantErr bool
	}{
		{"a", MustParseColumns("name:s", "year:i32", "day*:i32"), []byte{0, 4, 117, 115, 101, 114, 0, 0, 0, 7, 228, 0, 0, 0, 25}, []interface{}{"user", int32(2020), int32(25)}, false},
		{"a", MustParseColumns("name:s", "year:i32", "ok*:bool"), []byte{0, 4, 117, 115, 101, 114, 0, 0, 0, 7, 228, 1}, []interface{}{"user", int32(2020), true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &BinRecordReader{
				decoder: NewDecoder(tt.columns, context.Background()),
				r:       bytes.NewReader(tt.b),
			}
			gotRec, err := this.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("BinRecordReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRec, tt.wantRec) {
				t.Errorf("BinRecordReader.Read() = %v, want %v", gotRec, tt.wantRec)
			}
		})
	}
}

func TestBinRecordWriter_Write(t *testing.T) {
	tests := []struct {
		name    string
		columns []*Column
		rec     Recorde
		wantB   []byte
		wantErr bool
	}{
		{"a", MustParseColumns("name:s", "year:i32", "day*:i32"), []interface{}{"user", int32(2020), int32(25)}, []byte{0, 4, 117, 115, 101, 114, 0, 0, 0, 7, 228, 0, 0, 0, 25}, false},
		{"a", MustParseColumns("name:s", "year:i32", "ok*:bool"), []interface{}{"user", int32(2020), true}, []byte{0, 4, 117, 115, 101, 114, 0, 0, 0, 7, 228, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w bytes.Buffer
			this := &BinRecordWriter{
				encoder: NewEncoder(tt.columns),
				w:       &w,
			}
			if err := this.Write(tt.rec); (err != nil) != tt.wantErr {
				t.Errorf("BinRecordWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotB := w.Bytes(); string(gotB) != string(tt.wantB) {
				t.Errorf("BinRecordWriter.Write() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestBinColumnReader_Read(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		wantCol *Column
		wantErr bool
	}{
		{"name", []byte{6, 110, 97, 109, 101, 58, 115}, MustParseColumn("name:s"), false},
		{"date", []byte{10, 116, 104, 101, 95, 100, 97, 116, 101, 58, 100}, MustParseColumn("the_date:d"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &BinColumnReader{
				r: bytes.NewReader(tt.b),
			}
			gotCol, err := this.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("BinColumnReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCol, tt.wantCol) {
				t.Errorf("BinColumnReader.Read() = %v, want %v", gotCol, tt.wantCol)
			}
		})
	}
}

func TestBinColumnWriter_Write(t *testing.T) {
	tests := []struct {
		name    string
		col     *Column
		wantB   []byte
		wantErr bool
	}{
		{"name", MustParseColumn("name:s"), []byte{6, 110, 97, 109, 101, 58, 115}, false},
		{"date", MustParseColumn("the_date:d"), []byte{10, 116, 104, 101, 95, 100, 97, 116, 101, 58, 100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w bytes.Buffer
			this := &BinColumnWriter{
				w: &w,
			}
			if err := this.Write(tt.col); (err != nil) != tt.wantErr {
				t.Errorf("BinColumnWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotB := w.Bytes(); string(gotB) != string(tt.wantB) {
				t.Errorf("BinColumnWriter.Write() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
