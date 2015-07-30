package radgo

import (
	. "asdf"
)

type EAastValue uint32

func (me EAastValue) Tag() string {
	return "Acct-Status-Type"
}

func (me EAastValue) Begin() int {
	return int(aastBegin)
}

func (me EAastValue) End() int {
	return int(aastEnd)
}

func (me EAastValue) Int() int {
	return int(me)
}

func (me EAastValue) IsGood() bool {
	return IsGoodEnum(me) && 
		len(aastBind)==me.End() && 
		len(aastBind[me]) > 0
}

func (me EAastValue) ToString() string {
	var b EnumBinding = aastBind[:]

	return b.EntryShow(me)
}

const (
	aastBegin			EAastValue = 1
	
	AastStart 			EAastValue = 1
	AastStop			EAastValue = 2
	AastInterimUpdate	EAastValue = 3
	AastAccountingOn	EAastValue = 4
	AastAccountingOff	EAastValue = 5
	AastFailed 			EAastValue = 15

	aastEnd 			EAastValue = 16
)

var aastBind = [aastEnd]string{
	AastStart:			"Start",
	AastStop:			"Stop",
	AastInterimUpdate:	"Interim-Update",
	AastAccountingOn:	"Accounting-On",
	AastAccountingOff:	"Accounting-Off",
	AastFailed:			"Failed",
}
