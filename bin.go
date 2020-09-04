package bintb

import (
	"encoding/binary"
	"io"
)

var binOrder = binary.BigEndian

func br(r io.Reader, dst ...interface{}) error {
	for _, dst := range dst {
		if err := binary.Read(r, binOrder, dst); err != nil {
			return err
		}
	}
	return nil
}

func bw(w io.Writer, data ...interface{}) (err error) {
	for _, d := range data {
		if err = binary.Write(w, binOrder, d); err != nil {
			return
		}
	}
	return
}