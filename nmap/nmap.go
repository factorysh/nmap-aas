package nmap

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/factorysh/go-longrun/run"

	"github.com/bitwurx/jrpc2"

	_nmap "github.com/t94j0/nmap"
	"golang.org/x/sync/semaphore"
)

type NmapParams struct {
	Hosts []string `json:"hosts"`
	Ports []uint16 `json:"ports"`
}

type Scan struct {
	params *NmapParams
	run    *run.Run
}

func (p *NmapParams) FromPositional(params []interface{}) error {
	if len(params) != 2 {
		return errors.New("Two arguments")
	}
	p.Hosts = params[0].([]string)
	p.Ports = params[1].([]uint16)
	return nil
}

type NmapPool struct {
	runs    *run.Runs
	waiting chan *Scan
	sem     *semaphore.Weighted
	ctx     context.Context
}

func New(ctx context.Context, runs *run.Runs, workers int) *NmapPool {
	np := &NmapPool{
		runs:    runs,
		waiting: make(chan *Scan),
		sem:     semaphore.NewWeighted(int64(workers)),
		ctx:     ctx,
	}
	go func() {
		for {
			ctx := context.Background()
			err := np.sem.Acquire(ctx, 1)
			defer np.sem.Release(1)
			if err != nil {
				panic(err)
			}
			scan := <-np.waiting
			doNmap(scan)
		}
	}()
	return np
}

func doNmap(scan *Scan) {
	scans, err := _nmap.Init().AddPorts(scan.params.Ports...).AddHosts(scan.params.Hosts...).Run()
	if err != nil {
		scan.run.Error(err)
		return
	}
	for _, value := range scans.Hosts {
		scan.run.Run(value.ToString())
	}
	scan.run.Success(nil)
}

func (np *NmapPool) Nmap(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(NmapParams)
	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}
	u := np.runs.New()

	np.waiting <- &Scan{
		run:    u,
		params: p,
	}
	return u.Id(), nil
}
