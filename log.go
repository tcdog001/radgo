package radgo

import (
	. "asdf"
)

var log ILogger

func SetLogger(r ILogger) {
	log = r
}
