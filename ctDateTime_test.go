package bintb

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCTdtime_Decode(t *testing.T) {
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "2006-01-02 15:04:05", time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTdtime{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTdtime.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTdtime.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTdtime_Encode(t *testing.T) {
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC), "2006-01-02 15:04:05"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTdtime{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTdtime.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTdtime_BinRead(t *testing.T) {
	tests := []struct {
		name    string
		v    []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{0,0,0,0,67,185,64,229}, time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTdtime{}
			gotV, err := c.BinRead(bytes.NewReader(tt.v))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTdtime.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTdtime.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTdtime_BinWrite(t *testing.T) {
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC), []byte{0,0,0,0,67,185,64,229}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTdtime{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTdtime.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != string(tt.wantW) {
				t.Errorf("CTdtime.BinWrite() = %v, want %v", []byte(gotW), tt.wantW)
			}
		})
	}
}

func TestCTdtimeZ_Decode(t *testing.T) {
	tloc, _ := time.Parse("-0700", "-1325")
	loc := tloc.Location()
	tests := []struct {
		name    string
		value string
		wantV   interface{}
		wantErr bool
	}{
		{"a", "2006-01-02 15:04:05 -13:25", time.Date(2006, 1, 2, 15, 4, 5, 0, loc), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTdtimeZ{}
			gotV, err := c.Decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CTdtimeZ.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTdtimeZ.Decode() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTdtimeZ_Encode(t *testing.T) {
	tloc, _ := time.Parse("-0700", "-1325")
	loc := tloc.Location()
	tests := []struct {
		name string
		value interface{}
		want string
	}{
		{"a", time.Date(2006, 1, 2, 15, 4, 5, 0, loc), "2006-01-02 15:04:05 -13:25"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTdtimeZ{}
			if got := c.Encode(tt.value); got != tt.want {
				t.Errorf("CTdtimeZ.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCTdtimeZ_BinRead(t *testing.T) {
	tloc, _ := time.Parse("-0700", "-1325")
	loc := tloc.Location()
	tests := []struct {
		name    string
		r []byte
		wantV   interface{}
		wantErr bool
	}{
		{"a", []byte{0,0,0,0,67,185,253,145,243,25}, time.Date(2006, 1, 2, 15, 4, 5, 0, loc), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTdtimeZ{}
			gotV, err := c.BinRead(bytes.NewReader(tt.r))
			if (err != nil) != tt.wantErr {
				t.Errorf("CTdtimeZ.BinRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("CTdtimeZ.BinRead() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCTdtimeZ_BinWrite(t *testing.T) {
	tloc, _ := time.Parse("-0700", "-1325")
	loc := tloc.Location()
	tests := []struct {
		name    string
		v interface{}
		wantW   []byte
		wantErr bool
	}{
		{"a", time.Date(2006, 1, 2, 15, 4, 5, 0, loc), []byte{0,0,0,0,67,185,253,145,243,25}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CTdtimeZ{}
			w := &bytes.Buffer{}
			if err := c.BinWrite(w, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("CTdtimeZ.BinWrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != string(tt.wantW) {
				t.Errorf("CTdtimeZ.BinWrite() = %v, want %v", []byte(gotW), tt.wantW)
			}
		})
	}
}
