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
	if !IsGoodEnum(me) {
		Log.Error("bad attr(%s) value(%d)", me.Tag(), me)

		return false
	} else if 0 == len(atcBind[me]) {
		Log.Error("no support attr(%s) value(%d)", me.Tag(), me)

		return false
	}

	return true
}

func (me EAtcValue) ToString() string {
	var b EnumBinding = atcBind[:]

	return b.EntryShow(me)
}

func (me *EAtcValue) FromString(Name string) error {
	if e, ok := atcMap[Name]; ok {
		*me = e

		return nil
	}

	return ErrNoFound
}

const (
	atcBegin EAtcValue = 1

	AtcUserRequest        EAtcValue = 1
	AtcLostCarrier        EAtcValue = 2
	AtcLostService        EAtcValue = 3
	AtcIdleTimeout        EAtcValue = 4
	AtcSessionTimeout     EAtcValue = 5
	AtcAdminReset         EAtcValue = 6
	AtcAdminReboot        EAtcValue = 7
	AtcPortError          EAtcValue = 8
	AtcNasError           EAtcValue = 9
	AtcNasRequest         EAtcValue = 10
	AtcNasReboot          EAtcValue = 11
	AtcPortUnneeded       EAtcValue = 12
	AtcPortPreempted      EAtcValue = 13
	AtcPortSuspended      EAtcValue = 14
	AtcServiceUnavailable EAtcValue = 15
	AtcCallback           EAtcValue = 16
	AtcUserError          EAtcValue = 17
	AtcHostRequest        EAtcValue = 18

	atcEnd EAtcValue = 19
)

var atcBind = [atcEnd]string{
	AtcUserRequest:        "User Request",
	AtcLostCarrier:        "Lost Carrier",
	AtcLostService:        "Lost Service",
	AtcIdleTimeout:        "Idle Timeout",
	AtcSessionTimeout:     "Session Timeout",
	AtcAdminReset:         "Admin Reset",
	AtcAdminReboot:        "Admin Reboot",
	AtcPortError:          "Port Error",
	AtcNasError:           "NAS Error",
	AtcNasRequest:         "NAS Request",
	AtcNasReboot:          "NAS Reboot",
	AtcPortUnneeded:       "Port Unneeded",
	AtcPortPreempted:      "Port Preempted",
	AtcPortSuspended:      "Port Suspended",
	AtcServiceUnavailable: "Service Unavailable",
	AtcCallback:           "Callback",
	AtcUserError:          "User Error",
	AtcHostRequest:        "Host Request",
}

var atcMap = map[string]EAtcValue{}

func initAtc() {
	for i := atcBegin; i < atcEnd; i++ {
		atcMap[atcBind[i]] = i
	}
}

// device deauth resaon
type DeauthReason uint32

func (me DeauthReason) Tag() string {
	return "Dev-Deauth-Reason"
}

func (me DeauthReason) Begin() int {
	return int(DeauthReasonBegin)
}

func (me DeauthReason) End() int {
	return int(DeauthReasonEnd)
}

func (me DeauthReason) Int() int {
	return int(me)
}

func (me DeauthReason) IsGood() bool {
	if !IsGoodEnum(me) {
		Log.Error("bad %s(%d)", me.Tag(), me)

		return false
	} else if 0 == len(drBind[me]) {
		Log.Error("no support %s(%d)", me.Tag(), me)

		return false
	}

	return true
}

func (me DeauthReason) ToString() string {
	var b EnumBinding = drBind[:]

	return b.EntryShow(me)
}

func (me *DeauthReason) FromString(Name string) error {
	if e, ok := drMap[Name]; ok {
		*me = e

		return nil
	}

	return ErrNoFound
}

func (me DeauthReason) TerminateCause() uint32 {
	return uint32(ressonToCause[me])
}

const (
	DeauthReasonBegin DeauthReason = 0

	DeauthReasonNone       DeauthReason = 0
	DeauthReasonAuto       DeauthReason = 1
	DeauthReasonOnlineTime DeauthReason = 2
	DeauthReasonFlowLimit  DeauthReason = 3
	DeauthReasonAdmin      DeauthReason = 4
	DeauthReasonAging      DeauthReason = 5
	DeauthReasonInitiative DeauthReason = 6

	DeauthReasonEnd DeauthReason = 7
)

var drBind = [DeauthReasonEnd]string{
	DeauthReasonNone:       "None",
	DeauthReasonAuto:       "Auto",
	DeauthReasonOnlineTime: "OnlineTimeOut",
	DeauthReasonFlowLimit:  "FlowLimit",
	DeauthReasonAdmin:      "Admin",
	DeauthReasonAging:      "IdleTimeOut",
	DeauthReasonInitiative: "Initiative",
}

var ressonToCause = [DeauthReasonEnd]EAtcValue{
	DeauthReasonNone:       AtcUserError,
	DeauthReasonAuto:       AtcUserError,
	DeauthReasonOnlineTime: AtcSessionTimeout,
	DeauthReasonFlowLimit:  AtcLostService, // need define new cause ???
	DeauthReasonAdmin:      AtcAdminReset,
	DeauthReasonAging:      AtcIdleTimeout,
	DeauthReasonInitiative: AtcUserRequest,
}

var drMap = map[string]DeauthReason{}

func initDeauthReason() {
	for i := DeauthReasonBegin; i < DeauthReasonEnd; i++ {
		drMap[drBind[i]] = i
	}
}
