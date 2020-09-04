package bintb

import (
	"bytes"
	"io"
	"math"
	"reflect"
	"testing"
)

func TestCTFloat32_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"e", "3.4028235e+38", float32(3.4028235e+38), false},
		{"E", "3.4028235E+38", float32(3.4028235e+38), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTFloat32{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTFloat32.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTFloat32.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTFloat32_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", float32(3.4028235e+38), "3.4028235e+38"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTFloat32{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTFloat32.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTFloat32_BinRead(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{127,127,255,255}, float32(3.4028235e+38), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTFloat32{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTFloat32.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTFloat32.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTFloat32_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", float32(3.4028235e+38), []byte{127,127,255,255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTFloat32{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTFloat32.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTFloat32.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}


func TestCTFloat64_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"e", "1.7976931348623157e+308", float64(1.7976931348623157e+308), false},
		{"E", "1.7976931348623157E+308", float64(1.7976931348623157e+308), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTFloat64{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTFloat64.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTFloat64.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTFloat64_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", float64(math.MaxFloat64), "1.7976931348623157e+308"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTFloat64{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTFloat64.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTFloat64_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{127,239,255,255,255,255,255,255}, float64(math.MaxFloat64), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTFloat64{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTFloat64.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTFloat64.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTFloat64_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", float64(math.MaxFloat64), []byte{127,239,255,255,255,255,255,255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTFloat64{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTFloat64.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTFloat64.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
