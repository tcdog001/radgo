package radgo

import (
	. "asdf"
	"crypto/md5"
)

const passBlockSize = 16

type AttrUserPassword []byte

func (me AttrUserPassword) Count() int {
	return (len(me) + passBlockSize - 1) / passBlockSize
}

func (me AttrUserPassword) isAlign() bool {
	return len(me) == me.Count()*passBlockSize
}

func (me AttrUserPassword) Block(idx int) []byte {
	count := me.Count()
	if idx < 0 || idx >= count {
		Log.Error("bad user password block index(%d), should [%d, %d)",
			idx, 0, count)

		return nil
	} else if idx < count-1 || me.isAlign() {
		return me[idx*passBlockSize : (idx+1)*passBlockSize]
	}

	tail := [passBlockSize]byte{}
	copy(tail[:], me[idx*passBlockSize:])

	return tail[:]
}

func (me AttrUserPassword) Encrypt(auth PktAuth, secret []byte) []byte {
	crypt := []byte(nil)
	count := me.Count()
	var b = make([][passBlockSize]byte, count)
	var c = make([][passBlockSize]byte, count)

	for i := 0; i < count; i++ {
		md5 := md5.New()

		md5.Write(secret)
		md5.Write(auth)

		copy(b[i][:], md5.Sum(nil))
		xor(c[i][:], me.Block(i), b[i][:])
		crypt = append(crypt, c[i][:]...)
	}

	return crypt
}

func xor(r []byte, a []byte, b []byte) {
	Len := len(a)
	for i := 0; i < Len; i++ {
		r[i] = a[i] ^ b[i]
	}
}
