package radgo

import (
	. "asdf"
	"time"
	"net"
)

type IAuth interface {
	IAcct
	
	UserPassword() []byte
}

type IAcct interface {
	IParam
	
	SessionId() []byte
	UserName() []byte
	UserMac() []byte // binary mac
	UserIp() uint32
	AcctInputOctets() uint32
	AcctOutputOctets() uint32
	AcctInputGigawords() uint32
	AcctOutputGigawords() uint32
	AcctTerminateCause() uint32
}

type IParam interface {
	Secret() []byte
	NasIdentifier() []byte
	NasIpAddress() uint32
	NasPort() uint32
	NasPortType() uint32
	NasPortId() uint32
	ServiceType() uint32
	Server () string
	AuthPort() string
	AcctPort() string
	Timeout() int // ms
}

type Policy struct {
	IdleTimeout uint32
	OnlineTime 	uint32
	FlowLimit 	uint32
	RateLimit 	uint32
}

type client struct {
	request 	Packet
	response 	Packet
	bin [PktLengthMax]byte
	rlen 		int
	sessionId 	[AcctSessionIdLength]byte
	
	//cache
	mac 		[]byte
	
	// socket
	remote 		*net.UDPAddr
	conn 		*net.UDPConn
}

func clientNew(mac Mac) *client {
	c := &client{mac:mac}
	
	c.init()
	
	return c
}

func (me *client) init() {
	(&me.request).Init()
	(&me.response).Init()
}

func (me *client) debugError(err error) error{
	debugUserError(me.mac, err)
	
	return err
}
	
func (me *client) initConn(r IAcct) error {
	err := error(nil)
	
	me.conn, err = net.DialUDP("udp", nil, me.remote)
	if nil!=err {
		return me.debugError(err)
	}
	
	err = me.conn.SetReadDeadline(time.Now().Add(time.Duration(r.Timeout()) * time.Millisecond))
	if nil!=err {
		return me.debugError(err)
	}
	
	return nil
}

func (me *client) initAuth(r IAuth) error {
	pkt := &me.request
	
	pkt.Code = AccessRequest
	pkt.Id	= PktId()
	if err := PktAuth(pkt.Auth[:]).AuthRequest(r.UserMac()); nil!=err {
		return me.debugError(err)
	}
	
	if err := pkt.SetAttrStringList([]AttrString{
		{
			Type:UserName,
			Value:r.UserName(),
		},
		{
			Type:UserPassword,
			Value:r.UserPassword(),
		},
		{
			Type:CallingStationId,
			Value:[]byte(Mac(r.UserMac()).ToString()),
		},
		{
			Type:NasIdentifier,
			Value:r.NasIdentifier(),
		},
	}); nil!=err {
		return me.debugError(err)
	}
	
	if err := pkt.SetAttrNumberList([]AttrNumber{
		{
			Type:FramedIpAddress,
			Value:r.UserIp(),
		},
		{
			Type:NasIpAddress,
			Value:r.NasIpAddress(),
		},
		{
			Type:NasPort,
			Value:r.NasPort(),
		},
		{
			Type:NasPortType,
			Value:r.NasPortType(),
		},
		{
			Type:NasPortId,
			Value:r.NasPortId(),
		},
		{
			Type:ServiceType,
			Value:r.ServiceType(),
		},
	}); nil!=err {
		return me.debugError(err)
	}
	
	return nil
}

