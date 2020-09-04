package bintb

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCTbool_Decode(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"true01", "y", true, false},
		{"true02", "t", true, false},
		{"true03", "T", true, false},
		{"true04", "yes", true, false},
		{"true05", "true", true, false},
		{"true06", "1", true, false},

		{"false01", "f", false, false},
		{"false02", "n", false, false},
		{"false03", "no", false, false},
		{"false04", "N", false, false},
		{"false05", "false", false, false},
		{"false06", "0", false, false},

		{"bad", "3", true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbool{}
			gotV, err := c.Decode(tt.value)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("CTbool.Decode() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTbool.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTbool_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"true", true, "true"},
		{"false", false, "false"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbool{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTbool.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTbool_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"true", []byte{1}, true, false},
		{"false", []byte{0}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbool{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTbool.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTbool.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTbool_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"true", true, []byte{1}, false},
		{"false", false, []byte{0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbool{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTbool.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTbool.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestCTstring_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "a", "a", false},
		{"ab", "ab", "ab", false},
		{"abc", "abc", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTstring{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTstring.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTstring.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTstring_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", "a", "a"},
		{"ab", "ab", "ab"},
		{"abc", "abc", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTstring{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTstring.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTstring_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{

		{"a", append([]byte{0, 1}, []byte("a")...), "a", false},
		{"ab", append([]byte{0, 2}, []byte("ab")...), "ab", false},
		{"abc", append([]byte{0, 3}, []byte("abc")...), "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTstring{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTstring.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTstring.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTstring_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", "a", append([]byte{0,  1}, []byte("a")...), false},
		{"ab","ab", append([]byte{0, 2}, []byte("ab")...),  false},
		{"abc", "abc", append([]byte{0, 3}, []byte("abc")...),  false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTstring{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTstring.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTstring.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestCTbyte_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"41", "41", byte('A'), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbyte{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTbyte.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTbyte.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTbyte_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"41", byte('A'), "41"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbyte{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTbyte.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTbyte_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{
		{"41", []byte{'A'}, byte('A'), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbyte{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTbyte.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTbyte.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTbyte_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"41", byte('A'), []byte{'A'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbyte{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTbyte.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTbyte.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}


func TestCTbytes_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "61", []byte("a"), false},
		{"ab", "6162", []byte("ab"), false},
		{"abc", "616263", []byte("abc"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbytes{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTbytes.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTbytes.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTbytes_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", []byte("a"), "61"},
		{"ab", []byte("ab"), "6162"},
		{"abc", []byte("abc"), "616263"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbytes{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTbytes.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTbytes_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		b []byte
		wantV   interface{}
		wantErr bool
	}{

		{"a", append([]byte{0, 1}, []byte("a")...), []byte("a"), false},
		{"ab", append([]byte{0, 2}, []byte("ab")...), []byte("ab"), false},
		{"abc", append([]byte{0, 3}, []byte("abc")...), []byte("abc"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbytes{}
			gotV, err := c.BinRead(bytes.NewReader(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTbytes.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTbytes.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTbytes_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v string
		wantW   []byte
		wantErr bool
	}{
		{"a", "a", append([]byte{0, 1}, []byte("a")...), false},
		{"ab","ab", append([]byte{0, 2}, []byte("ab")...),  false},
		{"abc", "abc", append([]byte{0, 3}, []byte("abc")...),  false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTbytes{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, []byte(tt.v)); (err != nil) != tt.wantErr {
				t.Errorf("CTbytes.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.Bytes(); string(gotW) != string(tt.wantW) {
				t.Errorf("CTbytes.BinWrite() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
