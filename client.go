package radgo

import (
	. "asdf"
	"net"
	"time"
	"errors"
)

const (
	RadPrivateBegin 		RadPrivate = 0
	
	RadPrivateClass 		RadPrivate = 0
	RadPrivateChapChallenge	RadPrivate = 1
	
	RadPrivateEnd 			RadPrivate = 2
)
type RadPrivate uint32

type IAuth interface {
	IAcct

	UserPassword() []byte
}

type IAcct interface {
	IParam

	SSID() []byte
	DevMac() []byte
	SessionId() []byte
	UserName() []byte
	UserMac() []byte // binary mac
	UserIp() uint32
	AcctInputOctets() uint32
	AcctOutputOctets() uint32
	AcctInputGigawords() uint32
	AcctOutputGigawords() uint32
	AcctTerminateCause() uint32

	GetPrivate(t RadPrivate) interface{}
	SetPrivete(t RadPrivate, e interface{})
}

type IParam interface {
	Secret() []byte
	NasIdentifier() []byte
	NasIpAddress() uint32
	NasPort() uint32
	NasPortType() uint32
	//	NasPortId() uint32
	ServiceType() uint32
	Server() string
	AuthType() uint32 // 0:pap, 1:chap
	AuthPort() string
	AcctPort() string
	Timeout() uint32 // ms
}

type Policy struct {
	IdleTimeout uint32
	OnlineTime  uint32

	UpFlowLimit uint64
	UpRateMax   uint32
	UpRateAvg   uint32

	DownFlowLimit uint64
	DownRateMax   uint32
	DownRateAvg   uint32
}

type client struct {
	request   Packet
	response  Packet
	bin       [PktLengthMax]byte
	rlen      int
	sessionId [AcctSessionIdLength]byte

	//cache
	mac []byte

	// socket
	remote *net.UDPAddr
	conn   *net.UDPConn
}

func newClient(mac Mac) *client {
	c := &client{mac: mac}

	c.init()

	return c
}

func userInit(r IAuth) {
	if authChap==r.AuthType() {
		r.SetPrivete(RadPrivateChapChallenge, 
			newChapChallenge(r.UserMac(), r.DevMac()))
	}
}

func userFini(r IAuth) {
	for i:=RadPrivateBegin; i<RadPrivateEnd; i++ {
		r.SetPrivete(i, nil)
	}
}

func (me *client) init() {
	(&me.request).Init()
	(&me.response).Init()
}

func (me *client) debugError(err error) error {
	debugUserError(me.mac, err)

	return err
}

func (me *client) initConn(r IAcct) error {
	err := error(nil)

	me.conn, err = net.DialUDP("udp", nil, me.remote)
	if nil != err {
		return me.debugError(err)
	}

	err = me.conn.SetReadDeadline(time.Now().Add(time.Duration(r.Timeout()) * time.Millisecond))
	if nil != err {
		return me.debugError(err)
	}

	return nil
}

func (me *client) initAuth(r IAuth) error {
	q := &me.request

	q.Code = AccessRequest
	q.Id = PktId()
	if err := PktAuth(q.Auth[:]).AuthRequest(r.UserMac()); nil != err {
		return me.debugError(err)
	}

	switch r.AuthType() {
		case authPap:
			password := enPapPassword(q.Auth[:], r.Secret(), r.UserPassword())
			if err := q.SetAttrString(UserPassword, password); nil!=err {
				return err
			}
		case authChap:
			challenge, _ := r.GetPrivate(RadPrivateChapChallenge).([]byte)
			if nil==challenge {
				return ErrBadIntf
			}
			password := enChapPassword(r.UserPassword(), challenge)
			
			if err := q.SetAttrString(ChapChallenge, challenge); nil!=err {
				return err
			}
			if err := q.SetAttrString(ChapPassword, password); nil!=err {
				return err
			}
		default:
			return ErrBadType
	}
	
	if err := q.SetAttrStringList([]AttrString{
		{
			Type:  UserName,
			Value: r.UserName(),
		},
		{
			Type:  CalledStationId,
			Value: MakeCalledStationId(r.DevMac(), r.SSID()),
		},
		{
			Type:  CallingStationId,
			Value: []byte(Mac(r.UserMac()).ToString()),
		},
		{
			Type:  NasIdentifier,
			Value: r.NasIdentifier(),
		},
	}); nil != err {
		return me.debugError(err)
	}

	if err := q.SetAttrNumberList([]AttrNumber{
		{
			Type:  FramedIpAddress,
			Value: r.UserIp(),
		},
		{
			Type:  NasIpAddress,
			Value: r.NasIpAddress(),
		},
		{
			Type:  NasPort,
			Value: r.NasPort(),
		},
		{
			Type:  NasPortType,
			Value: r.NasPortType(),
		},
		{
			Type:  ServiceType,
			Value: r.ServiceType(),
		},
	}); nil != err {
		return me.debugError(err)
	}

	return nil
}

