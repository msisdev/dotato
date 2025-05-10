package app

import (
	"os"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

func (a App) WalkImportFile(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	onDot := func(dot gp.GardenPath, fi os.FileInfo) error {
		// Get dtt path
		dtt := a.e.DotToDtt(base, dot, group)

		// Get preview
		pre, err := a.PreviewImportFile(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return a.e.WalkDotDir(group, base, onDot)
}

func (a App) WalkImportLink(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	onDot := func(dot gp.GardenPath, fi os.FileInfo) error {
		// Get dtt path
		dtt := a.e.DotToDtt(base, dot, group)

		// Get preview
		pre, err := a.PreviewImportLink(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return a.e.WalkDotDir(group, base, onDot)
}

func (a App) WalkExportFile(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	onDtt := func(dtt gp.GardenPath, fi os.FileInfo) error {
		// Get dot path
		dot := a.e.DttToDot(base, dtt)

		// Get preview
		pre, err := a.PreviewExportFile(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return a.e.WalkDttDir(group, onDtt)
}

func (a App) WalkExportLink(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	onDtt := func(dtt gp.GardenPath, fi os.FileInfo) error {
		// Get dot path
		dot := a.e.DttToDot(base, dtt)

		// Get preview
		pre, err := a.PreviewExportLink(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return a.e.WalkDttDir(group, onDtt)
}

func (a App) WalkUnlink(
	group string,
	base gp.GardenPath,
	onPreview func(Preview) error,
) (err error) {
	onDot := func(dot gp.GardenPath, fi os.FileInfo) error {
		// Get dtt path
		dtt := a.e.DotToDtt(base, dot, group)

		// Get preview
		pre, err := a.PreviewUnlink(dot, dtt)
		if err != nil {
			return err
		}

		return onPreview(*pre)
	}

	return a.e.WalkDotDir(group, base, onDot)
}
