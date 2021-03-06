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
	if me < AttrTypeBegin || me > AttrTypeEnd {
		Log.Error("attr type is %d, should [%d, %d)",
			me,
			AttrTypeBegin,
			AttrTypeEnd)

		return false
	}

	return nil != attrTypeBind[me]
}

func (me EAttrType) ToString() string {
	if !me.IsGood() {
		return Empty
	}

	return attrTypeBind[me].name
}

func (me *EAttrType) FromString(s string) error {
	v, ok := attrTypeMap[s]
	if !ok {
		Log.Error("bad attr type string", s)

		return Error
	}

	*me = v.idx
	return nil
}

func (me EAttrType) ValueType() EAttrValueType {
	if !me.IsGood() {
		return AvtString
	}

	return attrTypeBind[me].avt
}

func (me EAttrType) IsGoodLength(Len byte) bool {
	if !me.IsGood() {
		return false
	} else if Len < (attrTypeBind[me].min+2) ||
		Len > (attrTypeBind[me].max+2) {
		Log.Error("attr(%s) Len is %d, should [%d, %d]",
			me.ToString(),
			Len-2,
			attrTypeBind[me].min,
			attrTypeBind[me].max)

		return false
	}

	return true
}

func (me EAttrType) IsGoodNumber(Value uint32) bool {
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

	UserName            EAttrType = 1
	UserPassword        EAttrType = 2
	ChapPassword        EAttrType = 3
	NasIpAddress        EAttrType = 4
	NasPort             EAttrType = 5
	ServiceType         EAttrType = 6
	FramedIpAddress     EAttrType = 8
	ReplyMessage        EAttrType = 18
	Class               EAttrType = 25
	SessionTimeout      EAttrType = 27
	IdleTimeout         EAttrType = 28
	CalledStationId     EAttrType = 30
	CallingStationId    EAttrType = 31
	NasIdentifier       EAttrType = 32
	AcctStatusType      EAttrType = 40
	AcctDelayTime       EAttrType = 41
	AcctInputOctets     EAttrType = 42
	AcctOutputOctets    EAttrType = 43
	AcctSessionId       EAttrType = 44
	AcctSessionTime     EAttrType = 46
	AcctTerminateCause  EAttrType = 49
	AcctInputGigawords  EAttrType = 52
	AcctOutputGigawords EAttrType = 53
	EventTimestamp      EAttrType = 55
	ChapChallenge       EAttrType = 60
	NasPortType         EAttrType = 61
	NasPortId           EAttrType = 87
	
	ErrorCause			EAttrType = 101
	
	AttrTypeEnd 		EAttrType = 102
)

type attrType struct {
	idx   EAttrType
	avt   EAttrValueType
	min   byte // just value size
	max   byte // just value size
	name  string
	table [PktCodeEnd]EAttrTableValue
}

var attrTypeMap = map[string]*attrType{}

