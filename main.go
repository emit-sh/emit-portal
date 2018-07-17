package main

import (
	"fmt"
	"github.com/emit-sh/emit-portal/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		fmt.Print(err)
		return
	}
	server.Start()
}