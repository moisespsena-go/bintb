package bintb

import (
	"fmt"
	"io"
	"reflect"
	"sort"
)

type ColumnTool interface {
	Limited() bool
	Zero() interface{}
	Decode(value string) (v interface{}, err error)
	Encode(value interface{}) string
	BinRead(r io.Reader) (v interface{}, err error)
	BinWrite(w io.Writer, v interface{}) (err error)
}

type LimitColumnTool struct{}

func (LimitColumnTool) Limited() bool {
	return true
}

type UnLimitColumnTool struct{}

func (UnLimitColumnTool) Limited() bool {
	return false
}

type columnTypeTool struct {
	data    map[ColumnType]ColumnTool
	alias   map[ColumnType]ColumnType
	aliases map[ColumnType][]ColumnType
}

func newColumnTypeTool() *columnTypeTool {
	return &columnTypeTool{map[ColumnType]ColumnTool{}, map[ColumnType]ColumnType{}, map[ColumnType][]ColumnType{}}
}

func (this *columnTypeTool) Pairs(typ ...ColumnType) (pairs [][]string) {
	if len(typ) == 0 {
		for k := range this.data {
			pairs = append(pairs, []string{string(k), this.Description(k)})
		}
	} else {
		for _, k := range typ {
			pairs = append(pairs, []string{string(k), this.Description(k)})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	return
}

func (this *columnTypeTool) Description(key ColumnType) (desc string) {
	tool := this.Get(key)
	if tool == nil {
		return
	}
	desc = reflect.TypeOf(tool.Zero()).String()
	if d, ok := tool.(interface{ TypeName() string }); ok {
		desc = d.TypeName()
	}
	if d, ok := tool.(interface{ Description() string }); ok {
		desc += ": " + d.Description()
	}
	return
}

func (this *columnTypeTool) Each(cb func(typ ColumnType, too ColumnTool)) {
	for k, v := range this.data {
		cb(k, v)
	}
}

func (this *columnTypeTool) Set(typ ColumnType, tool ColumnTool, alias ...interface{}) {
	if _, ok := this.data[typ]; ok {
		log.Fatalf("ColumnTypeTool.Set: type %q has be registered", typ)
	}
	this.data[typ] = tool
	for _, alias := range alias {
		var aliasTyp ColumnType
		switch t := alias.(type) {
		case ColumnType:
			aliasTyp = typ
		case string:
			aliasTyp = ColumnType(t)
		default:
			aliasTyp = ColumnType(fmt.Sprint(t))
		}
		if _, ok := this.alias[aliasTyp]; ok {
			log.Fatalf("ColumnTypeTool.Set: alias %q has be registered", aliasTyp)
		}
		this.alias[aliasTyp] = typ
		if aliases, ok := this.aliases[typ]; ok {
			this.aliases[typ] = append(aliases, aliasTyp)
		} else {
			this.aliases[typ] = []ColumnType{aliasTyp}
		}
	}
}

func (this *columnTypeTool) Resolve(key interface{}) (typ ColumnType, ok bool) {
	switch t := key.(type) {
	case ColumnType:
		if _, ok = this.data[t]; ok {
			typ = t
			return
		}
		typ, ok = this.alias[t]
	case string:
		typ, ok = this.alias[ColumnType(t)]
	default:
		typ, ok = this.alias[ColumnType(fmt.Sprint(t))]
	}
	return
}

func (this *columnTypeTool) GetAliases(key ColumnType) []ColumnType {
	return this.aliases[key]
}

func (this *columnTypeTool) Get(key ColumnType) ColumnTool {
	if key, ok := this.Resolve(key); ok {
		return this.data[key]
	}
	return nil
}

var ColumnTypeTool = newColumnTypeTool()