package radgo

import (
	"math/rand"
	"time"
)

var rSeed = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
	initDebug()
	initPktCode()
	initAttrType()
	initAttrTable()
	initDeauthReason()
	initAec()
	initAtc()
	initAst()
	initAvt()
	initAast()
	initAnpt()
	
	radRun()
}
