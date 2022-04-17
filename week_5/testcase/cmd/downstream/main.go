package main

import "week_5/testcase/server"

func main() {
	server.NewDownStreamServer(0.2).Run(":8000")
}
