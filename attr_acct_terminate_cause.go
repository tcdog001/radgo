package radgo

import (
	. "asdf"
)

type EAtcValue uint32

func (me EAtcValue) Tag() string {
	return "Acct-Terminate-Cause"
}

func (me EAtcValue) Begin() int {
	return int(atcBegin)
}

func (me EAtcValue) End() int {
	return int(atcEnd)
}

func (me EAtcValue) Int() int {
	return int(me)
}

func (me EAtcValue) IsGood() bool {
	return IsGoodEnum(me) && 
		len(atcBind)==me.End() && 
		len(atcBind[me]) > 0
}

func (me EAtcValue) ToString() string {
	var b EnumBinding = atcBind[:]

	return b.EntryShow(me)
}

const (
	atcBegin 			EAtcValue = 1
	
	AtcUserRequest		EAtcValue = 1
	AtcLostCarrier		EAtcValue = 2
	AtcLostService		EAtcValue = 3
	AtcIdleTimeout		EAtcValue = 4
	AtcSessionTimeout	EAtcValue = 5
	AtcAdminReset		EAtcValue = 6
	AtcAdminReboot		EAtcValue = 7
	AtcPortError		EAtcValue = 8
	AtcNasError			EAtcValue = 9
	AtcNasRequest		EAtcValue = 10
	AtcNasReboot		EAtcValue = 11
	AtcPortUnneeded		EAtcValue = 12
	AtcPortPreempted	EAtcValue = 13
	AtcPortSuspended	EAtcValue = 14
	AtcServiceUnavailable	EAtcValue = 15
	AtcCallback			EAtcValue = 16
	AtcUserError		EAtcValue = 17
	AtcHostRequest		EAtcValue = 18

	atcEnd 				EAtcValue = 19
)

var atcBind = [atcEnd]string{
	AtcUserRequest:		"User Request",
	AtcLostCarrier:		"Lost Carrier",
	AtcLostService:		"Lost Service",
	AtcIdleTimeout:		"Idle Timeout",
	AtcSessionTimeout:	"Session Timeout",
	AtcAdminReset:		"Admin Reset",
	AtcAdminReboot:		"Admin Reboot",
	AtcPortError:		"Port Error",
	AtcNasError:		"NAS Error",
	AtcNasRequest:		"NAS Request",
	AtcNasReboot:		"NAS Reboot",
	AtcPortUnneeded:	"Port Unneeded",
	AtcPortPreempted:	"Port Preempted",
	AtcPortSuspended:	"Port Suspended",
	AtcServiceUnavailable:"Service Unavailable",
	AtcCallback:		"Callback",
	AtcUserError:		"User Error",
	AtcHostRequest:		"Host Request",
}
