package bintb

import "io"

type ioWriter func(data []byte) (n int, err error)

func (this ioWriter) Write(p []byte) (n int, err error) {
	return this(p)
}

func IOWriter(f func(data []byte) (n int, err error)) io.Writer {
	return ioWriter(f)
}

