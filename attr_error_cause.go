package radgo

import (
	. "asdf"
)

type EAecValue uint32

func (me EAecValue) Tag() string {
	return "Error-Cause"
}

func (me EAecValue) Begin() int {
	return int(aecBegin)
}

func (me EAecValue) End() int {
	return int(aecEnd)
}

func (me EAecValue) Int() int {
	return int(me)
}

func (me EAecValue) IsGood() bool {
	if !IsGoodEnum(me) {
		Log.Error("bad attr(%s) value(%d)", me.Tag(), me)

		return false
	} else if 0 == len(aecBind[me]) {
		Log.Error("no support attr(%s) value(%d)", me.Tag(), me)

		return false
	}

	return true
}

func (me EAecValue) ToString() string {
	var b EnumBinding = aecBind[:]

	return b.EntryShow(me)
}

func (me *EAecValue) FromString(Name string) error {
	if e, ok := aecMap[Name]; ok {
		*me = e

		return nil
	}

	return ErrNoFound
}

const (
	aecBegin	EAecValue = 201
	
	AecResidualSessionContextRemoved	EAecValue = 201
	AecInvalidEapPacket					EAecValue = 202
	AecUnsupportedAttribute 			EAecValue = 401
	AecMissingAttribute					EAecValue = 402
	AecNasIdentificationMismatch		EAecValue = 403
	AecInvalidRequest					EAecValue = 404
	AecUnsupportedService				EAecValue = 405
	AecUnsupportedExtension				EAecValue = 406
	AecAdministrativelyProhibited		EAecValue = 501
	AecRequestNotRoutable				EAecValue = 502
	AecSessionContextNotFound			EAecValue = 503
	AecSessionContextNotRemovable		EAecValue = 504
	AecOtherProxyProcessingError		EAecValue = 505
	AecResourcesUnavailable				EAecValue = 506
	AecRequestInitiated					EAecValue = 507
	
	aecEnd	EAecValue = 508
)

var aecBind = [aecEnd]string{
	AecResidualSessionContextRemoved:	"Residual Session Context Removed",
	AecInvalidEapPacket:				"Invalid EAP Packet",
	AecUnsupportedAttribute:			"Unsupported Attribute",
	AecMissingAttribute:				"Missing Attribute",
	AecNasIdentificationMismatch:		"NAS Identification Mismatch",
	AecInvalidRequest:					"Invalid Request",
	AecUnsupportedService:				"Unsupported Service",
	AecUnsupportedExtension:			"Unsupported Extension",
	AecAdministrativelyProhibited:		"Administratively Prohibited",
	AecRequestNotRoutable:				"Request Not Routable",
	AecSessionContextNotFound:			"Session Context Not Found",
	AecSessionContextNotRemovable:		"Session Context Not Removable",
	AecOtherProxyProcessingError:		"Other Proxy Processing Error",
	AecResourcesUnavailable:			"Resources Unavailable",
	AecRequestInitiated:				"Request Initiated",
}

var aecMap = map[string]EAecValue{}

func initAec() {
	for i := aecBegin; i < aecEnd; i++ {
		aecMap[aecBind[i]] = i
	}
}

