// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/han-y-s/hertz_demo/caller"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8887"))
	//caller.ClientInit()
	caller.GeneralCliInit()
	register(h)
	h.Spin()
}
