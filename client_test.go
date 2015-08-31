package radgo

import (
	. "asdf"
//	"fmt"
	"testing"
)

type Param struct {
	Secret        []byte
	NasIdentifier []byte
	NasIpAddress  uint32
	NasPort       uint32
	NasPortType   uint32
	ServiceType   uint32
	Server        string
	AuthPort      string
	AcctPort      string
	Timeout       int // ms
}

var param = &Param{
	Secret:        []byte("testing123"),
	NasIdentifier: []byte("ums.autelan.com"),
	NasPort:       0,
	NasPortType:   uint32(AnptIeee80211),
	ServiceType:   uint32(AstLogin),
	Server:        "116.228.184.202",
	AuthPort:      "1812",
	AcctPort:      "1813",
	Timeout:       3000,
}

type User struct {
	ssid []byte
	dev  [6]byte

	passwd    []byte
	sessionid [AcctSessionIdLength]byte
	name      []byte
	mac       [6]byte // binary mac
	ip        uint32
	input     uint32
	output    uint32
	inputg    uint32
	outputg   uint32
	reason    uint32

	class []byte
}

var user = &User{
	ssid: []byte("i-shanghai"),

	name:    []byte("10000000000@windfind.static@ish"),
	passwd:  []byte("123456"),
	input:   1000 * 1000,
	output:  1000 * 2000,
	inputg:  0,
	outputg: 0,
	reason:  uint32(AtcUserRequest),
}

func testInit(t *testing.T) {
	//param init
	param.NasIpAddress = uint32(IpAddressFromString("120.26.47.127"))
	t.Logf("test init param:%#v" + Crlf, param)
	
	//user init
	Mac(user.mac[:]).FromString("F8:95:C7:D9:37:74")
	Mac(user.dev[:]).FromString("00:1f:64:00:00:01")
	user.ip = uint32(IpAddressFromString("192.168.100.200"))
	ClientSessionId(user.mac[:], user.sessionid[:])
	t.Logf("test init user:%#v" + Crlf, user)
}

func (me *User) SSID() []byte {
	return me.ssid
}

func (me *User) DevMac() []byte {
	return me.dev[:]
}

func (me *User) UserPassword() []byte {
	return me.passwd
}

func (me *User) SessionId() []byte {
	return me.sessionid[:]
}

func (me *User) UserName() []byte {
	return me.name
}

func (me *User) UserMac() []byte {
	return me.mac[:]
}

func (me *User) UserIp() uint32 {
	return me.ip
}

func (me *User) AcctInputOctets() uint32 {
	return me.input
}

func (me *User) AcctOutputOctets() uint32 {
	return me.output
}

func (me *User) AcctInputGigawords() uint32 {
	return me.inputg
}

func (me *User) AcctOutputGigawords() uint32 {
	return me.outputg
}

func (me *User) AcctTerminateCause() uint32 {
	return me.reason
}

func (me *User) GetClass() []byte {
	return me.class
}

func (me *User) SetClass(class []byte) {
	me.class = class
}

func (me *User) Secret() []byte {
	return param.Secret
}

func (me *User) NasIdentifier() []byte {
	return param.NasIdentifier
}

func (me *User) NasIpAddress() uint32 {
	return param.NasIpAddress
}

func (me *User) NasPort() uint32 {
	return param.NasPort
}

func (me *User) NasPortType() uint32 {
	return param.NasPortType
}

func (me *User) ServiceType() uint32 {
	return param.ServiceType
}

func (me *User) Server() string {
	return param.Server
}

func (me *User) AuthPort() string {
	return param.AuthPort
}

func (me *User) AcctPort() string {
	return param.AcctPort
}

func (me *User) Timeout() int {
	return param.Timeout
}

func testAuth(t *testing.T) {
	t.Log("testing auth start ...")

	policy, err, authError := ClientAuth(user)
	if nil != err || nil != authError {
		t.Fatal("test auth error:", err, authError)
	}
	t.Logf("test auth get policy:%#v" + Crlf, policy)
	
	t.Log("test auth PASS")
}

func testAcctStart(t *testing.T) {
	t.Log("testing acct start ...")
	
	if err, acctError := ClientAcctStart(user); nil != err || nil != acctError {
		t.Fatal("test acct start error:", err, acctError)
	}
	
	t.Log("test acct start PASS")
}

func testAcctUpdate(t *testing.T) {
	t.Log("testing acct update ...")
	
	if err, acctError := ClientAcctUpdate(user); nil != err || nil != acctError {
		t.Fatal("test acct update error:", err, acctError)
	}
	
	t.Log("test acct update PASS")
}

func testAcctStop(t *testing.T) {
	t.Log("testing acct stop ...")
	
	if err, acctError := ClientAcctStop(user); nil != err || nil != acctError {
		t.Fatal("test acct stop error:", err, acctError)
	}
	
	t.Log("test acct stop PASS")
}

func TestClient(t *testing.T) {
	testInit(t)
	testAuth(t)
	testAcctStart(t)
	testAcctUpdate(t)
	testAcctStop(t)
}