var attrTypeBind = [AttrTypeEnd]*attrType{
	UserName: &attrType{
		name: "User-Name",
		avt:  AvtString,
		max:  63,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccessAccept:       AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	UserPassword: &attrType{
		name: "User-Password",
		avt:  AvtString,
		min:  16,
		max:  128,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
		},
	},
	ChapPassword: &attrType{
		name: "CHAP-Password",
		avt:  AvtString,
		min:  17,
		max:  17,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
		},
	},
	NasIpAddress: &attrType{
		name: "NAS-IP-Address",
		avt:  AvtAddress,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	NasPort: &attrType{
		name: "NAS-Port",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	ServiceType: &attrType{
		name: "Service-Type",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccessAccept:       AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
			DisconnectNak:		AttrTableZeroOne,
		},
	},
	FramedIpAddress: &attrType{
		name: "Framed-IP-Address",
		avt:  AvtAddress,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccessAccept:       AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	ReplyMessage: &attrType{
		name: "Reply-Message",
		avt:  AvtText,
		table: [PktCodeEnd]EAttrTableValue{
			AccessAccept:       AttrTableZeroMore,
			AccessReject:       AttrTableZeroMore,
			AccessChallenge:    AttrTableZeroMore,
			DisconnectRequest:	AttrTableZeroMore,
		},
	},
	Class: &attrType{
		name: "Class",
		avt:  AvtString,
		table: [PktCodeEnd]EAttrTableValue{
			AccessAccept:       AttrTableZeroMore,
			AccountingRequest:  AttrTableZeroMore,
			DisconnectRequest:	AttrTableZeroMore,
		},
	},
	SessionTimeout: &attrType{
		name: "Session-Timeout",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccessAccept:       AttrTableZeroOne,
			AccessChallenge:    AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
		},
	},
	IdleTimeout: &attrType{
		name: "Idle-Timeout",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccessAccept:       AttrTableZeroOne,
			AccessChallenge:    AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
		},
	},
	CalledStationId: &attrType{
		name: "Called-Station-Id",
		avt:  AvtString,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	CallingStationId: &attrType{
		name: "Calling-Station-Id",
		avt:  AvtString,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	NasIdentifier: &attrType{
		name: "NAS-Identifier",
		avt:  AvtString,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	EventTimestamp: &attrType{
		name: "Event-Timestamp",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
			DisconnectAck:		AttrTableZeroOne,
			DisconnectNak:		AttrTableZeroOne,
		},
	},
	ChapChallenge: &attrType{
		name: "CHAP-Challenge",
		avt:  AvtString,
		min:  5,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
		},
	},
	NasPortType: &attrType{
		name: "NAS-Port-Type",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	NasPortId: &attrType{
		name: "NAS-Port-Id",
		avt:  AvtText,
		table: [PktCodeEnd]EAttrTableValue{
			AccessRequest:      AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},

	AcctStatusType: &attrType{
		name: "Acct-Status-Type",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableOne,
		},
	},
	AcctDelayTime: &attrType{
		name: "Acct-Delay-Time",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableZeroOne,
		},
	},
	AcctInputOctets: &attrType{
		name: "Acct-Input-Octets",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableZeroOne,
		},
	},
	AcctOutputOctets: &attrType{
		name: "Acct-Output-Octets",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableZeroOne,
		},
	},
	AcctSessionId: &attrType{
		name: "Acct-Session-Id",
		avt:  AvtText,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableOne,
			DisconnectRequest:	AttrTableZeroOne,
		},
	},
	AcctSessionTime: &attrType{
		name: "Acct-Session-Time",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableZeroOne,
		},
	},
	AcctTerminateCause: &attrType{
		name: "Acct-Terminate-Cause",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableZeroOne,
			DisconnectRequest:	AttrTableZeroOne,
			DisconnectAck:		AttrTableZeroOne,
		},
	},
	AcctInputGigawords: &attrType{
		name: "Acct-Input-Gigawords",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableZeroOne,
		},
	},
	AcctOutputGigawords: &attrType{
		name: "Acct-Output-Gigawords",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			AccountingRequest:  AttrTableZeroOne,
		},
	},
	ErrorCause: &attrType{
		name: "Error-Cause",
		avt:  AvtInteger,
		table: [PktCodeEnd]EAttrTableValue{
			DisconnectAck:		AttrTableZeroMore,
			DisconnectNak:		AttrTableZeroMore,
		},
	},
}

func initAttrType() {
	for idx := AttrTypeBegin; idx < AttrTypeEnd; idx++ {
		v := attrTypeBind[idx]
		if nil == v {
			continue
		}
		v.idx = idx

		attrTypeMap[v.name] = v

		switch v.avt {
		case AvtText:fallthrough
		case AvtString:
			if 0 == v.min {
				v.min = 1
			}
			if 0 == v.max {
				v.max = AttrLengthMax - 2
			}
		case AvtInteger:fallthrough
		case AvtAddress:fallthrough
		case AvtTime:
			if 0 == v.min {
				v.min = 4
			}
			if 0 == v.max {
				v.max = 4
			}
		}
	}
}
