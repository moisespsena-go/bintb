package bintb

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

func TestCTint8_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "127", int8(math.MaxInt8), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint8{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTint8.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTint8.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTint8_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", int8(math.MaxInt8), "127"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint8{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTint8.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTint8_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{127}, int8(math.MaxInt8), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint8{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTint8.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTint8.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTint8_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", int8(math.MaxInt8), []byte{127}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint8{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTint8.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTint8.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestCTint16_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "32767", int16(math.MaxInt16), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint16{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTint16.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTint16.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTint16_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", int16(math.MaxInt16), "32767"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint16{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTint16.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTint16_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{127,255}, int16(math.MaxInt16), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint16{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTint16.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTint16.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTint16_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", int16(math.MaxInt16), []byte{127,255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint16{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTint16.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTint16.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}


func TestCTint32_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "2147483647", int32(math.MaxInt32), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint32{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTint32.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTint32.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTint32_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", int32(math.MaxInt32), "2147483647"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint32{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTint32.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTint32_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{127,255,255,255}, int32(math.MaxInt32), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint32{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTint32.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTint32.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTint32_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", int32(math.MaxInt32), []byte{127,255,255,255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint32{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTint32.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTint32.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}


func TestCTint64_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "9223372036854775807", int64(math.MaxInt64), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint64{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTint64.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTint64.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTint64_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", int64(math.MaxInt64), "9223372036854775807"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint64{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTint64.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTint64_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{127,255,255,255,255,255,255,255}, int64(math.MaxInt64), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint64{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTint64.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTint64.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTint64_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", int64(math.MaxInt64), []byte{127,255,255,255,255,255,255,255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTint64{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTint64.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTint64.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}