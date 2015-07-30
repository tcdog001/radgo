package radgo

import (
	. "asdf"
	"crypto/md5"
	"encoding/binary"
	"time"
)

const AuthSize = 16

var pktAuthSeq uint16 = 0

func authSeq() uint16 {
	pktAuthSeq += 1
	
	return pktAuthSeq
}

// Authenticator Private Format in Access-Request/Accept/Reject
type privateAuth struct {
	mac [6]byte // 6, user mac
	seq  uint16	// 2
	unix uint64	// 8
}

func (me *privateAuth) init(mac Mac) {
	copy(me.mac[:], mac)
	me.seq = authSeq() // needn't lock
	me.unix = uint64(time.Now().Unix())
}

func (me *privateAuth) ToBinary(bin []byte) error {
	if len(bin) < AuthSize {
		return Error
	}
	
	copy(bin, me.mac[:])
	binary.BigEndian.PutUint16(bin[6:], me.seq)
	binary.BigEndian.PutUint64(bin[8:], me.unix)
	
	return nil
}

func (me *privateAuth) FromBinary(bin []byte) error {
	if len(bin) < AuthSize {
		return Error
	}
	
	copy(me.mac[:], bin)
	me.seq = binary.BigEndian.Uint16(bin[6:])
	me.unix = binary.BigEndian.Uint64(bin[8:])
	
	return nil
}

type PktAuth []byte

func (me PktAuth) md5(pkt []byte, auth PktAuth, secret []byte) error {
	md5 := md5.New()
	
	md5.Write(pkt[:4]) 	// Code+ID+Length	
	md5.Write(auth)		// auth
	md5.Write(pkt[PktHdrSize:]) // Attributes
	md5.Write(secret)	// Key

	sum := md5.Sum(nil)
	
	if len(me) < len(sum) {
		return Error
	}
	
	copy(me, sum)
	return nil
}

// AccessRequest=Authenticator
func (me PktAuth) AuthRequest(mac Mac) error {
	auth := &privateAuth{}
	
	auth.init(mac)
	
	return auth.ToBinary(me)
}

// AccessReponse = MD5(Code+ID+Length+AccessRequest+Attributes+Key)
// pkt is the Access-Accept/Reject packet
func (me PktAuth) AuthReponse(pkt []byte, req PktAuth, secret []byte) error {
	return me.md5(pkt, req, secret)
}

// AcctRequest = MD5(Code+ID+Length+16ZeroOctets+Attributes+Key)
// pkt is the Accounting-Request packet
func (me PktAuth) AcctRequest(pkt []byte, secret []byte) error {
	zero := PktAuth{}
	
	return me.md5(pkt, zero, secret)
}

// AcctReponse = MD5(Code+ID+Length+AcctRequest+Attributes+Key)
// pkt is the Accounting-Response packet
func (me PktAuth) AcctReponse(pkt []byte, req PktAuth, secret []byte) error {
	return me.md5(pkt, req, secret)
}
