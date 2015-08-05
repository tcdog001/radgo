package radgo

import (
	. "asdf"
)

type EAaaValue uint32

func (me EAaaValue) Tag() string {
	return "Acct-Authentic"
}

func (me EAaaValue) Begin() int {
	return int(aaaBegin)
}

func (me EAaaValue) End() int {
	return int(aaaEnd)
}

func (me EAaaValue) Int() int {
	return int(me)
}

func (me EAaaValue) IsGood() bool {
	return IsGoodEnum(me) &&
		len(aaaBind) == me.End() &&
		len(aaaBind[me]) > 0
}

func (me EAaaValue) ToString() string {
	var b EnumBinding = aaaBind[:]

	return b.EntryShow(me)
}

func (me *EAaaValue) FromString(Name string) error {
	if e, ok := aaaMap[Name]; ok {
		*me = e

		return nil
	}

	return ErrNoFound
}

const (
	aaaBegin EAaaValue = 1

	AaaRadius EAaaValue = 1
	AaaLocal  EAaaValue = 2
	AaaRemote EAaaValue = 3

	aaaEnd EAaaValue = 4
)

var aaaBind = [aaaEnd]string{
	AaaRadius: "RADIUS",
	AaaLocal:  "Local",
	AaaRemote: "Remote",
}

var aaaMap = map[string]EAaaValue{}

func initAaa() {
	for i := aaaBegin; i < aaaEnd; i++ {
		aaaMap[aaaBind[i]] = i
	}
}
