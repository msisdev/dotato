package cfg

import "sort"

type GroupList []string

func (g GroupList) ToGroupSet() GroupSet {
	gs := make(GroupSet)
	for _, item := range g {
		gs[item] = true
	}
	return gs
}

func (g GroupList) IsEqual(other GroupList) bool {
	// Compare length first
	if len(g) != len(other) {
		return false
	}

	// Sort and compare items
	sort.Slice(g, func(i, j int) bool { return g[i] < g[j] })
	sort.Slice(other, func(i, j int) bool { return other[i] < other[j] })
	for i := range g {
		if g[i] != other[i] {
			return false
		}
	}
	return true
}

///////////////////////////////////////////////////////////////////////////////

type GroupSet map[string]bool

func (gs GroupSet) ToGroupList() GroupList {
	gl := make(GroupList, 0, len(gs))
	for item := range gs {
		gl = append(gl, item)
	}
	return gl
}

func (gs GroupSet) IsEqual(other GroupSet) bool {	
	// Compare length first
	if len(gs) != len(other) {
		return false
	}

	// Compare items
	for item := range gs {
		if !other[item] {
			return false
		}
	}
	return true
}
