package main

import (
	"context"

	"github.com/bitwurx/jrpc2"
	"github.com/factorysh/go-longrun/longrun"
	"github.com/factorysh/nmap-aas/nmap"
)

func main() {

	s := jrpc2.NewServer(":8888", "/api/v1/rpc", map[string]string{})

	l := longrun.New()
	s.Register("longrun.next", jrpc2.Method{Method: l.Next})
	n := nmap.New(context.Background(), l.Runs, 5)
	s.Register("nmap.scan", jrpc2.Method{Method: n.Nmap})
	s.Start()

}
