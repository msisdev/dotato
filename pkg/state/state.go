package state

import "github.com/msisdev/dotato/pkg/config"

type State struct {
	d *DB
}

func NewState(path string) (*State, error) {
	d, err := NewDB(path)
	if err != nil {
		return nil, err
	}
	return &State{d: d}, nil
}

func (s State) GetAllLink() ([]History, error) {
	return s.d.v1_getAllByMode(config.ModeLink)
}

