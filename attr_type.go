package radgo

import (
	. "asdf"
)

type EAttrType byte

func (me EAttrType) Tag() string {
	return "Attribute-Type"
}

func (me EAttrType) Begin() int {
	return int(AttrTypeBegin)
}

func (me EAttrType) End() int {
	return int(AttrTypeEnd)
}

func (me EAttrType) Int() int {
	return int(me)
}

func (me EAttrType) IsGood() bool {
	return me >= AttrTypeBegin && me < AttrTypeEnd &&
		len(attrTypeBind)==me.End() &&
		nil != attrTypeBind[me]
}

func (me EAttrType) ToString() string {
	if me.IsGood() {
		return attrTypeBind[me].name
	}
	
	return Empty
}

func (me *EAttrType) FromString(s string) error {
	v, ok := attrTypeMap[s]
	if !ok {
		return Error
	}
	
	*me = v.idx
	return nil
}

func (me EAttrType) ValueType() EAttrValueType {
	if !me.IsGood() || nil==attrTypeBind[me] {
		return AvtString
	}
	
	return attrTypeBind[me].avt
}

func (me EAttrType) IsGoodLength(Len byte) bool {
	return me.IsGood() &&
		Len >= (attrTypeBind[me].min + 2) &&
		Len <= (attrTypeBind[me].max + 2)
}

func (me EAttrType) IsGoodValue(Value uint32) bool {
	if !me.IsGood() {
		return false
	}
	
	if me.ValueType().IsNumber() {
		switch me {
			case AcctStatusType:
				if !EAastValue(Value).IsGood() {
					return false
				}
			case AcctTerminateCause:
				if !EAtcValue(Value).IsGood() {
					return false
				}
			case NasPortType:
				if !EAnptValue(Value).IsGood() {
					return false
				}
			case ServiceType:
				if !EAstValue(Value).IsGood() {
					return false
				}
		}
	}

	return true
}

const (
	AttrTypeBegin 		EAttrType = 1
	
	UserName 			EAttrType = 1
	UserPassword 		EAttrType = 2
	ChapPassword		EAttrType = 3
	NasIpAddress		EAttrType = 4
	NasPort				EAttrType = 5
	ServiceType			EAttrType = 6
	FramedIpAddress		EAttrType = 8
	ReplyMessage		EAttrType = 18
	Class				EAttrType = 25
	SessionTimeout		EAttrType = 27
	IdleTimeout			EAttrType = 28
	CalledStationId		EAttrType = 30
	CallingStationId	EAttrType = 31
	NasIdentifier		EAttrType = 32
	AcctStatusType		EAttrType = 40	
	AcctDelayTime		EAttrType = 41
	AcctInputOctets		EAttrType = 42
	AcctOutputOctets	EAttrType = 43
	AcctSessionId		EAttrType = 44
	AcctSessionTime		EAttrType = 46
	AcctTerminateCause	EAttrType = 49
	AcctInputGigawords	EAttrType = 52
	AcctOutputGigawords EAttrType = 53
	EventTimestamp		EAttrType = 55
	ChapChallenge		EAttrType = 60
	NasPortType			EAttrType = 61
	NasPortId			EAttrType = 87
	
	AttrTypeEnd 		EAttrType = 88
)

type attrType struct {
	idx 		EAttrType
	avt 		EAttrValueType
	min 		byte // just value size
	max 		byte // just value size
	name 		string
	table 		[PktCodeEnd]EAttrTableValue
}

var attrTypeMap = map[string]*attrType{}

