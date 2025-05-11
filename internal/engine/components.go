package engine

import (
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

// Config /////////////////////////////////////////////////////////////////////

func (e Engine) GetConfigDir() (gp.GardenPath, error) {
	if err := e.readConfig(); err != nil {
		return nil, err
	}

	return e.cdir, nil
}

func (e Engine) GetConfigVersion() (string, error) {
	if err := e.readConfig(); err != nil {
		return "", err
	}

	return e.cfg.Version, nil
}

func (e Engine) GetConfigMode() (string, error) {
	if err := e.readConfig(); err != nil {
		return "", err
	}

	return e.cfg.Mode, nil
}

func (e Engine) GetConfigPlans() (map[string][]string, error) {
	if err := e.readConfig(); err != nil {
		return nil, err
	}

	return e.cfg.Plans, nil
}

func (e Engine) GetConfigGroups(plan string) (map[string]bool, bool, error) {
	if err := e.readConfig(); err != nil {
		return nil, false, err
	}

	// Get groups
	groupList, ok := e.cfg.Plans[plan]
	if !ok {
		return nil, false, nil
	}

	// Convert to map[string]bool
	groupSet := make(map[string]bool)
	for _, group := range groupList {
		groupSet[group] = true
	}

	return groupSet, true, nil
}

func (e Engine) GetConfigGroupAll() (map[string]bool, error) {
	if err := e.readConfig(); err != nil {
		return nil, err
	}

	// Get groups
	groups := make(map[string]bool)
	for group := range e.cfg.Groups {
		groups[group] = true
	}

	return groups, nil
}

func (e Engine) GetConfigGroupBase(group, resolver string) (gp.GardenPath, []string, error) {
	if err := e.readConfig(); err != nil {
		return nil, nil, err
	}

	// Get group base
	return e.cfg.GetGroupBase(group, resolver)
}

func (e Engine) GetConfigGroupResolvers(group string) (map[string]string, error) {
	if err := e.readConfig(); err != nil {
		return nil, err
	}

	rs := make(map[string]string)

	// For each resolver
	for name, resolver := range e.cfg.Groups[group] {
		// Collect resolver
		rs[name] = resolver
	}

	return rs, nil
}
