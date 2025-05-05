package cli

import (
	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/pkg/config"
	"github.com/msisdev/dotato/pkg/dotato"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/state"
)

func _importLink(
	fs billy.Filesystem,
	dtt *dotato.Dotato,
	dotfile gp.GardenPath,
	dotato gp.GardenPath,
) error {

	// Create directory
	if err := fs.MkdirAll(dotato.Parent().Abs(), 0755); err != nil {
		return err
	}

	// Move file
	if err := fs.Rename(dotfile.Abs(), dotato.Abs()); err != nil {
		return err
	}

	// Create link
	if err := fs.Symlink(dotato.Abs(), dotfile.Abs()); err != nil {
		return err
	}

	// Update history
	return dtt.PutHistory(state.History{
		DotPath: dotfile.Abs(),
		DttPath: dotato.Abs(),
		Mode: config.ModeLink,
	})
}

func _exportLink() {}

func _exportFile() {}
