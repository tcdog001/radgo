package radgo

import (
	. "asdf"
)

type EAstValue uint32

func (me EAstValue) Tag() string {
	return "Service-Type"
}

func (me EAstValue) Begin() int {
	return int(astBegin)
}

func (me EAstValue) End() int {
	return int(astEnd)
}

func (me EAstValue) Int() int {
	return int(me)
}

func (me EAstValue) IsGood() bool {
	return IsGoodEnum(me) && 
		len(astBind)==me.End() && 
		len(astBind[me]) > 0
}

func (me EAstValue) ToString() string {
	var b EnumBinding = astBind[:]

	return b.EntryShow(me)
}

const (
	astBegin				EAstValue = 1
	
	AstLogin 				EAstValue = 1
	AstFramed				EAstValue = 2
	AstCallbackLogin		EAstValue = 3
	AstCallbackFramed		EAstValue = 4
	AstOutbound				EAstValue = 5
	AstAdministrative		EAstValue = 6
	AstNasPrompt			EAstValue = 7
	AstAuthenticateOnly		EAstValue = 8
	AstCallbackNasPrompt	EAstValue = 9
	AstCallCheck			EAstValue = 10
	AstCallbackAdministrative 	EAstValue = 11

	astEnd 					EAstValue = 12
)

var astBind = [astEnd]string{
	AstLogin:			"Login",
	AstFramed:			"Framed",
	AstCallbackLogin:	"Callback Login",
	AstCallbackFramed:	"Callback Framed",
	AstOutbound:		"Outbound",
	AstAdministrative:	"Administrative",
	AstNasPrompt:		"NAS Prompt",
	AstAuthenticateOnly:		"Authenticate Only",
	AstCallbackNasPrompt:		"Callback NAS Prompt",
	AstCallCheck:				"Call Check",
	AstCallbackAdministrative:	"Callback Administrative",
}

