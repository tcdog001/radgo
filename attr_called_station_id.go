package radgo

import (
	. "asdf"
)

func MakeCalledStationId(mac []byte, ssid []byte) string {
	return Mac(mac).ToString() + ":" + string(ssid)
}
