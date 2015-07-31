package radgo

import (
	. "asdf"
)

type EAnptValue uint32

func (me EAnptValue) Tag() string {
	return "NAS-Port-Type"
}

func (me EAnptValue) Begin() int {
	return int(anptBegin)
}

func (me EAnptValue) End() int {
	return int(anptEnd)
}

func (me EAnptValue) Int() int {
	return int(me)
}

func (me EAnptValue) IsGood() bool {
	if !IsGoodEnum(me) {
		log.Info("bad %s %d", me.Tag(), me)
		
		return false
	} else if 0==len(anptBind[me]) {
		log.Info("no support %s %d", me.Tag(), me)
		
		return false
	}
	
	return true
}

func (me EAnptValue) ToString() string {
	var b EnumBinding = anptBind[:]

	return b.EntryShow(me)
}

const (
	anptBegin				EAnptValue = 0
	
	AnptAsync				EAnptValue = 0
	AnptSync				EAnptValue = 1
	AnptIsdnSync			EAnptValue = 2
	AnptIsdnAsyncV1210		EAnptValue = 3
	AnptIsdnAsyncV1110		EAnptValue = 4
	AnptVirtual				EAnptValue = 5
	AnptPiafs 				EAnptValue = 6 // PIAFS
	AnptHdlcClearChannel	EAnptValue = 7 // HDLC Clear Channel
	AnptX25					EAnptValue = 8
	AnptX75					EAnptValue = 9
	AnptG3Fax 				EAnptValue = 10// G.3 Fax
	AnptSdsl 				EAnptValue = 11// SDSL - Symmetric DSL
	AnptAdslCap				EAnptValue = 12// ADSL-CAP - Asymmetric DSL, Carrierless Amplitude Phase Modulation
	AnptAdslDmt				EAnptValue = 13// ADSL-DMT - Asymmetric DSL, Discrete Multi-Tone
	AnptIdsl 				EAnptValue = 14// IDSL - ISDN Digital Subscriber Line
	AnptEthernet			EAnptValue = 15
	AnptXdsl				EAnptValue = 16// xDSL - Digital Subscriber Line of unknown type
	AnptCable				EAnptValue = 17
	AnptWirelessOther		EAnptValue = 18// Wireless - Other
	AnptIeee80211			EAnptValue = 19//Wireless - IEEE 802.11

	anptEnd 				EAnptValue = 20
)

var anptBind = [anptEnd]string{
	AnptAsync:				"Async",
	AnptSync:				"Sync",
	AnptIsdnSync:			"ISDN Sync",
	AnptIsdnAsyncV1210:		"ISDN Async V.120",
	AnptIsdnAsyncV1110:		"ISDN Async V.110",
	AnptVirtual:			"Virtual",
	AnptPiafs:				"PIAFS",
	AnptHdlcClearChannel:	"HDLC Clear Channel",
	AnptX25:				"X.25",
	AnptX75:				"X.75",
	AnptG3Fax:				"G.3 Fax",
	AnptSdsl:				"SDSL",
	AnptAdslCap:			"ADSL-CAP",
	AnptAdslDmt:			"ADSL-DMT",
	AnptIdsl:				"IDSL",
	AnptEthernet:			"Ethernet",
	AnptXdsl:				"xDSL",
	AnptCable:				"Cable",
	AnptWirelessOther:		"Wireless - Other",
	AnptIeee80211:			"Wireless - IEEE 802.11",
}

