package app

import (
	"testing"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func TestNewPreview(t *testing.T) {
	// assertOS(t)

	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: file / Dtt: file
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_File_Eq)
		p, err := app.newPreview(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, dot, p.Dot.Path)
		assert.Equal(t, dot, p.Dot.Target)
		assert.Equal(t, dot, p.Dot.Real)
		assert.Equal(t, dtt, p.Dtt.Path)
		assert.Equal(t, dtt, p.Dtt.Target)
		assert.Equal(t, dtt, p.Dtt.Real)
	}
}

func TestPreviewImportFile(t *testing.T) {
	assertOS(t)

	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: file / Dtt: not exists
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_Empty)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpCreate, p.DttOp)
	}

	// Dot: file / Dtt: file, not equal
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_File_NotEq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: file / Dtt: file, equal
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_File_Eq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file / Dtt: symlink, at diff
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: file / Dtt: symlink, at same
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_Link_Same)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: not exists
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_Empty)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpCreate, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: file, not equal
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_File_NotEq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: file, equal
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_File_Eq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: symlink, diff
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: symlink at diff / Dtt: symlink, same
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Same)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: symlink at same / Dtt: not exists
	{
		app := requestApp(dot, FirstReq_Link_Same, dtt, SecondReq_Empty)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpSkip, p.DttOp)
	}

	// Dot: symlink at same / Dtt: file, not equal
	{
		app := requestApp(dot, FirstReq_Link_Same, dtt, SecondReq_File_NotEq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpSkip, p.DttOp)
	}

	// Dot: symlink at same / Dtt: file, equal
	{
		app := requestApp(dot, FirstReq_Link_Same, dtt, SecondReq_File_Eq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpSkip, p.DttOp)
	}

	// Dot: symlink at same / Dtt: symlink, diff
	{
		app := requestApp(dot, FirstReq_Link_Same, dtt, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpSkip, p.DttOp)
	}

	// Dot: symlink at same / Dtt: symlink, same
	{
		app := requestApp(dot, FirstReq_Link_Same, dtt, SecondReq_Link_Same)
		_, err := app.PreviewImportFile(dot, dtt)
		assert.Error(t, err)
	}
}

func TestPreviewImportLink(t *testing.T) {
	assertOS(t)

	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)
	
	// Dot: file / Dtt: file, not exists
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_Empty)
		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpCreate, p.DttOp)
	}

	// Dot: file / Dtt: file, not equal
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_File_NotEq)
		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: file / Dtt: file, equal
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_File_Eq)
		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file / Dtt: link, at diff
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_Link_Diff_NotEq)

		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: file / Dtt: link, at same
	{
		app := requestApp(dot, FirstReq_File, dtt, SecondReq_Link_Same)

		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: link at diff / Dtt: empty
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_Empty)

		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpCreate, p.DttOp)
	}

	// Dot: link at diff / Dtt: file, not equal
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_File_NotEq)

		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: link at diff / Dtt: file, equal
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_File_Eq)
		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link at diff / Dtt: link, diff, not equal
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: link at diff / Dtt: link, diff, equal
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Diff_Eq)
		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}

	// Dot: link at diff / Dtt: link, same
	{
		app := requestApp(dot, FirstReq_Link_Diff, dtt, SecondReq_Link_Same)
		p, err := app.PreviewImportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpOverwrite, p.DttOp)
	}
}

func TestPreviewExportFile(t *testing.T) {
	assertOS(t)

	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: empty / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Empty)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpCreate, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_File_NotEq)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_File_Eq)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Diff_Eq)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Same)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: empty / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_Empty)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpCreate, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_File_NotEq)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_File_Eq)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not eq / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, eq / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Diff_Eq)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Same)
		p, err := app.PreviewExportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}
}

func TestPreviewExportLink(t *testing.T) {
	assertOS(t)

	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: empty / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Empty)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpCreate, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_File_NotEq)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_File_Eq)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Diff_Eq)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Same)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: empty / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_Empty)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpCreate, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_File_NotEq)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_File_Eq)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not eq / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, eq / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Diff_Eq)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: link, at diff
	{
		app := requestApp(dtt, FirstReq_Link_Diff, dot, SecondReq_Link_Same)
		p, err := app.PreviewExportLink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}
}

func TestPreviewUnlink(t *testing.T) {
	assertOS(t)

	var (
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
	)

	// Dot: empty / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Empty)
		p, err := app.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, not equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_File_NotEq)
		p, err := app.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: file, equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_File_Eq)
		p, err := app.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, not equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Diff_NotEq)
		p, err := app.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at diff, equal / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Diff_Eq)
		p, err := app.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpNone, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}

	// Dot: link, at same / Dtt: file
	{
		app := requestApp(dtt, FirstReq_File, dot, SecondReq_Link_Same)
		p, err := app.PreviewUnlink(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, FileOpOverwrite, p.DotOp)
		assert.Equal(t, FileOpNone, p.DttOp)
	}
}
