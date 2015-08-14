package radgo

import (
	. "asdf"
	"encoding/binary"
	"math/rand"
	"time"
)

const AcctSessionIdLength = 20

var sessionSeq uint32 = 0

// Acct-Session-Id Private Format
type SessionId struct {
	Mac  [6]byte // 6, user mac
	Rand uint16  // 2
	Seq  uint32  // 4
	Unix uint64  // 8
}

func (me *SessionId) Init(mac Mac) {
	sessionSeq += 1

	copy(me.Mac[:], mac)
	me.Rand = uint16(rand.Uint32())
	me.Seq = sessionSeq // needn't lock
	me.Unix = uint64(time.Now().Unix())
}

func (me *SessionId) checkLengh(bin []byte) error {
	Len := len(bin)

	if Len < AcctSessionIdLength {
		Log.Error("SessionId bin length is %d, must >= %d",
			Len,
			AcctSessionIdLength)

		return Error
	}

	return nil
}

func (me *SessionId) ToBinary(bin []byte) error {
	if err := me.checkLengh(bin); nil != err {
		return err
	}

	copy(bin, me.Mac[:])
	binary.BigEndian.PutUint16(bin[6:], me.Rand)
	binary.BigEndian.PutUint32(bin[8:], me.Seq)
	binary.BigEndian.PutUint64(bin[12:], me.Unix)

	return nil
}

func (me *SessionId) FromBinary(bin []byte) error {
	if err := me.checkLengh(bin); nil != err {
		return err
	}

	copy(me.Mac[:], bin)
	me.Rand = binary.BigEndian.Uint16(bin[6:])
	me.Seq = binary.BigEndian.Uint32(bin[8:])
	me.Unix = binary.BigEndian.Uint64(bin[12:])

	return nil
}
