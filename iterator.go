package bintb

import "github.com/pkg/errors"

type RecordsIterator interface {
	Start() (state interface{}, err error)
	Done(state interface{}) (ok bool)
	Next(state interface{}) (rec Recorde, newState interface{}, err error)
}

func Iterate(it RecordsIterator) (err error) {
	var state interface{}
	for state, err = it.Start(); err == nil && !it.Done(state); _, state, err = it.Next(state) {
	}
	return
}

type RecordsIteratorOpener interface {
	Open() (header *ReadHeader, it RecordsIterator, err error)
}

func OpenIterate(opener RecordsIteratorOpener) (err error) {
	var it RecordsIterator
	if _, it, err = opener.Open(); err != nil {
		return errors.Wrap(err, "OpenIterate > Open() failed")
	}
	return Iterate(it)
}

func Each(it RecordsIterator, cb func(i int, r Recorde) error) (err error) {
	var (
		state interface{}
		rec   Recorde
		i     int
	)
	for state, err = it.Start(); err == nil && !it.Done(state); i++ {
		if rec, state, err = it.Next(state); err == nil {
			err = cb(i, rec)
		}
	}
	return
}
