package radgo

import (
	. "asdf"
	"encoding/binary"
)

const (
	AttrLengthMax = 255
)

type Attr struct {
	Type 	EAttrType
	Len 	byte
	Value 	[AttrLengthMax-2]byte
	Number 	uint32
}

// me is full
func (me *Attr) IsGood() bool {
	return me.Type.IsGoodLength(me.Len) && me.Type.IsGoodValue(me.Number)
}

func (me *Attr) SetString(Value []byte) error {
	Type := me.Type
	
	// check value type
	if !Type.ValueType().IsString() {
		return Error
	}
	
	// check length
	Len := byte(len(Value))
	if !Type.IsGoodLength(Len) {
		return Error
	}
	
	copy(me.Value[:], Value)
	me.Len = Len
	
	return nil
}

func (me *Attr) SetNumber(Value uint32) error {
	// check value type
	if !me.Type.ValueType().IsNumber() {
		return Error
	}
	
	// check value
	if !me.Type.IsGoodValue(Value) {
		return Error
	}
	
	me.Number = Value
	me.Len = 6
	
	return nil
}


// before call
// 1. me is full
//		me.Type is set
//  	me.Len is set
//		me.Value/U32 is set
// 2. bin is empty
//
// after call
// 1. me put into bin
func (me *Attr) ToBinary(bin []byte) error {
	if nil==me {
		return Error
	}
	
	if !me.IsGood() || me.Len > byte(len(bin)) {
		return Error
	}
	
	bin[0] = byte(me.Type)
	bin[1] = me.Len
	
	if me.Type.ValueType().IsNumber() {
		binary.BigEndian.PutUint32(bin[2:], me.Number)
	} else {
		copy(bin[2:], me.Value[:me.Len-2])
	}
	
	return nil
}


// before call
// 1. *me is empty
// 2. bin is full
//
// after call
// 1. *me is full
func (me *Attr) FromBinary(bin []byte) error {
	if nil==me {
		return Error
	}
	
	Type := EAttrType(bin[0])
	Len  := bin[1]
	if !Type.IsGoodLength(Len) || Len > byte(len(bin)) {
		return Error
	}
	
	me.Type = Type
	me.Len  = Len
	
	if Type.ValueType().IsNumber() {
		me.Number = binary.BigEndian.Uint32(bin[2:])
	} else {
		copy(me.Value[:], bin[2:Len-2])
	}
	
	if !Type.IsGoodValue(me.Number) {
		return Error
	}
	
	return nil
}

