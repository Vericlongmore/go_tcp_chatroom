package main

import (
	"time"
)

func main() {
	initRedis("localhost:6379", 16, 1024, time.Second*300)
	initUserMgr()
	runServer("localhost:20001")
}
