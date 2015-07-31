package radgo

import (
	. "strconv"
)

// fengxun Class
type AttrClass []byte

func (me AttrClass) UpRateMax() uint32 {
	if v, ok := Atoi(string(me)[:8]); nil!=ok {
		return uint32(v)
	}
	
	return 0
}

func (me AttrClass) UpRateAvg() uint32 {
	if v, ok := Atoi(string(me)[8:16]); nil!=ok {
		return uint32(v)
	}
	
	return 0
}

func (me AttrClass) DownRateMax() uint32 {
	if v, ok := Atoi(string(me)[16:24]); nil!=ok {
		return uint32(v)
	}
	
	return 0
}

func (me AttrClass) DownRateAvg() uint32 {
	if v, ok := Atoi(string(me)[16:24]); nil!=ok {
		return uint32(v)
	}
	
	return 0
}
