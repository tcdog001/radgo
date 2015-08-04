package radgo

import (
	. "asdf"
	"encoding/binary"
)

const (
	AttrLengthMax = 255
)


type AttrBinary []byte

func (me AttrBinary) Bin() []byte {
	return []byte(me)
}

func (me AttrBinary) Type() EAttrType {
	return EAttrType(me[0])
}

func (me AttrBinary) Len() byte {
	return me[1]
}

func (me AttrBinary) Value() []byte {
	return me[2:me.Len()]
}

func (me AttrBinary) Next() AttrBinary {
	Len := byte(len(me))
	
	if Len <= me.Len() {
		return nil
	}
	
	return me[me.Len():]
}

func (me AttrBinary) Number() uint32 {
	return binary.BigEndian.Uint32(me.Value())	
}

func (me AttrBinary) SetNumber(Type EAttrType, Number uint32) {
	me[0] = byte(Type)
	me[1] = 6
	
	binary.BigEndian.PutUint32(me.Value(), Number)
}

func (me AttrBinary) SetString(Type EAttrType, Value []byte) {
	me[0] = byte(Type)
	me[1] = 2 + byte(len(Value))
	
	copy(me.Value(), Value)
}

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
	
	return me.Type.IsGoodLength(me.Len) && me.Type.IsGoodNumber(me.Number)
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
	
	me.Len = Len
	copy(me.GetString(), Value)

	log.Info("set attr(%s) Len(%d) String(%s)",
		Type.ToString(),
		Len,
		me.GetString())
	
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
	if !me.Type.IsGoodNumber(Value) {
		return Error
	}
	
	me.Len = 6
	me.Number = Value

	log.Info("set attr(%s) Number(%d)",
		me.Type.ToString(),
		me.Number)
	
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
	
	Len := byte(len(bin))
	
	if !me.IsGood() {
		return ErrBadObj
	} 
	
	if me.Len > Len {
		log.Error("attr(%s) Len(%d) < bin Len(%d)",
			me.Type.ToString(),
			me.Len,
			Len)
		
		return ErrTooShortBuffer
	}
	
	ab := AttrBinary(bin)
	
	if me.Type.ValueType().IsNumber() {
		ab.SetNumber(me.Type, me.Number)
	
		log.Info("write attr(%s) Number(%d)",
			me.Type.ToString(),
			me.Number)
	} else {
		ab.SetString(me.Type, me.GetString())

		log.Info("write attr(%s) Len(%d) String(%s)",
			me.Type.ToString(),
			me.Len,
			me.GetString())
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
	
	ab := AttrBinary(bin)
	
	Type := ab.Type()
	Len  := ab.Len()
	if !Type.IsGoodLength(Len) {
		return ErrBadObj
	} 
	
	if Len > byte(len(bin)) {
		log.Error("bin attr(%s) Len(%d) < bin Len(%d)",
			Type.ToString(),
			Len,
			len(bin))
		return Error
	}
	
	me.Type = Type
	me.Len  = Len
	
	if Type.ValueType().IsNumber() {
		me.Number = ab.Number()
		
		if !Type.IsGoodNumber(me.Number) {
			return Error
		}
		
		log.Info("read attr(%s) Len(%d) Number(%d)",
			me.Type.ToString(),
			me.Len,
			me.Number)
	} else {
		copy(me.GetString(), ab.Value())
		
		log.Info("read attr(%s) Len(%d) String(%s)",
			me.Type.ToString(),
			me.Len,
			me.GetString())
	}
	
	return nil
}

