package radgo

import (
	. "asdf"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const (
	fileDebug = "radius.debug"
)

type DebugConfig struct {
	Users []string
}

func (me *DebugConfig) Load() {
	f, err := os.Open(fileDebug)
	if nil != err {
		fmt.Println("no found", fileDebug)

		return // no debug file, just return
	}
	defer f.Close()

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(f)
	if nil != err {
		panic(err)
	}

	json.Unmarshal(b.Bytes(), me)

	fmt.Printf("load debug file(%s):%+v" + Crlf, fileDebug, me)
}

type DebugControl struct {
	Users map[[6]byte]bool
}

func (me *DebugControl) Load(c *DebugConfig) {
	me.Users = make(map[[6]byte]bool)

	mac := [6]byte{}
	me.Users[mac] = true

	for _, s := range c.Users {
		Mac(mac[:]).FromString(s)

		me.Users[mac] = true
		fmt.Println("load debug user:", s)
	}
}

var debugControl = DebugControl{
	Users: nil,
}

func initDebug() {
	c := &DebugConfig{}

	c.Load()

	(&debugControl).Load(c)
}

func debugUser(mac Mac, format string, v ...interface{}) {
	if nil != debugControl.Users {
		user := [6]byte{mac[0], mac[1], mac[2], mac[3], mac[4], mac[5]}

		if d, ok := debugControl.Users[user]; d && ok {
			Log.Debug(format, v)
		}
	}
}

func debugUserError(mac Mac, v ...interface{}) {
	debugUser(mac, Empty, v)
}
