package engine

import "github.com/msisdev/dotato/internal/factory"

func (e *Engine) readConfig() (err error) {
	if e.cfg != nil {
		return
	}

	e.cfg, e.cdir, err = factory.ReadConfig(e.fs)
	return
}

func (e *Engine) readIgnore() (err error) {
	if e.ig != nil {
		return
	}
	if err = e.readConfig(); err != nil {
		return
	}

	e.ig, err = factory.ReadIgnore(e.fs, e.cdir)
	return
}
