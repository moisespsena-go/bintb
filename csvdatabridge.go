package bintb

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/moisespsena-go/logging"
	path_helpers "github.com/moisespsena-go/path-helpers"

	"github.com/moisespsena-go/task"
)


var (
	log = logging.GetOrCreateLogger(path_helpers.GetCalledDir())
)

type CsvDataBridge struct {
	r          *csv.Reader
	w          io.Writer
	csvReader  *csv.Reader
	cache      *Cache
	UidColumns []int
	Columns    []*Column
	close      chan struct{}
	Log        logging.Logger
	encdec     *Decoder
}

func New(cache *Cache, columns []*Column, r ...*csv.Reader) *CsvDataBridge {
	return &CsvDataBridge{
		cache:   cache,
		Columns: columns,
		Log:     logging.WithPrefix(log, "data bridge"),
		encdec:  NewDecoder(columns, context.Background()),
	}
}

func (this *CsvDataBridge) Stop() {
	close(this.close)
}

func (this *CsvDataBridge) IsRunning() bool {
	select {
	case <-this.close:
		return true
	default:
		return false
	}
}

func (this *CsvDataBridge) Start(done func()) (stop task.Stoper, err error) {
	r, w := io.Pipe()
	this.w = w
	this.csvReader = csv.NewReader(r)
	this.close = make(chan struct{})
	return this, nil
}

func (this *CsvDataBridge) readLine() (rec Recorde, err error) {
	var record []string
	if record, err = this.csvReader.Read(); err != nil {
		return
	}

	if len(record) != len(this.Columns) {
		return nil, fmt.Errorf("bad csv columns count. Expected %d, but get %d. %v", len(this.Columns), len(record), record)
	}

	rec, err = this.encdec.Decode(record...)
	return
}

func (this *CsvDataBridge) readStoreLine() (err error) {
	var rec Recorde
	if rec, err = this.readLine(); err != nil {
		if err == io.EOF {
			close(this.close)
			return
		}
		this.Log.Error(err)
		return nil
	}

	var uidRec = make([]string, len(this.UidColumns))
	for i, j := range this.UidColumns {
		uidRec[i] = fmt.Sprint(rec[j])
	}
	uid := strings.Join(uidRec, ":")
	if uid == "" {
		uid = strconv.Itoa(int(time.Now().UnixNano()))
	}

	this.cache.Store(uid, rec)
	return
}

func (this *CsvDataBridge) loop() (err error) {
	for {
		select {
		case <-this.close:
			return io.EOF
		default:
			if err = this.readStoreLine(); err != nil {
				return
			}
		}
	}
}
