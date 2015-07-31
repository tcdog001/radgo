package radgo

import (
	. "asdf"
)

type EPktCode byte

func (me EPktCode) Tag() string {
	return "Packet-Code"
}

func (me EPktCode) Begin() int {
	return int(PktCodeBegin)
}

func (me EPktCode) End() int {
	return int(PktCodeEnd)
}

func (me EPktCode) Int() int {
	return int(me)
}

func (me EPktCode) IsGood() bool {
	return IsGoodEnum(me) && 
		len(pktCodeBind)==me.End() && 
		len(pktCodeBind[me]) > 0
}

func (me EPktCode) ToString() string {
	var b EnumBinding = pktCodeBind[:]

	return b.EntryShow(me)
}

func (me EPktCode) Match(Type EAttrType) EAttrTableValue {
	if me.IsGood() && Type.IsGood() {
		return attrTypeBind[Type].table[me]
	}
	
	return AttrTableZero
}

func (me EPktCode) IsMatch(Type EAttrType) bool {
	if AttrTableZero==me.Match(Type) {
		log.Info("code %s and type %s should match, but is %s",
			me.ToString(), 
			Type.ToString(),
			AttrTableZero.ToString())
		
		return false
	}
	
	return true
}

func (me EPktCode) IsMust(Type EAttrType) bool {
	return AttrTableOne==me.Match(Type)
}

func (me EPktCode) NetworkDir() ENetworkDir {
	return codeDirBind[me]
}

const (
	PktCodeBegin 		EPktCode = 1
	
	AccessRequest		EPktCode = 1
	AccessAccept		EPktCode = 2
	AccessReject		EPktCode = 3
	AccountingRequest	EPktCode = 4
	AccountingResponse	EPktCode = 5
	AccessChallenge		EPktCode = 11

	PktCodeEnd			EPktCode = 12
)

var pktCodeBind = [PktCodeEnd]string{
	AccessRequest:		"Access-Request",
	AccessAccept:		"Access-Accept",
	AccessReject:		"Access-Reject",
	AccessChallenge:	"Access-Challenge",
	AccountingRequest:	"Accounting-Request",
	AccountingResponse:	"Accounting-Response",
}

var codeDirBind = [PktCodeEnd]ENetworkDir{
	AccessRequest:		ToServer,
	AccessAccept:		ToClient,
	AccessReject:		ToClient,
	AccessChallenge:	ToClient,
	AccountingRequest:	ToServer,
	AccountingResponse:	ToClient,
}

