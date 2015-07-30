package radgo

import (
	. "asdf"
)

type EAfrValue uint32

func (me EAfrValue) Tag() string {
	return "Framed-Routing"
}

func (me EAfrValue) Begin() int {
	return int(afrBegin)
}

func (me EAfrValue) End() int {
	return int(afrEnd)
}

func (me EAfrValue) Int() int {
	return int(me)
}

func (me EAfrValue) IsGood() bool {
	return IsGoodEnum(me) && 
		len(afrBind)==me.End() && 
		len(afrBind[me]) > 0
}

func (me EAfrValue) ToString() string {
	var b EnumBinding = afrBind[:]

	return b.EntryShow(me)
}

// Framed-Routing value
const (
	afrBegin			EAfrValue = 0
	
	AfrNone				EAfrValue = 0
	AfrSend				EAfrValue = 1
	AfrListen 			EAfrValue = 2
	AfrSendAndListen	EAfrValue = 3

	afrEnd 				EAfrValue = 4
)

var afrBind = [afpEnd]string{
	AfrNone:			"None",
	AfrSend:			"Send routing packets",
	AfrListen:			"Listen for routing packets",
	AfrSendAndListen:	"Send and Listen",
}
