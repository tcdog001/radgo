package radgo

import (
	. "strconv"
)

// fengxun Class
type AttrClass []byte

func (me AttrClass) UpRateMax() uint32 {
	s := string(me)[:8]
	
	if v, ok := Atoi(s); nil==ok {
		return uint32(v)
	}
	
	log.Error("bad attr class UpRateMax(%s)", s)
	
	return 0
}

func (me AttrClass) UpRateAvg() uint32 {
	s := string(me)[8:16]
	
	if v, ok := Atoi(s); nil==ok {
		return uint32(v)
	}
	
	log.Error("bad attr class UpRateAvg(%s)", s)
	
	return 0
}

func (me AttrClass) DownRateMax() uint32 {
	s := string(me)[16:24]
	
	if v, ok := Atoi(s); nil==ok {
		return uint32(v)
	}
	
	log.Error("bad attr class DownRateMax(%s)", s)
	
	return 0
}

func (me AttrClass) DownRateAvg() uint32 {
	s := string(me)[24:32]
	
	if v, ok := Atoi(s); nil==ok {
		return uint32(v)
	}
	
	log.Error("bad attr class DownRateAvg(%s)", s)
	
	return 0
}
