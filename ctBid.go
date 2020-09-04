package bintb

import (
	"io"

	"cloud.google.com/go/civil"

	"github.com/moisespsena-go/bid"
)

type CTbid struct{}

func (CTbid) Limited() bool {
	return false
}

func (CTbid) Zero() interface{} {
	return bid.BID{}
}

func (CTbid) Decode(value string) (v interface{}, err error) {
	var bid bid.BID
	if err = bid.Scan(value); err != nil {
		return
	}
	v = bid
	return
}

func (CTbid) Encode(value interface{}) string {
	return value.(bid.BID).B64()
}

func (CTbid) BinRead(r io.Reader) (v interface{}, err error) {
	var bid = make(bid.BID, 12)
	if _, err = r.Read(bid); err != nil {
		return
	}
	v = bid
	return
}

func (CTbid) BinWrite(w io.Writer, v interface{}) (err error) {
	_, err = w.Write(v.(bid.BID))
	return
}

func (this *CTbid) DateOf(value interface{}) civil.Date {
	return civil.DateOf(value.(bid.BID).Time())
}

const (
	CtBid ColumnType = "bid"
)

func init() {
	ColumnTypeTool.Set(CtBid, CTbid{})
}
