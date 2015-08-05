package radgo

import (
	. "asdf"
)

type EAlsValue uint32

func (me EAlsValue) Tag() string {
	return "Login-Service"
}

func (me EAlsValue) Begin() int {
	return int(alsBegin)
}

func (me EAlsValue) End() int {
	return int(alsEnd)
}

func (me EAlsValue) Int() int {
	return int(me)
}

func (me EAlsValue) IsGood() bool {
	return IsGoodEnum(me) &&
		len(alsBind) == me.End() &&
		len(alsBind[me]) > 0
}

func (me EAlsValue) ToString() string {
	var b EnumBinding = alsBind[:]

	return b.EntryShow(me)
}

func (me *EAlsValue) FromString(Name string) error {
	if e, ok := alsMap[Name]; ok {
		*me = e

		return nil
	}

	return ErrNoFound
}

const (
	alsBegin EAlsValue = 0

	AlsTelnet     EAlsValue = 0
	AlsRlogin     EAlsValue = 1
	AlsTcp        EAlsValue = 2
	AlsPortMaster EAlsValue = 3
	AlsLat        EAlsValue = 4
	AlsX25Pad     EAlsValue = 5
	AlsX25T3Pos   EAlsValue = 6
	AlsTcpQuiet   EAlsValue = 7

	alsEnd EAlsValue = 8
)

var alsBind = [alsEnd]string{
	AlsTelnet:     "Telnet",
	AlsRlogin:     "Rlogin",
	AlsTcp:        "TCP Clear",
	AlsPortMaster: "PortMaster",
	AlsLat:        "Lat",
	AlsX25Pad:     "X25-PAD",
	AlsX25T3Pos:   "X25-T3POS",
	AlsTcpQuiet:   "TCP Clear Quiet",
}

var alsMap = map[string]EAlsValue{}

func initAls() {
	for i := alsBegin; i < alsEnd; i++ {
		alsMap[alsBind[i]] = i
	}
}
