package radgo

import (
	. "asdf"
)

type EAfcValue uint32

func (me EAfcValue) Tag() string {
	return "Framed-Compression"
}

func (me EAfcValue) Begin() int {
	return int(afcBegin)
}

func (me EAfcValue) End() int {
	return int(afcEnd)
}

func (me EAfcValue) Int() int {
	return int(me)
}

func (me EAfcValue) IsGood() bool {
	return IsGoodEnum(me) && 
		len(afcBind)==me.End() && 
		len(afcBind[me]) > 0
}

func (me EAfcValue) ToString() string {
	var b EnumBinding = afcBind[:]

	return b.EntryShow(me)
}

func (me *EAfcValue) FromString(Name string) error {
	if e, ok := afcMap[Name]; ok {
		*me = e
		
		return nil
	}

	return ErrNoFound
}

const (
	afcBegin 		EAfcValue = 0
	
	AfcNone			EAfcValue = 0
	AfcTcpip		EAfcValue = 1
	AfcIpx			EAfcValue = 2
	AfcLzs 			EAfcValue = 3

	afcEnd 			EAfcValue = 4
)

var afcBind = [afcEnd]string{
	AfcNone:	"None",
	AfcTcpip:	"VJ TCP/IP header compression",
	AfcIpx:		"IPX header compression",
	AfcLzs:		"Stac-LZS compression",
}

var afcMap = map[string]EAfcValue{}

func initAfc() {
	for i:=afcBegin; i<afcEnd; i++ {
		afcMap[afcBind[i]] = i
	}
}