func (me *client) initAcct(r IAcct, action EAastValue) error {
	pkt := &me.request
	
	pkt.Code = AccountingRequest
	pkt.Id	= PktId()
	if err := PktAuth(pkt.Auth[:]).AcctRequest(me.bin[:], r.Secret()); nil!=err {
		return me.debugError(err)
	}
		
	if err := pkt.SetAttrStringList([]AttrString{
		{
			Type:UserName,
			Value:r.UserName(),
		},
		{
			Type:CallingStationId,
			Value:[]byte(Mac(r.UserMac()).ToString()),
		},
		{
			Type:AcctSessionId,
			Value:r.SessionId(),
		},
		{
			Type:NasIdentifier,
			Value:r.NasIdentifier(),
		},
	}); nil!=err {
		return me.debugError(err)
	}
	
	if err := pkt.SetAttrNumberList([]AttrNumber{
		{
			Type:FramedIpAddress,
			Value:r.UserIp(),
		},
		{
			Type:AcctStatusType,
			Value:uint32(action),
		},
		{
			Type:AcctDelayTime,
			Value:0, // fix 0 ???
		},
		{
			Type:EventTimestamp,
			Value:uint32(time.Now().Unix()),
		},
		{
			Type:NasIpAddress,
			Value:r.NasIpAddress(),
		},
	}); nil!=err {
		return me.debugError(err)
	}
	
	if AastStart!=action {
		if err := pkt.SetAttrNumberList([]AttrNumber{
			{
				Type:AcctInputOctets,
				Value:r.AcctInputOctets(),
			},
			{
				Type:AcctInputGigawords,
				Value:r.AcctInputGigawords(),
			},
			{
				Type:AcctOutputOctets,
				Value:r.AcctOutputOctets(),
			},
			{
				Type:AcctOutputGigawords,
				Value:r.AcctOutputGigawords(),
			},
			{
				Type:AcctDelayTime,
				Value:0, // fix 0 ???
			},
		}); nil!=err {
			return me.debugError(err)
		}
	}
	
	if AastStop==action {
		if err := pkt.SetAttrNumberList([]AttrNumber{
			{
				Type:AcctTerminateCause,
				Value:r.AcctTerminateCause(),
			},
		}); nil!=err {
			return me.debugError(err)
		}
	}
	
	return nil
}

func (me *client) net() error {
	err := error(nil)
	
	if _, err = me.conn.Write(me.bin[:me.request.Len]); nil!=err {
		return me.debugError(err)
	}
	
	if me.rlen, err = me.conn.Read(me.bin[:]); nil!=err {
		return me.debugError(err)
	}
	
	return nil
}

func (me *client) auth(r IAuth) (*Policy, error) {
	err := error(nil)
	
	me.remote , err = net.ResolveUDPAddr("udp4", r.Server() + ":" + r.AuthPort())
	if nil!=err {
		return nil, me.debugError(err)
	}
	
	if err := me.initConn(r); nil!=err {
		return nil, me.debugError(err)
	}
	
	if err := me.initAuth(r); nil!=err {
		return nil, me.debugError(err)
	}
	
	q := &me.request
	if err := q.ToBinary(me.bin[:]); nil!=err {
		return nil, me.debugError(err)
	}
	
	if err := me.net(); nil!=err {
		return nil, me.debugError(err)
	}
	
	p := &me.response
	if err := p.FromBinary(me.bin[:]); nil!=err {
		return nil, me.debugError(err)
	}
	
	policy := &Policy{}
	p.Policy(policy)
	
	return policy, nil
}

func (me *client) acct(r IAcct, action EAastValue) (bool, error) {
	err := error(nil)
	
	me.remote , err = net.ResolveUDPAddr("udp4", r.Server() + ":" + r.AcctPort())
	if nil!=err {
		return false, me.debugError(err)
	}
	
	if err := me.initConn(r); nil!=err {
		return false, me.debugError(err)
	}
	
	if err := me.initAcct(r, action); nil!=err {
		return false, me.debugError(err)
	}
	
	q := &me.request
	if err := q.ToBinary(me.bin[:]); nil!=err {
		return false, me.debugError(err)
	}
	
	if err := me.net(); nil!=err {
		return false, me.debugError(err)
	}
	
	p := &me.response
	if err := p.FromBinary(me.bin[:]); nil!=err {
		return false, me.debugError(err)
	}
	
	return true, nil
}

func ClientSessionId(mac Mac, session []byte) error{
	s := &SessionId{}
	s.Init(mac)
	
	return s.ToBinary(session)
}

func ClientAuth(r IAuth) (*Policy, error) {
	c := clientNew(r.UserMac())
	defer func() {c=nil}()
	
	return c.auth(r)
}

func ClientAcctStart(r IAcct) (bool, error) {
	c := clientNew(r.UserMac())
	defer func() {c=nil}()
	
	return c.acct(r, AastStart)
}

func ClientAcctUpdate(r IAcct) (bool, error) {
	c := clientNew(r.UserMac())
	defer func() {c=nil}()
	
	return c.acct(r, AastInterimUpdate)
}

func ClientAcctStop(r IAcct) (bool, error) {
	c := clientNew(r.UserMac())
	defer func() {c=nil}()
	
	return c.acct(r, AastStop)
}
