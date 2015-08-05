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
	if !IsGoodEnum(me) {
		log.Error("bad attr(%s) value(%d)", me.Tag(), me)

		return false
	} else if 0 == len(aastBind[me]) {
		log.Error("no support attr(%s) value(%d)", me.Tag(), me)

		return false
	}

	return true
}

func (me EAastValue) ToString() string {
	var b EnumBinding = aastBind[:]

	return b.EntryShow(me)
}

func (me *EAastValue) FromString(Name string) error {
	if e, ok := aastMap[Name]; ok {
		*me = e

		return nil
	}

	return ErrNoFound
}

const (
	aastBegin EAastValue = 1

	AastStart         EAastValue = 1
	AastStop          EAastValue = 2
	AastInterimUpdate EAastValue = 3
	AastAccountingOn  EAastValue = 4
	AastAccountingOff EAastValue = 5
	AastFailed        EAastValue = 15

	aastEnd EAastValue = 16
)

var aastBind = [aastEnd]string{
	AastStart:         "Start",
	AastStop:          "Stop",
	AastInterimUpdate: "Interim-Update",
	AastAccountingOn:  "Accounting-On",
	AastAccountingOff: "Accounting-Off",
	AastFailed:        "Failed",
}

var aastMap = map[string]EAastValue{}

func initAast() {
	for i := aastBegin; i < aastEnd; i++ {
		aastMap[aastBind[i]] = i
	}
}
