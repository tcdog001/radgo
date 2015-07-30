package radgo

import (
	. "asdf"
)

type EAtaValue uint32

func (me EAtaValue) Tag() string {
	return "Termination-Action"
}

func (me EAtaValue) Begin() int {
	return int(ataBegin)
}

func (me EAtaValue) End() int {
	return int(ataEnd)
}

func (me EAtaValue) Int() int {
	return int(me)
}

func (me EAtaValue) IsGood() bool {
	return IsGoodEnum(me) && 
		len(ataBind)==me.End() && 
		len(ataBind[me]) > 0
}

func (me EAtaValue) ToString() string {
	var b EnumBinding = ataBind[:]

	return b.EntryShow(me)
}

const (
	ataBegin			EAtaValue = 0
	
	AtaDefault 			EAtaValue = 0
	AtaRadiusRequest	EAtaValue = 1

	ataEnd 				EAtaValue = 2
)

var ataBind = [ataEnd]string{
	AtaDefault:			"Default",
	AtaRadiusRequest:	"RADIUS-Request",
}
