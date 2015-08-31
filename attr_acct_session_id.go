package radgo

import (
	. "asdf"
	"time"
)

const AcctSessionIdLength = 12*2+14

func NewSessionId(user Mac, dev Mac) string {
	return 	user.ToStringS() + 
			dev.ToStringS() + 
			time.Now().Format("20060102150405")
}

