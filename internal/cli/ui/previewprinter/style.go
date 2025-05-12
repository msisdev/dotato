package previewprinter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/ui"
)

var (
	mutedForegroundStyle = lipgloss.NewStyle().Foreground(ui.MutedColor)
	mutedBackgroundStyle = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.MutedColor)
	infoForegroundStyle = lipgloss.NewStyle().Foreground(ui.InfoColor)
	infoBackgroundStyle = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.InfoColor)
	positiveForegroundStyle = lipgloss.NewStyle().Foreground(ui.PositiveColor)
	positiveBackgroundStyle = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.PositiveColor)
	criticalForegroundStyle = lipgloss.NewStyle().Foreground(ui.CriticalColor)
	criticalBackgroundStyle = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.CriticalColor)
	negativeForegroundStyle = lipgloss.NewStyle().Foreground(ui.NegativeColor)
	negativeBackgroundStyle = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.NegativeColor)
)

// https://www.unicode.org/charts/ -> Block Elements
const (
	iconLeftBar = "▐"
	iconRightBar = "▌"
	arrowImportFile = "->"
	arrowImportLink = "->"
	arrowExportFile = "<-"
	arrowExportLink = "<-"
	arrowUnlink     = "<->"
)

var (
	iconNone = (
		mutedForegroundStyle.Render(iconLeftBar) +
		mutedBackgroundStyle.Bold(true).Render("o") +
		mutedForegroundStyle.Render(iconRightBar))
	iconSkip 			= (
		infoForegroundStyle.Render(iconLeftBar) +
		infoBackgroundStyle.Bold(true).Render("s") +
		infoForegroundStyle.Render(iconRightBar))
	iconCreate = (
		positiveForegroundStyle.Render(iconLeftBar) +
		positiveBackgroundStyle.Bold(true).Render("c") +
		positiveForegroundStyle.Render(iconRightBar))
	iconOverwrite = (
		criticalForegroundStyle.Render(iconLeftBar) +
		criticalBackgroundStyle.Bold(true).Render("w") +
		criticalForegroundStyle.Render(iconRightBar))
	iconUnknown = (
		negativeForegroundStyle.Render(iconLeftBar) +
		negativeBackgroundStyle.Bold(true).Render("?") +
		negativeForegroundStyle.Render(iconRightBar))
	footer = fmt.Sprintf(
		"%s okay / %s skip / %s create / %s overwrite\n",
		iconNone, iconSkip, iconCreate, iconOverwrite,
	)
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

func renderItem(p app.Preview, arrow string) string {
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
		dotPath = mutedForegroundStyle.Render(p.Dot.Path.Abs())
		dttPath = mutedForegroundStyle.Render(p.Dtt.Path.Abs())
		arrow = mutedForegroundStyle.Render(arrow)
	} else {
		dotPath = p.Dot.Path.Abs()
		dttPath = p.Dtt.Path.Abs()
	}

	return fmt.Sprintf(
		"%s%s\n %s%s%s\n",
		dotIcon, dotPath, arrow, dttIcon, dttPath,
	)
}
