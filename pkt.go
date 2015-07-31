package radgo

import (
	. "asdf"
	"fmt"
	"encoding/binary"
)

const (
	PktHdrSize 		= 20
	
	PktLengthMin 	= 20
	PktLengthMax 	= 4096
)

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

func isGoodPktLength(Len uint16) bool{
	if Len < PktLengthMin || Len > PktLengthMax {
		log.Error("pkt length is %d, should [%d, %d]",
			Len,
			PktLengthMin,
			PktLengthMax)
		
		return false
	}
	
	return true
}

func (me *Header) IsGood() bool {
	return isGoodPktLength(me.Len) && me.Code.IsGood()
}

func (me *Header) ToBinary(bin []byte) error {
	if nil==me {
		log.Error("empty packet")
		
		return ErrNilObj
	}
	
	if !me.IsGood() {
		return ErrBadObj
	}
	
	if len(bin) < int(me.Len) {
		log.Error("packet==>bin, not enough space")
		
		return ErrNoSpace
	}
	
	bin[0] = byte(me.Code)
	bin[1] = me.Id
	binary.BigEndian.PutUint16(bin[2:], me.Len)
	copy(bin[4:], me.Auth[:])
	
	return nil
}

func (me *Header) FromBinary(bin []byte) error {
	if nil==me {
		log.Error("empty packet")
		
		return ErrNilObj
	}
	
	if len(bin) < PktHdrSize {
		log.Error("packet buffer is tool short")
		
		return ErrTooShortBuffer
	}
	
	code := EPktCode(bin[0])
	if !code.IsGood() {
		return Error
	}
	
	Len := binary.BigEndian.Uint16(bin[2:])
	if !isGoodPktLength(Len) {
		return ErrBadPktLen
	}
	
	if len(bin) < int(Len) {
		return ErrPktLenNoMatchBufferLen
	}
	
	me.Code = code
	me.Id 	= bin[1]
	me.Len 	= Len
	copy(me.Auth[:], bin[4:4+AuthSize])
	
	return nil
}

type Packet struct {
	Header
	Attrs [AttrTypeEnd]*Attr
}

func (me *Packet) Init() {
	me.Len = PktHdrSize
}

func (me *Packet) attr(Type EAttrType) (*Attr, bool) {
	if nil!=me.Attrs[Type] {
		return me.Attrs[Type], false
	}
	
	me.Attrs[Type] = new(Attr)
	me.Attrs[Type].Type = Type
	
	fmt.Println("create new attr", Type.ToString())
	
	return me.Attrs[Type], true
}

func (me *Packet) SetAttrNumber(Type EAttrType, Value uint32) error {
	if !me.Code.IsMatch(Type) {
		return Error
	}
	
	attr, create := me.attr(Type)
	err := attr.SetNumber(Value)
	if create && nil==err {
		me.Len += uint16(attr.Len)
	}
	
	return err
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
	
	attr, create := me.attr(Type)
	err := attr.SetString(Value)
	if create && nil==err {
		me.Len += uint16(attr.Len)
	}
	
	return err
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
		attr := me.Attrs[i]
		
		// if the code is must, but attr is empty
		if me.Code.IsMust(i) && !attr.IsGood() {
			log.Error("attr type %s must match code %s, but attr is bad",
				i.ToString(),
				me.Code.ToString())
				
			return Error
		}
	}
	
	return nil
}

/*
func (me *Packet) CalcLength() uint16 {
	Len := uint16(PktHdrSize)
	
	for i:=AttrTypeBegin; i<AttrTypeEnd; i++ {
		if attr := me.Attrs[i]; nil!=attr {
			Len += uint16(attr.Len)
		}
	}
	
	return Len
}
*/

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
		attr := me.Attrs[i]
		if nil==attr {
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
	
	Len := uint16(len(bin))
	if !isGoodPktLength(Len) {
		return Error
	}
	
	// bin==>hdr
	if err:=me.Header.FromBinary(bin); nil!=err {
		return err
	}
	bin = bin[PktHdrSize:]

	// bin==>attr
	for len(bin) > 0 {
		attr := me.Attrs[bin[0]]
		
		if err:=attr.FromBinary(bin); nil!=err {
			return err
		}
		
		if !me.Code.IsMatch(attr.Type) {
			return Error
		}
		
		bin = bin[attr.Len:]
		
		me.Len += uint16(attr.Len)
	}
	
	if err:=me.CheckMust(); nil!=err {
		return err
	}
	
	return nil
}

func (me *Packet) Policy() *Policy {
	if AccessAccept!=me.Code {
		return nil
	}
	
	var attr *Attr
	policy := &Policy{}
	
	attr = me.Attrs[SessionTimeout]
	if attr.IsGood() {
		policy.OnlineTime = attr.Number
	}
	
	attr = me.Attrs[IdleTimeout]
	if attr.IsGood() {
		policy.IdleTimeout = attr.Number
	}
	
	attr = me.Attrs[Class]
	if attr.IsGood() {
		class := AttrClass(attr.GetString())
		
		policy.UpRateMax 	= class.UpRateMax()
		policy.UpRateAvg 	= class.UpRateAvg()
		policy.DownRateMax 	= class.DownRateMax()
		policy.DownRateAvg 	= class.DownRateAvg()
	}
	
	return policy
}