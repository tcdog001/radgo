package radgo

import (
	. "asdf"
)

type EAfpValue uint32

func (me EAfpValue) Tag() string {
	return "Framed-Protocol"
}

func (me EAfpValue) Begin() int {
	return int(afpBegin)
}

func (me EAfpValue) End() int {
	return int(afpEnd)
}

func (me EAfpValue) Int() int {
	return int(me)
}

func (me EAfpValue) IsGood() bool {
	return IsGoodEnum(me) &&
		len(afpBind) == me.End() &&
		len(afpBind[me]) > 0
}

func (me EAfpValue) ToString() string {
	var b EnumBinding = afpBind[:]

	return b.EntryShow(me)
}

func (me *EAfpValue) FromString(Name string) error {
	if e, ok := afpMap[Name]; ok {
		*me = e

		return nil
	}

	return ErrNoFound
}

const (
	afpBegin EAfpValue = 1

	AfpPpp     EAfpValue = 1
	AfpSlip    EAfpValue = 2
	AfpArap    EAfpValue = 3 // AppleTalk Remote Access Protocol (ARAP)
	AfpGandalf EAfpValue = 4 // Gandalf proprietary SingleLink/MultiLink protocol
	AfpIpx     EAfpValue = 5 // Xylogics proprietary IPX/SLIP
	AfpX75     EAfpValue = 6 // X.75 Synchronous

	afpEnd EAfpValue = 7
)

var afpBind = [afpEnd]string{
	AfpPpp:     "PPP",
	AfpSlip:    "SLIP",
	AfpArap:    "AppleTalk Remote Access Protocol (ARAP)",
	AfpGandalf: "Gandalf proprietary SingleLink/MultiLink protocol",
	AfpIpx:     "Xylogics proprietary IPX/SLIP",
	AfpX75:     "X.75 Synchronous",
}

var afpMap = map[string]EAfpValue{}

func initAfp() {
	for i := afpBegin; i < afpEnd; i++ {
		afpMap[afpBind[i]] = i
	}
}