var attrTypeBind = [AttrTypeEnd]*attrType{
	UserName:&attrType{
		name:"User-Name",
		avt:AvtString,
		max:63,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZeroOne,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,
		},
	},
	UserPassword:&attrType{
		name:"User-Password",
		avt:AvtString,
		min:16,
		max:128,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZero,
			AccountingResponse:	AttrTableZero,		
		},
	},
	ChapPassword:&attrType{
		name:"CHAP-Password",
		avt:AvtString,
		min:17,
		max:17,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZero,
			AccountingResponse:	AttrTableZero,		
		},
	},
	NasIpAddress:&attrType{
		name:"NAS-IP-Address",
		avt:AvtAddress,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	NasPort:&attrType{
		name:"NAS-Port",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	ServiceType:&attrType{
		name:"Service-Type",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZeroOne,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	FramedIpAddress:&attrType{
		name:"Framed-IP-Address",
		avt:AvtAddress,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZeroOne,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	ReplyMessage:&attrType{
		name:"Reply-Message",
		avt:AvtText,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZeroMore,
			AccessReject:		AttrTableZeroMore,
			AccessChallenge:	AttrTableZeroMore,
			AccountingRequest:	AttrTableZero,
			AccountingResponse:	AttrTableZero,		
		},
	},
	Class:&attrType{
		name:"Class",
		avt:AvtString,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZeroMore,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroMore,
			AccountingResponse:	AttrTableZero,		
		},
	},
	SessionTimeout:&attrType{
		name:"Session-Timeout",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZeroOne,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZeroOne,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	IdleTimeout:&attrType{
		name:"Idle-Timeout",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZeroOne,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZeroOne,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	CalledStationId:&attrType{
		name:"Called-Station-Id",
		avt:AvtString,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	CallingStationId:&attrType{
		name:"Calling-Station-Id",
		avt:AvtString,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,	
		},
	},
	NasIdentifier:&attrType{
		name:"NAS-Identifier",
		avt:AvtString,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	
	EventTimestamp:&attrType{
		name:"Event-Timestamp",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	ChapChallenge:&attrType{
		name:"CHAP-Challenge",
		avt:AvtString,
		min:5,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZero,
			AccountingResponse:	AttrTableZero,		
		},
	},
	NasPortType:&attrType{
		name:"NAS-Port-Type",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	NasPortId:&attrType{
		name:"NAS-Port-Id",
		avt:AvtText,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZeroOne,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZero,
			AccountingResponse:	AttrTableZero,		
		},
	},

	AcctStatusType:&attrType{
		name:"Acct-Status-Type",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	AcctDelayTime:&attrType{
		name:"Acct-Delay-Time",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	AcctInputOctets:&attrType{
		name:"Acct-Input-Octets",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	AcctOutputOctets:&attrType{
		name:"Acct-Output-Octets",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	AcctSessionId:&attrType{
		name:"Acct-Session-Id",
		avt:AvtText,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	AcctSessionTime:&attrType{
		name:"Acct-Session-Time",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	AcctTerminateCause:&attrType{
		name:"Acct-Terminate-Cause",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	AcctInputGigawords:&attrType{
		name:"Acct-Input-Gigawords",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
	AcctOutputGigawords:&attrType{
		name:"Acct-Output-Gigawords",
		avt:AvtInteger,
		table:[PktCodeEnd]EAttrTableValue{
			AccessRequest:		AttrTableZero,
			AccessAccept:		AttrTableZero,
			AccessReject:		AttrTableZero,
			AccessChallenge:	AttrTableZero,
			AccountingRequest:	AttrTableZeroOne,
			AccountingResponse:	AttrTableZero,		
		},
	},
}

func initAttrType(){
	for idx := AttrTypeBegin; idx<AttrTypeEnd; idx++ {
		v := attrTypeBind[idx]
		if nil==v {
			continue
		}
		v.idx = idx
		
		attrTypeMap[v.name] = v
		
		switch v.avt {
		case AvtText:fallthrough
		case AvtString:
			if 0==v.min {
				v.min = 1
			}
			if 0==v.max {
				v.max = AttrLengthMax - 2
			}
		case AvtInteger: fallthrough
		case AvtAddress: fallthrough
		case AvtTime:
			if 0==v.min {
				v.min = 4
			}
			if 0==v.max {
				v.max = 4
			}
		}
	}
}