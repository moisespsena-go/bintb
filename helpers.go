package bintb

func NextRecords(recs []Recorde) (next func() (rec Recorde, err error)) {
	var (
		i int
		l = len(recs)
	)
	return func() (rec Recorde, err error) {
		if i == l {
			return
		}
		rec = recs[i]
		i++
		return
	}
}
