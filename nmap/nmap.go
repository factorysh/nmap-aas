package nmap

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bitwurx/jrpc2"

	_nmap "github.com/t94j0/nmap"
)

type NmapParams struct {
	Hosts []string `json:"hosts"`
	Ports []uint16 `json:"ports"`
}

func (p *NmapParams) FromPositional(params []interface{}) error {
	if len(params) != 2 {
		return errors.New("Two arguments")
	}
	p.Hosts = params[0].([]string)
	p.Ports = params[1].([]uint16)
	return nil
}

func Nmap(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(NmapParams)
	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}
	scan, err := _nmap.Init().AddPorts(p.Ports...).AddHosts(p.Hosts...).Run()
	if err != nil {
		return nil, &jrpc2.ErrorObject{
			Code:    jrpc2.InternalErrorCode,
			Message: jrpc2.InternalErrorMsg,
			Data:    err.Error(),
		}
	}
	fmt.Println(scan)

	return scan.ToString(), nil
}
