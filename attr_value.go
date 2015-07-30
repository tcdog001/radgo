package radgo

import (
	. "asdf"
)

type EAttrValueType 	int

func (me EAttrValueType) Tag() string {
	return "Attribute-Value-Type"
}

func (me EAttrValueType) Begin() int {
	return int(avtBegin)
}

func (me EAttrValueType) End() int {
	return int(avtEnd)
}

func (me EAttrValueType) Int() int {
	return int(me)
}

func (me EAttrValueType) IsGood() bool {
	return IsGoodEnum(me) && 
		len(avtBind)==me.End() && 
		len(avtBind[me]) > 0
}

func (me EAttrValueType) IsNumber() bool {
	switch me {
		case AvtAddress:fallthrough
		case AvtInteger:fallthrough
		case AvtTime:
			return true
		default:
			return false
	}
}

func (me EAttrValueType) IsString() bool {
	return !me.IsNumber()
}

func (me EAttrValueType) ToString() string {
	var b EnumBinding = avtBind[:]

	return b.EntryShow(me)
}

const (
	avtBegin 		EAttrValueType = 1
	
	AvtText 		EAttrValueType = 1
	AvtString		EAttrValueType = 2
	AvtAddress		EAttrValueType = 3
	AvtInteger		EAttrValueType = 4
	AvtTime			EAttrValueType = 5
	
	avtEnd 			EAttrValueType = 6
)

var avtBind = [avtEnd]string{
	AvtText:	"Text",
	AvtString:	"String",
	AvtAddress:	"Address",
	AvtInteger:	"Integer",
	AvtTime:	"Time",
}
