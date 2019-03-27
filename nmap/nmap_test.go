package nmap

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/google/uuid"

	"github.com/factorysh/go-longrun/run"

	"github.com/stretchr/testify/assert"
)

func TestNmap(t *testing.T) {
	ctx := context.Background()
	runs := run.New()
	np := New(ctx, runs, 3)

	id, jerr := np.Nmap([]byte(`
	{ 
		"hosts": ["blog.garambrogne.net", "factory.sh", "bearstech.com"],
		"ports": [22, 80, 443]
	}
	`))
	assert.Nil(t, jerr)
	uid, ok := id.(uuid.UUID)
	assert.True(t, ok)
	i := 0
	stop := false
	for {
		evts, err := runs.Get(uid, i)
		assert.Nil(t, err)
		i += len(evts)
		for _, evt := range evts {
			spew.Dump(evt)
			stop = stop || evt.Ended()
		}
		if stop {
			break
		}
	}
}
