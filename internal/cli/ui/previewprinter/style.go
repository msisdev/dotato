package previewprinter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/ui"
)

var (
	mutedStyle = lipgloss.NewStyle().Foreground(ui.MutedColor)
)

var (
	iconNone 			= lipgloss.NewStyle().Foreground(ui.MutedColor).Render("✔")
	iconSkip 			= lipgloss.NewStyle().Foreground(ui.InfoColor).Render("✘")
	iconCreate 		= lipgloss.NewStyle().Foreground(ui.PositiveColor).Render("+")
	iconOverwrite	= lipgloss.NewStyle().Foreground(ui.CriticalColor).Render("!")
	iconUnknown 	= lipgloss.NewStyle().Foreground(ui.NegativeColor).Render("?")
)

func renderIcon(op app.FileOp) (string, bool) {
	switch op {
	case app.FileOpNone:
		return iconNone, true
	case app.FileOpSkip:
		return iconSkip, false
	case app.FileOpCreate:
		return iconCreate, false
	case app.FileOpOverwrite:
		return iconOverwrite, false
	default:
		return iconUnknown, false
	}
}

const (
	arrowImportFile = "->"
	arrowImportLink = "->"
	arrowExportFile = "<-"
	arrowExportLink = "<-"
	arrowUnlink     = "<->"
)

func render(p app.Preview, arrow string) string {
	var (
		dotIcon string
		dotPath string
		dotMuted bool

		dttIcon string
		dttPath string
		dttMuted bool
	)
	// Get muted status
	dotIcon, dotMuted = renderIcon(p.DotOp)
	dttIcon, dttMuted = renderIcon(p.DttOp)

	// If both are muted, render all in muted style
	if dotMuted && dttMuted {
		dotPath = mutedStyle.Render(p.Dot.Path.Abs())
		dttPath = mutedStyle.Render(p.Dtt.Path.Abs())
		arrow = mutedStyle.Render(arrow)
	} else {
		dotPath = p.Dot.Path.Abs()
		dttPath = p.Dtt.Path.Abs()
	}

	return fmt.Sprintf(
		"%s %s\n%s %s %s\n",
		dotIcon, dotPath, arrow, dttIcon, dttPath,
	)
}
