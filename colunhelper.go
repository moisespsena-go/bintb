package bintb

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ParseColumn(def string) (c *Column, err error) {
	parts := strings.Split(def, ":")
	if len(parts) != 2 {
		err = errors.New("bad column format")
		return
	}

	var (
		name = parts[0]
		typ  = ColumnType(parts[1])
		clen uint16
	)
	var flags Flags
	ln := len(name) - 1
	switch name[ln] {
	case '*':
		flags = flags.Set(Required)
		name = name[:ln]
	case '+':
		flags = flags.Set(Unique)
		name = name[:ln]
	}
	if lpos := strings.IndexRune(name, '['); lpos > 0 {
		rpos := strings.IndexRune(name, ']')
		if rpos < lpos+1 {
			err = errors.New("bad column maxLength format")
			return
		}
		csize := name[lpos+1 : rpos]
		var i uint64
		if i, err = strconv.ParseUint(csize, 10, 16); err != nil {
			err = errors.Wrap(err, "bad column maxLength")
			return
		}
		clen = uint16(i)
		name = name[0:lpos]
	}
	c = NewColumn(name, typ, flags)
	c.maxLength = clen
	return
}

func MustParseColumn(def string) (c *Column) {
	var err error
	if c, err = ParseColumn(def); err != nil {
		log.Fatal("MustParseColumn:", err)
	}
	return
}

func ParseColumns(defs ...string) (columns []*Column, err error) {
	columns = make([]*Column, len(defs))
	for i, def := range defs {
		if columns[i], err = ParseColumn(def); err != nil {
			return nil, errors.Wrapf(err, "parse column#%d `%s`", i, def)
		}
	}
	return
}

func MustParseColumns(def ...string) (columns []*Column) {
	var err error
	if columns, err = ParseColumns(def...); err != nil {
		log.Fatal("MustParseColumns:", err)
	}
	return
}