func (me *client) initAcct(r IAcct, action EAastValue) error {
	q := &me.request

	q.Code = AccountingRequest
	q.Id = PktId()

	class, _ := r.GetPrivate(RadPrivateClass).([]byte)
	
	if err := q.SetAttrStringList([]AttrString{
		{
			Type:  UserName,
			Value: r.UserName(),
		},
		{
			Type:  CalledStationId,
			Value: MakeCalledStationId(r.DevMac(), r.SSID()),
		},
		{
			Type:  CallingStationId,
			Value: []byte(Mac(r.UserMac()).ToString()),
		},
		{
			Type:  AcctSessionId,
			Value: r.SessionId(),
		},
		{
			Type:  NasIdentifier,
			Value: r.NasIdentifier(),
		},
		{
			Type:  Class,
			Value: class,
		},
	}); nil != err {
		return me.debugError(err)
	}

	if err := q.SetAttrNumberList([]AttrNumber{
		{
			Type:  FramedIpAddress,
			Value: r.UserIp(),
		},
		{
			Type:  AcctStatusType,
			Value: uint32(action),
		},
		{
			Type:  AcctDelayTime,
			Value: 0, // fix 0 ???
		},
		{
			Type:  EventTimestamp,
			Value: uint32(time.Now().Unix()),
		},
		{
			Type:  NasIpAddress,
			Value: r.NasIpAddress(),
		},
		{
			Type:  NasPort,
			Value: r.NasPort(),
		},
		{
			Type:  NasPortType,
			Value: r.NasPortType(),
		},
		{
			Type:  ServiceType,
			Value: r.ServiceType(),
		},
	}); nil != err {
		return me.debugError(err)
	}

	if AastStop == action || AastInterimUpdate == action {
		if err := q.SetAttrNumberList([]AttrNumber{
			{
				Type:  AcctInputOctets,
				Value: r.AcctInputOctets(),
			},
			{
				Type:  AcctInputGigawords,
				Value: r.AcctInputGigawords(),
			},
			{
				Type:  AcctOutputOctets,
				Value: r.AcctOutputOctets(),
			},
			{
				Type:  AcctOutputGigawords,
				Value: r.AcctOutputGigawords(),
			},
		}); nil != err {
			return me.debugError(err)
		}
	}

	if AastStop == action {
		if err := q.SetAttrNumberList([]AttrNumber{
			{
				Type:  AcctTerminateCause,
				Value: r.AcctTerminateCause(),
			},
		}); nil != err {
			return me.debugError(err)
		}
	}

	return nil
}

func (me *client) net() error {
	err := error(nil)

	if _, err = me.conn.Write(me.bin[:me.request.Len]); nil != err {
		return me.debugError(err)
	}

	if me.rlen, err = me.conn.Read(me.bin[:]); nil != err {
		return me.debugError(err)
	}

	return nil
}

type AuthError error

func (me *client) auth(r IAuth) (*Policy, error, AuthError) {
	err := error(nil)

	me.remote, err = net.ResolveUDPAddr("udp4", r.Server()+":"+r.AuthPort())
	if nil != err {
		return nil, me.debugError(err), nil
	}

	if err := me.initConn(r); nil != err {
		return nil, me.debugError(err), nil
	}

	if err := me.initAuth(r); nil != err {
		return nil, me.debugError(err), nil
	}

	q := &me.request
	if err := q.ToBinary(me.bin[:]); nil != err {
		return nil, me.debugError(err), nil
	}

	if err := me.net(); nil != err {
		return nil, me.debugError(err), nil
	}

	p := &me.response
	if err := p.FromBinary(me.bin[:me.rlen]); nil != err {
		return nil, me.debugError(err), nil
	}

	if AccessAccept != p.Code {
		return nil, me.debugError(Error), nil
	}
	
	if authError := p.Attrs[ReplyMessage].GetString(); nil != authError {
		err := errors.New(string(authError))
		if IsGoodReplyMessage(authError) {
			return nil, nil, err
		} else {
			return nil, ErrUnknowReplyMessage, err
		}
	}
	
	if authClass := p.Attrs[Class].GetString(); nil!=authClass {
		r.SetPrivete(RadPrivateClass, authClass)
	}

	return p.Policy(), nil, nil
}

type AcctError error

func (me *client) acct(r IAcct, action EAastValue) (error, AcctError) {
	err := error(nil)

	me.remote, err = net.ResolveUDPAddr("udp4", r.Server()+":"+r.AcctPort())
	if nil != err {
		return me.debugError(err), nil
	}

	if err := me.initConn(r); nil != err {
		return me.debugError(err), nil
	}

	if err := me.initAcct(r, action); nil != err {
		return me.debugError(err), nil
	}

	q := &me.request
	if err := q.ToBinary(me.bin[:]); nil != err {
		return me.debugError(err), nil
	}

	if err := PktAuth(me.bin[4:PktHdrSize]).AcctRequest(me.bin[:me.request.Len], r.Secret()); nil != err {
		return me.debugError(err), nil
	}

	if err := me.net(); nil != err {
		return me.debugError(err), nil
	}

	p := &me.response
	if err := p.FromBinary(me.bin[:me.rlen]); nil != err {
		return me.debugError(err), nil
	}

	if AccountingResponse != p.Code {
		return me.debugError(Error), nil
	}
	
	if acctError := p.Attrs[ReplyMessage].GetString(); nil != acctError {
		err := errors.New(string(acctError))
		if IsGoodReplyMessage(acctError) {
			return nil, err
		} else {
			return ErrUnknowReplyMessage, err
		}
	}

	return nil, nil
}

func ClientAuth(r IAuth) (*Policy, error, AuthError) {
	c := newClient(r.UserMac())
	userInit(r)
	defer func() { c = nil }()

	return c.auth(r)
}

func ClientAcctStart(r IAcct) (error, AcctError) {
	c := newClient(r.UserMac())
	defer func() { c = nil }()

	return c.acct(r, AastStart)
}

func ClientAcctUpdate(r IAcct) (error, AcctError) {
	c := newClient(r.UserMac())
	defer func() { c = nil }()

	return c.acct(r, AastInterimUpdate)
}

func ClientAcctStop(r IAcct) (error, AcctError) {
	c := newClient(r.UserMac())
	defer func() { c = nil }()

	return c.acct(r, AastStop)
}
