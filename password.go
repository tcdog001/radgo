package radgo

import (
	. "asdf"
	"crypto/md5"
)

const (
	authPap 	= 0
	authChap	= 1
	
	papBlockSize 		= 16
	chapChallengeSize 	= 16
)

type papPassword []byte

func (me papPassword) Count() int {
	return (len(me) + papBlockSize - 1) / papBlockSize
}

func (me papPassword) isAlign() bool {
	return len(me) == me.Count()*papBlockSize
}

func (me papPassword) Block(idx int) []byte {
	count := me.Count()
	if idx < 0 || idx >= count {
		Log.Error("bad user password block index(%d), should [%d, %d)",
			idx, 0, count)

		return nil
	} else if idx < count-1 || me.isAlign() {
		return me[idx*papBlockSize : (idx+1)*papBlockSize]
	}

	tail := [papBlockSize]byte{}
	copy(tail[:], me[idx*papBlockSize:])

	return tail[:]
}

func enPapPassword(auth PktAuth, secret []byte, password []byte) []byte {
	me := papPassword(password)
	
	crypt := []byte(nil)
	count := me.Count()
	var b = make([][papBlockSize]byte, count)
	var c = make([][papBlockSize]byte, count)

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


// Chap-Challenge = user-mac + dev-mac + random + random + random + id
func newChapChallenge(user, dev Mac) []byte {
	c := [chapChallengeSize]byte{}
	
	copy(c[0:], user)
	copy(c[6:], dev)
	c[12] = byte(rSeed.Uint32())
	c[13] = byte(rSeed.Uint32())
	c[14] = byte(rSeed.Uint32())
	c[15] = byte(rSeed.Uint32()) // id

	return c[:]
}

// Chap-Password = id + md5(id|password|challenge)
func enChapPassword(password []byte, challenge []byte) []byte {
	id := []byte{challenge[15]}
	
	md5 := md5.New()
	
	md5.Write(id[:])
	md5.Write(password)
	md5.Write(challenge)
	
	return append(id[:], md5.Sum(nil)...)
}