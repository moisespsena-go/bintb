package bintb

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/moisespsena-go/aorm/types"
)

func TestCTtime_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "12:35:27", types.NewTime(12, 35, 27), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTtime{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTtime.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTtime.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTtime_Encode(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"a", types.NewTime(12, 35, 27), "12:35:27"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTtime{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTtime.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTtime_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		r       []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{12, 35, 27}, types.NewTime(12, 35, 27), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTtime{}
			gotV, err := c.BinRead(bytes.NewReader(tt.r))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTtime.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTtime.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTtime_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v       interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", types.NewTime(12, 35, 27), []byte{12, 35, 27}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTtime{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTtime.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != string(tt.wantW) {
				t.Errorf("CTtime.BinWrite() = %v, want %v", []byte(gotW), tt.wantW)
			}
		})
	}
}

func TestCTtimeZ_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "12:35:27 -13:25", types.NewTimeZ(12, 35, 27, -13, 25), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTtimeZ{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTtime.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTtime.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTtimeZ_Encode(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"a", types.NewTimeZ(12, 35, 27, -13, 25), "12:35:27 -13:25"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTtimeZ{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTtime.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTtimeZ_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		r       []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{12, 35, 27, 243, 25}, types.NewTimeZ(12, 35, 27, -13, 25), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTtimeZ{}
			gotV, err := c.BinRead(bytes.NewReader(tt.r))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTtime.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTtime.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTtimeZ_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v       interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", types.NewTimeZ(12, 35, 27, -13, 25), []byte{12, 35, 27, 243, 25}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTtimeZ{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTtime.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != string(tt.wantW) {
				t.Errorf("CTtime.BinWrite() = %v, want %v", []byte(gotW), tt.wantW)
			}
		})
	}
}
