package radgo

import (
	. "asdf"
	"encoding/binary"
)

const (
	PktLengthMin 	= 20
	PktLengthMax 	= 4096
)

const PktHdrSize 	= 20

var pktId byte = 0

func PktId() byte {
	pktId += 1
	
	return pktId
}

type PktLength uint16

type Header struct {
	Code 	EPktCode
	Id   	byte
	Len 	uint16
	Auth 	[16]byte
}

func isGoodPktLength(length uint16) bool{
	return length >= PktLengthMin && length <= PktLengthMax
}

func (me *Header) IsGood() bool {
	return isGoodPktLength(me.Len) && me.Code.IsGood()
}

func (me *Header) ToBinary(bin []byte) error {
	if nil==me {
		return Error
	}
	
	if !me.IsGood() {
		return Error
	}
	
	if len(bin) < int(me.Len) {
		return Error
	}
	
	bin[0] = byte(me.Code)
	bin[1] = me.Id
	binary.BigEndian.PutUint16(bin[2:], me.Len)
	copy(bin[4:], me.Auth[:])
	
	return nil
}

func (me *Header) FromBinary(bin []byte) error {
	if nil==me {
		return Error
	}
	
	if len(bin) < PktHdrSize {
		return Error
	}
	
	code := EPktCode(bin[0])
	if !code.IsGood() {
		return Error
	}
	
	Len := binary.BigEndian.Uint16(bin[2:])
	if !isGoodPktLength(Len) {
		return Error
	}
	
	if len(bin) < int(Len) {
		return Error
	}
	
	me.Code = code
	me.Id 	= bin[1]
	me.Len 	= Len
	copy(me.Auth[:], bin[4:4+AuthSize])
	
	return nil
}

type Packet struct {
	Header
	Attrs [AttrTypeEnd]Attr
}

func (me *Packet) Init() {
	for i:=AttrTypeBegin; i<AttrTypeEnd; i++ {
		me.Attrs[i].Type = i
	}
}

func (me *Packet) SetAttrNumber(Type EAttrType, Value uint32) error {
	if !me.Code.IsMatch(Type) {
		return Error
	}
	
	return (&me.Attrs[Type]).SetNumber(Value)
}

type AttrNumber struct {
	Type 	EAttrType
	Value 	uint32
}

func (me *Packet) SetAttrNumberList(list []AttrNumber) error {
	for _, v := range list {		
		if err := me.SetAttrNumber(v.Type, v.Value); nil!=err {
			return err
		}
	}
	
	return nil
}

func (me *Packet) SetAttrString(Type EAttrType, Value []byte) error {
	if nil==Value {
		return nil // Not Error
	}
	
	if !me.Code.IsMatch(Type) {
		return Error
	}
	
	return (&me.Attrs[Type]).SetString(Value)
}

type AttrString struct {
	Type 	EAttrType
	Value 	[]byte
}

func (me *Packet) SetAttrStringList(list []AttrString) error {
	for _, v := range list {
		if err := me.SetAttrString(v.Type, v.Value); nil!=err {
			return err
		}
	}
	
	return nil
}

func (me *Packet) CheckMust() error {
	for i:=AttrTypeBegin; i<AttrTypeEnd; i++ {
		attr := &me.Attrs[i]
		
		if me.Code.IsMust(i) && !attr.IsGood() {
			return Error
		}
	}
	
	return nil
}

func (me *Packet) ToBinary(bin []byte) error {
	if nil==me {
		return Error
	}
	
	if !me.IsGood() {
		return Error
	}
	
	if err := me.CheckMust(); nil!=err {
		return err
	}
	
	// hdr==>bin
	if err := me.Header.ToBinary(bin); nil!=err {
		return err
	}
	bin = bin[PktHdrSize:]
	
	// attr==>bin
	for i:=AttrTypeBegin; i<AttrTypeEnd; i++ {
		attr := &me.Attrs[i]
		if !attr.IsGood() {
			continue
		}
		
		if err:=attr.ToBinary(bin); nil!=err {
			return err
		}
		bin = bin[attr.Len:]
	}
	
	return nil
}

func (me *Packet) FromBinary(bin []byte) error {
	if nil==me {
		return Error
	}
	
	if !isGoodPktLength(uint16(len(bin))) {
		return Error
	}
	
	// bin==>hdr
	if err:=me.Header.FromBinary(bin); nil!=err {
		return err
	}
	bin = bin[PktHdrSize:]

	// bin==>attr
	for len(bin) > 0 {
		attr := &me.Attrs[bin[0]]
		
		if err:=attr.FromBinary(bin); nil!=err {
			return err
		}
		
		if !me.Code.IsMatch(attr.Type) {
			return Error
		}
		
		bin = bin[attr.Len:]
	}
	
	return me.CheckMust()
}

func (me *Packet) Policy(policy *Policy) {
	var attr *Attr
	
	attr = &me.Attrs[SessionTimeout]
	if attr.IsGood() {
		policy.OnlineTime = attr.Number
	}
	
	attr = &me.Attrs[IdleTimeout]
	if attr.IsGood() {
		policy.IdleTimeout = attr.Number
	}
	
	attr = &me.Attrs[Class]
	if attr.IsGood() {
		policy.FlowLimit = 0
		policy.RateLimit = 0
	}
}