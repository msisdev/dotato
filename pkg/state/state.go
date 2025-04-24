package state

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

