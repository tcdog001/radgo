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
	if nil==me {
		return false
	}
	
	return me.Type.IsGoodLength(me.Len) && me.Type.IsGoodValue(me.Number)
}

func (me *Attr) GetString() []byte {
	if nil==me {
		return nil
	}
	
	return me.Value[:me.Len-2]
}
	
func (me *Attr) SetString(Value []byte) error {
	Type := me.Type
	
	if nil==me {
		return ErrNilObj
	}
	
	// check value type
	if !Type.ValueType().IsString() {
		log.Error("attr %s value is not string", Type.ToString())
		return Error
	}
	
	// check length
	Len := 2 + byte(len(Value))
	if !Type.IsGoodLength(Len) {
		return Error
	}
	
	copy(me.Value[:], Value)
	me.Len = Len
	
	return nil
}

func (me *Attr) SetNumber(Value uint32) error {
	if nil==me {
		return ErrNilObj
	}
	
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
		return ErrNilObj
	}
	
	if !me.IsGood() {
		return Error
	} else if me.Len > byte(len(bin)) {
		log.Error("attr(%s) Len(%d) < bin Len(%d)",
			me.Type.ToString(),
			me.Len,
			len(bin))
	}
	
	bin[0] = byte(me.Type)
	bin[1] = me.Len
	
	if me.Type.ValueType().IsNumber() {
		binary.BigEndian.PutUint32(bin[2:], me.Number)
		
		log.Info("write attr(%s) Len(%d) Number(%d)",
			me.Type.ToString(),
			me.Len,
			me.Number)
	} else {
		copy(bin[2:], me.GetString())
		
		log.Info("write attr(%s) Len(%d) String(%d)",
			me.Type.ToString(),
			me.Len,
			me.Value[:me.Len-2])
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
		return ErrNilObj
	}
	
	Type := EAttrType(bin[0])
	Len  := bin[1]
	if !Type.IsGoodLength(Len) {
		return Error
	} else if Len > byte(len(bin)) {
		log.Error("bin attr(%s) Len(%d) < bin Len(%d)",
			Type.ToString(),
			Len,
			len(bin))
		return Error
	}
	
	me.Type = Type
	me.Len  = Len
	
	if Type.ValueType().IsNumber() {
		me.Number = binary.BigEndian.Uint32(bin[2:])
		
		log.Info("read attr(%s) Len(%d) Number(%d)",
			me.Type.ToString(),
			me.Len,
			me.Number)
	} else {
		copy(me.Value[:], bin[2:Len-2])
		
		log.Info("read attr(%s) Len(%d) String(%d)",
			me.Type.ToString(),
			me.Len,
			me.Value[:me.Len-2])
	}
	
	if !Type.IsGoodValue(me.Number) {
		return Error
	}
	
	return nil
}

