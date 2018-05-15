package g

import (
	"log"
	"runtime"
)

var consulAdress string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	consulAdress = "127.0.0.1:8500"
}
