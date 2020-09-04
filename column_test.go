package bintb

import (
	"testing"
)

func TestParseColumn(t *testing.T) {
	tests := []struct {
		name    string
		def     string
		wantC   Column
		wantErr bool
	}{
		{"maxLength", "name[32]:s", Column{Name: "name", maxLength: 32, typ: CtString}, false},
		{"maxLength", "name*:s", Column{Name: "name", typ: CtString, flags: Required}, false},
		{"maxLength", "name+:s", Column{Name: "name", typ: CtString, flags: Unique}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := ParseColumn(tt.def)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseColumn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotC.String() != tt.wantC.String() {
				t.Errorf("ParseColumn() = %v, want %v", gotC.String(), tt.wantC.String())
			}
		})
	}
}
