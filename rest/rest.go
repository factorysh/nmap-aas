package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/factorysh/go-longrun/run"
	"github.com/factorysh/nmap-aas/nmap"
)

type Nmap struct {
	pool *nmap.NmapPool
}

func New(ctx context.Context, run *run.Runs) *Nmap {
	return &Nmap{
		pool: nmap.New(ctx, run, 5),
	}
}

func (n *Nmap) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var p nmap.NmapParams
	err = json.Unmarshal(raw, &p)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	u := n.pool.AddScan(&p)
	w.WriteHeader(201)
	w.Header().Set("X-Longrun-Id", u.Id().String())
	fmt.Fprintf(w, `{"longrun":"%v"}`, u.Id())
}
