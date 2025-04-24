package run

import (
	"github.com/msisdev/dotato/internal/arg"
)

// Recover all linked files by dotato.
//
// 1. Determine linked files by dotato.
//
// 2. Replace each symlink with copied file.
func (r Runner) DangerCmd(arg arg.DangerUnlinkArgs) error {
	r.initState()

	return nil
}
