package bintb

func JsonStreamWriteF(config *JsonStreamWriterConfig, next func() (rec Recorde, err error)) (err error) {
	csw := NewJsonStreamWriter(config)

	defer func() {
		if err == nil {
			err = csw.Close()
		} else {
			csw.Close()
		}
	}()

	var rec Recorde
	for {
		if rec, err = next(); err != nil {
			return
		}
		if rec == nil {
			break
		}

		if err = csw.Write(rec); err != nil {
			return
		}
	}
	return
}