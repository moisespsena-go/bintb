package bintb

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/moisespsena-go/logging"

	"github.com/moisespsena-go/bid"
)

type Cache struct {
	DB    *sql.DB
	Table string
	Qinsert,
	Qshow,
	Qdelete string
	log logging.Logger
}

func NewCache(DB *sql.DB, Table string) *Cache {
	return &Cache{
		DB:      DB,
		Table:   Table,
		Qinsert: "INSERT INTO " + Table + " (sys, uid, data) VALUES (?,?,?)",
		Qshow:   "SELECT sys, uid, data FROM " + Table + " ORDER BY sys ASC",
		Qdelete: "DELETE FROM " + Table + " WHERE sys IN ?",
		log:     logging.WithPrefix(log, "cache"),
	}
}

func (this *Cache) Init() (err error) {
	_, err = this.DB.Exec(`
CREATE TABLE IF NOT EXISTS ` + this.Table + ` (
  sys CHAR(12) PRIMARY KEY,
  uid VARCHAR(32), 
  data TEXT, 
  UNIQUE (uid)
)`)
	return
}

func (this *Cache) Store(uid string, rec Recorde) (sysId bid.BID) {
	sysId = bid.New()
	if _, err := this.DB.Exec(this.Qinsert, sysId, uid); err != nil {
		this.log.Error("store "+fmt.Sprintf("[sys=%s, uid=%v] `%s`", sysId.Hex(), uid, rec.String())+"` failed: ", err.Error())
	}
	return
}

type Recorde []interface{}

func (this Recorde) String() string {
	b, _ := json.Marshal(this)
	return string(b)
}
