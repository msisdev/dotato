package dotato

import (
	"os"
	"runtime"
	"testing"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func getGardenPathFirstEl() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("SystemDrive")
	}
	return ""
}

func TestNewPathStat(t *testing.T) {
	var el = getGardenPathFirstEl()

	path := gp.GardenPath{el, "home", "user", ".bashrc"}
	real := gp.GardenPath{el, "home", "user", "Documents", "dotato", "bash", ".bashrc"}

	// Path: link / Real: file
	{
		d := requestDotato(real, FirstReq_File, path, SecondReq_Link_Same)
		stat, err := d.newPathStat(path)
		assert.NoError(t, err)
		assert.Equal(t, path, stat.Path)
		assert.Equal(t, real, stat.Real)
		assert.Equal(t, false, stat.IsFile)
		assert.Equal(t, true, stat.Exists)
	}

	// Path: file
	{
		d := requestDotato(real, FirstReq_File, path, SecondReq_File_Eq)
		stat, err := d.newPathStat(path)
		assert.NoError(t, err)
		assert.Equal(t, path, stat.Path)
		assert.Equal(t, path, stat.Real)
		assert.Equal(t, true, stat.IsFile)
		assert.Equal(t, true, stat.Exists)
	}
}

func TestNewPreview(t *testing.T) {
	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: file / Dtt: file
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_File_Eq)
		p, err := d.newPreview(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, dot, p.Dot.Path)
		assert.Equal(t, dot, p.Dot.Real)
		assert.Equal(t, dtt, p.Dtt.Path)
		assert.Equal(t, dtt, p.Dtt.Real)
	}
}

func TestPreviewImportFile(t *testing.T) {
	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: file / Dtt: not exists
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_Empty)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpCreate, p.DttOp)
	}

	// Dot: file / Dtt: file, not equal
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_File_NotEq)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: file / Dtt: file, equal
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_File_Eq)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file / Dtt: symlink, at diff
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_Link_Diff_NotEq)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: file / Dtt: symlink, at same
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_Link_Same)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: not exists
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_Empty)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpCreate, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: file, not equal
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_File_NotEq)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: file, equal
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_File_Eq)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: symlink, diff
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Diff_NotEq)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: symlink, same
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Same)

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}
}

func TestPreviewImportLink(t *testing.T) {
	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)
	
	// Dot: file / Dtt: file, not exists
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_Empty)

		// Preview
		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpCreate, p.DttOp)
	}

	// Dot: file / Dtt: file, not equal
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_File_NotEq)

		// Preview
		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: file / Dtt: file, equal
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_File_Eq)

		// Preview
		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file / Dtt: link, at diff
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_Link_Diff_NotEq)

		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: file / Dtt: link, at same
	{
		d := requestDotato(dot, FirstReq_File, dtt, SecondReq_Link_Same)

		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: link at diff / Dtt: empty
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_Empty)

		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpCreate, p.DttOp)
	}

	// Dot: link at diff / Dtt: file, not equal
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_File_NotEq)

		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: link at diff / Dtt: file, equal
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_File_Eq)
		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link at diff / Dtt: link, diff, not equal
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Diff_NotEq)
		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: link at diff / Dtt: link, diff, equal
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Diff_Eq)
		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: link at diff / Dtt: link, same
	{
		d := requestDotato(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Same)
		p, err := d.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}
}

func TestPreviewExportFile(t *testing.T) {
	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: empty / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Empty)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpCreate, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_File_NotEq)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_File_Eq)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Diff_NotEq)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Diff_Eq)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Same)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: empty / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_Empty)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpCreate, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_File_NotEq)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_File_Eq)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not eq / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Diff_NotEq)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, eq / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Diff_Eq)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Same)
		p, err := d.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}
}

func TestPreviewExportLink(t *testing.T) {
	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: empty / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Empty)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpCreate, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_File_NotEq)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_File_Eq)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Diff_NotEq)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Diff_Eq)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Same)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: empty / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_Empty)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpCreate, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_File_NotEq)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_File_Eq)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not eq / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Diff_NotEq)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, eq / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Diff_Eq)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: link, at diff
	{
		d := requestDotato(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Same)
		p, err := d.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}
}

func TestPreviewUnlink(t *testing.T) {
	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: empty / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Empty)
		p, err := d.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_File_NotEq)
		p, err := d.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_File_Eq)
		p, err := d.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Diff_NotEq)
		p, err := d.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, equal / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Diff_Eq)
		p, err := d.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: file
	{
		d := requestDotato(dtt, FirstReq_File, dot, SecondReq_Link_Same)
		p, err := d.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}
}
