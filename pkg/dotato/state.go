package dotato

import (
	"github.com/msisdev/dotato/pkg/state"
)

func (d *Dotato) ReadState() (err error) {
	if d.state != nil {
		return
	}

	d.state, err = state.New(getStatePathUnsafe())
	return
}

