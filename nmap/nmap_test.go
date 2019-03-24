package nmap

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	_nmap "github.com/t94j0/nmap"
)

func TestNmap(t *testing.T) {
	scan, err := _nmap.Init().AddHosts("blog.garambrogne.net").AddPorts(22, 80, 443).Run()
	assert.NoError(t, err)
	h, ok := scan.GetHost("blog.garambrogne.net")
	assert.True(t, ok)
	b, err := json.MarshalIndent(h, "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(b))
}
