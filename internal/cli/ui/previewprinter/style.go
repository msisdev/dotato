package previewprinter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/msisdev/dotato/internal/cli/app"
	"github.com/msisdev/dotato/internal/cli/ui"
)

// Styles
var (
	mutedForegroundStyle    = lipgloss.NewStyle().Foreground(ui.MutedColor)
	mutedBackgroundStyle    = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.MutedColor)
	infoForegroundStyle     = lipgloss.NewStyle().Foreground(ui.InfoColor)
	infoBackgroundStyle     = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.InfoColor)
	positiveForegroundStyle = lipgloss.NewStyle().Foreground(ui.PositiveColor)
	positiveBackgroundStyle = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.PositiveColor)
	criticalForegroundStyle = lipgloss.NewStyle().Foreground(ui.CriticalColor)
	criticalBackgroundStyle = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.CriticalColor)
	negativeForegroundStyle = lipgloss.NewStyle().Foreground(ui.NegativeColor)
	negativeBackgroundStyle = lipgloss.NewStyle().Foreground(ui.EmptyColor).Background(ui.NegativeColor)

	border = lipgloss.Border{
		// Top:         "─",
		// Bottom:      "─",
		Top:         " ",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}
	blockStyle = lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(40).
		BorderForeground(ui.MutedColor).
		Border(border)
)

// https://www.unicode.org/charts/ -> Block Elements
const (
	iconLeftBar     = "▐"
	iconRightBar    = "▌"
	arrowImportFile = "->"
	arrowImportLink = "->"
	arrowExportFile = "<-"
	arrowExportLink = "<-"
	arrowUnlink     = "<->"
)

// Rendered icons
var (
	iconNone = (mutedForegroundStyle.Render(iconLeftBar) +
		mutedBackgroundStyle.Bold(true).Render("o") +
		mutedForegroundStyle.Render(iconRightBar))
	iconSkip = (infoForegroundStyle.Render(iconLeftBar) +
		infoBackgroundStyle.Bold(true).Render("s") +
		infoForegroundStyle.Render(iconRightBar))
	iconCreate = (positiveForegroundStyle.Render(iconLeftBar) +
		positiveBackgroundStyle.Bold(true).Render("c") +
		positiveForegroundStyle.Render(iconRightBar))
	iconOverwrite = (criticalForegroundStyle.Render(iconLeftBar) +
		criticalBackgroundStyle.Bold(true).Render("w") +
		criticalForegroundStyle.Render(iconRightBar))
	iconUnknown = (negativeForegroundStyle.Render(iconLeftBar) +
		negativeBackgroundStyle.Bold(true).Render("?") +
		negativeForegroundStyle.Render(iconRightBar))
)

func renderHeader(updates, total int) string {
	line := fmt.Sprintf("🥔 Preview │ update %d │ total %d", updates, total)
	return blockStyle.Render(line)
}

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
		dotIcon  string
		dotPath  string
		dotMuted bool

		dttIcon  string
		dttPath  string
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
		"%s%s\n %s%s%s",
		dotIcon, dotPath, arrow, dttIcon, dttPath,
	)
}

func renderBlock() string {
	line := fmt.Sprintf(
		"%sokay %sskip %screate %soverwrite",
		iconNone, iconSkip, iconCreate, iconOverwrite,
	)
	return blockStyle.Render(line)
}