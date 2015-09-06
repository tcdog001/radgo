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
		len(pktCodeBind) == me.End() &&
		len(pktCodeBind[me]) > 0
}

func (me EPktCode) ToString() string {
	var b EnumBinding = pktCodeBind[:]

	return b.EntryShow(me)
}

func (me *EPktCode) FromString(Name string) error {
	if e, ok := pktCodeMap[Name]; ok {
		*me = e

		return nil
	}

	return ErrNoFound
}

func (me EPktCode) Match(Type EAttrType) EAttrTableValue {
	if me.IsGood() && Type.IsGood() {
		return attrTypeBind[Type].table[me]
	}

	return AttrTableZero
}

func (me EPktCode) IsMatch(Type EAttrType) bool {
	if AttrTableZero == me.Match(Type) {
		Log.Info("code %s and type %s should match, but is %s",
			me.ToString(),
			Type.ToString(),
			AttrTableZero.ToString())

		return false
	}

	return true
}

func (me EPktCode) IsMust(Type EAttrType) bool {
	return AttrTableOne == me.Match(Type)
}

func (me EPktCode) NetworkDir() ENetworkDir {
	return codeDirBind[me]
}

const (
	PktCodeBegin 		EPktCode = 1

	AccessRequest      	EPktCode = 1
	AccessAccept       	EPktCode = 2
	AccessReject       	EPktCode = 3
	AccountingRequest  	EPktCode = 4
	AccountingResponse 	EPktCode = 5
//	AccountingStatus	EPktCode = 6
//	PasswordRequest		EPktCode = 7
//	PasswordAck			EPktCode = 8
//	PasswordReject		EPktCode = 9
//	AccountingMessage	EPktCode = 10
	AccessChallenge    	EPktCode = 11

//	ResourceFreeRequest 	EPktCode = 21
//	ResourceFreeResponse	EPktCode = 22
//	ResourceQueryRequest	EPktCode = 23
//	ResourceQueryResponse	EPktCode = 24
//	AlternateResourceReclaimRequest EPktCode = 25
//	NasRebootRequest		EPktCode = 26
//	NasRebootResponse		EPktCode = 27
	
//	NextPasscode		EPktCode = 29
//	NewPin				EPktCode = 30
//	TerminateSession 	EPktCode = 31
//	PasswordExpired 	EPktCode = 32
//	EventRequest 		EPktCode = 33
//	EventResponse 		EPktCode = 34

	DisconnectRequest 	EPktCode = 40
	DisconnectAck 		EPktCode = 41
	DisconnectNak 		EPktCode = 42
//	ChangeFiltersRequest 	EPktCode = 43
//	ChangeFiltersAck 		EPktCode = 44
//	ChangeFiltersNak 		EPktCode = 45

//	IpAddressAllocate 	EPktCode = 50
//	IpAddressRelease 	EPktCode = 51

	PktCodeEnd 			EPktCode = 52
)

var pktCodeBind = [PktCodeEnd]string{
	AccessRequest:      "Access-Request",
	AccessAccept:       "Access-Accept",
	AccessReject:       "Access-Reject",
	AccessChallenge:    "Access-Challenge",
	AccountingRequest:  "Accounting-Request",
	AccountingResponse: "Accounting-Response",
	DisconnectRequest: 	"Disconnect-Request",
	DisconnectAck: 		"Disconnect-Ack",
	DisconnectNak: 		"Disconnect-Nak",
}

var pktCodeMap = map[string]EPktCode{}

var codeDirBind = [PktCodeEnd]ENetworkDir{
	AccessRequest:      ToServer,
	AccessAccept:       ToClient,
	AccessReject:       ToClient,
	AccessChallenge:    ToClient,
	AccountingRequest:  ToServer,
	AccountingResponse: ToClient,
	DisconnectRequest:	ToClient,
	DisconnectAck:		ToServer,
	DisconnectNak:		ToServer,
}

func initPktCode() {
	for i := PktCodeBegin; i < PktCodeEnd; i++ {
		pktCodeMap[pktCodeBind[i]] = i
	}
}
