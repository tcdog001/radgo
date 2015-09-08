package radgo

import (
	. "asdf"
	"time"
)

const AcctSessionIdLength = 12*2+14

func NewSessionId(user Mac, dev Mac) []byte {
	session := []byte{}
	
	session = append(session, []byte(user.ToStringS())...)
	session = append(session, []byte(dev.ToStringS())...)
	session = append(session, []byte(time.Now().Format("20060102150405"))...)
	
	return session
}
