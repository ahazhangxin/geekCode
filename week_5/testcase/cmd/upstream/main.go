package main

import (
	"time"
	"week_5/testcase/server"
)

func main() {
	server.NewUpStreamServer(
		10,
		50,
		0.8,
		time.Second*5,
	).Run(":9000")
}
