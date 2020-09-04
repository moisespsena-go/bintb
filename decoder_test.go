package bintb

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeDecode_Decode(t *testing.T) {
	columns := MustParseColumns("name*:s", "year:i32", "day*:i32")
	tests := []struct {
		name    string
		values  []string
		wantRec []interface{}
		wantErr bool
	}{
		{"a", []string{"user", "2020", "25"}, []interface{}{"user", int32(2020), int32(25)}, false},
		{"b", []string{"user", "", "25"}, []interface{}{"user", nil, int32(25)}, false},
		{"c", []string{"user", "", ""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Decoder{
				Columns: columns,
			}
			gotRec, err := this.Decode(tt.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRec, tt.wantRec) {
				t.Errorf("Decoder.Decode() = %v, want %v", gotRec, tt.wantRec)
			}
		})
	}
}

func TestEncodeDecode_Encode(t *testing.T) {
	columns := MustParseColumns("name*:s", "year:i32", "day*:i32")
	tests := []struct {
		name    string
		wantRec []string
		values  []interface{}
		wantErr bool
	}{
		{"c", nil, []interface{}{nil, nil, nil}, true},
		{"a", []string{"user", "2020", "25"}, []interface{}{"user", int32(2020), int32(25)}, false},
		{"b", []string{"user", "", "25"}, []interface{}{"user", nil, int32(25)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewEncoder(columns)
			gotRec, err := this.Encode(tt.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRec, tt.wantRec) {
				t.Errorf("Decoder.Encode() = %v, want %v", gotRec, tt.wantRec)
			}
		})
	}
}

func TestEncodeDecode_BinRead(t *testing.T) {
	columns := MustParseColumns("name*:s", "year:i32", "day*:i32")
	tests := []struct {
		name    string
		b       []byte
		wantRec []interface{}
		wantErr bool
	}{
		{"a", []byte{0, 4, 117, 115, 101, 114, 0, 0, 0, 7, 228, 0, 0, 0, 25}, []interface{}{"user", int32(2020), int32(25)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewDecoder(columns, nil)
			gotRec, err := this.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRec, tt.wantRec) {
				t.Errorf("Decoder.BinRead() = %v, want %v", gotRec, tt.wantRec)
			}
		})
	}
}

func TestEncodeDecode_BinWrite(t *testing.T) {
	columns := MustParseColumns("name:s", "year:i32", "day*:i32")
	tests := []struct {
		name    string
		values  []interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", []interface{}{"user", int32(2020), int32(25)}, []byte{0, 4, 117, 115, 101, 114, 0, 0, 0, 7, 228, 0, 0, 0, 25}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Encoder{
				Columns: columns,
			}
			w := &bytes.Buffer{}
			if err := this.BinWrite(w, tt.values...); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("Decoder.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
