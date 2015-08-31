package radgo

import (
	. "asdf"
	"time"
)

const AcctSessionIdLength = 32

func newSessionId(user Mac, dev Mac) string {
	return user.ToStringS() + dev.ToStringS() + time.Now().Format("20060102150405")
}

