package radgo

import (
	. "asdf"
)

type EAfrValue uint32

func (me EAfrValue) Tag() string {
	return "Acct-Authentic"
}

func (me EAfrValue) Begin() int {
	return int(aaaBegin)
}

func (me EAfrValue) End() int {
	return int(aaaEnd)
}

func (me EAfrValue) Int() int {
	return int(me)
}

func (me EAfrValue) IsGood() bool {
	return IsGoodEnum(me) && 
		len(aaaBind)==me.End() && 
		len(aaaBind[me]) > 0
}

func (me EAfrValue) ToString() string {
	var b EnumBinding = aaaBind[:]

	return b.EntryShow(me)
}

const (
	aaaBegin	EAfrValue = 1
	
	AaaRadius	EAfrValue = 1
	AaaLocal	EAfrValue = 2
	AaaRemote	EAfrValue = 3

	aaaEnd		EAfrValue = 4
)

var aaaBind = [aaaEnd]string{
	AaaRadius:	"RADIUS",
	AaaLocal:	"Local",
	AaaRemote:	"Remote",
}
