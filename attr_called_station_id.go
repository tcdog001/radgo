package radgo

import (
	. "asdf"
)

// fengxun
//
// CalledStationId=apmac + ":" + ssid
// apmac as AA-BB-CC-DD-EE-FF(windows style)
//
// so, CalledStationId as AA-BB-CC-DD-EE-FF:SSID
type AttrCalledStationId []byte

func (me AttrCalledStationId) ApMac() []byte {
	return me[:17]
}

func (me AttrCalledStationId) SSID() []byte {
	return me[19:]
}

func MakeCalledStationId(mac []byte, ssid []byte) []byte {
	return []byte(Mac(mac).ToStringLU() + ":" + string(ssid))
}
