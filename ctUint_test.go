package bintb

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

func TestCTuint8_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "255", uint8(math.MaxUint8), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint8{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTuint8.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTuint8.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTuint8_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", uint8(math.MaxUint8), "255"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint8{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTuint8.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTuint8_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{255}, uint8(math.MaxUint8), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint8{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTuint8.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTuint8.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTuint8_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", uint8(math.MaxUint8), []byte{255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint8{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTuint8.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTuint8.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestCTuint16_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "65535", uint16(math.MaxUint16), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint16{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTuint16.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTuint16.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTuint16_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", uint16(math.MaxUint16), "65535"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint16{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTuint16.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTuint16_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{255,255}, uint16(math.MaxUint16), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint16{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTuint16.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTuint16.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTuint16_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", uint16(math.MaxUint16), []byte{255,255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint16{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTuint16.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTuint16.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}


func TestCTuint32_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "4294967295", uint32(math.MaxUint32), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint32{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTuint32.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTuint32.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTuint32_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", uint32(math.MaxUint32), "4294967295"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint32{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTuint32.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTuint32_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{255,255,255,255}, uint32(math.MaxUint32), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint32{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTuint32.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTuint32.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTuint32_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", uint32(math.MaxUint32), []byte{255,255,255,255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint32{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTuint32.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTuint32.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}


func TestCTuint64_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "18446744073709551615", uint64(math.MaxUint64), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint64{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTuint64.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTuint64.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTuint64_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", uint64(math.MaxUint64), "18446744073709551615"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint64{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTuint64.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTuint64_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{255,255,255,255,255,255,255,255}, uint64(math.MaxUint64), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint64{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTuint64.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTuint64.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTuint64_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", uint64(math.MaxUint64), []byte{255,255,255,255,255,255,255,255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTuint64{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTuint64.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTuint64.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}