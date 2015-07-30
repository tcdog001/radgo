package radgo

import (
	. "asdf"
)

type EAttrTableValue byte

func (me EAttrTableValue) Tag() string {
	return "Attribute-Table"
}

func (me EAttrTableValue) Begin() int {
	return int(attrTableBegin)
}

func (me EAttrTableValue) End() int {
	return int(attrTableEnd)
}

func (me EAttrTableValue) Int() int {
	return int(me)
}

func (me EAttrTableValue) IsGood() bool {
	return IsGoodEnum(me) && 
		len(attrTableBind)==me.End() && 
		len(attrTableBind[me]) > 0
}

func (me EAttrTableValue) ToString() string {
	var b EnumBinding = attrTableBind[:]

	return b.EntryShow(me)
}

// 0     This attribute MUST NOT be present in packet.
// 0+    Zero or more instances of this attribute MAY be present in packet.
// 0-1   Zero or one instance of this attribute MAY be present in packet.
// 1     Exactly one instance of this attribute MUST be present in packet.
const (
	attrTableBegin 		EAttrTableValue = 0
	
	AttrTableZero		EAttrTableValue = 0
	AttrTableZeroMore	EAttrTableValue = 1
	AttrTableZeroOne	EAttrTableValue = 2
	AttrTableOne		EAttrTableValue = 3

	attrTableEnd		EAttrTableValue = 4
)

var attrTableBind = [attrTableEnd]string{
	AttrTableZero:		"Zero",
	AttrTableZeroMore:	"ZeroOrMore",
	AttrTableZeroOne:	"ZeroOrOne",
	AttrTableOne:		"One",
}
