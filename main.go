package main

import (
	"github.com/bitwurx/jrpc2"
	"github.com/factorysh/go-longrun/longrun"
)

func main() {

	s := jrpc2.NewServer(":8888", "/api/v1/rpc", map[string]string{})

	l := longrun.New()
	s.Register("longrun.next", jrpc2.Method{Method: l.Next})

	s.Start()

}
