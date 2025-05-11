package previewprinter

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/msisdev/dotato/internal/cli/app"
)

var (
	iconNone 			= lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render("✔")
	iconSkip 			= lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render("✘")
	iconCreate 		= lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Render("+")
	iconOverwrite	= lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render("!")
	iconUnknown 	= lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("?")
)


func getIcon(op app.FileOp) string {
	switch op {
	case app.FileOpNone:
		return iconNone
	case app.FileOpSkip:
		return iconSkip
	case app.FileOpCreate:
		return iconCreate
	case app.FileOpOverwrite:
		return iconOverwrite
	default:
		return iconUnknown
	}
}

var (
	dotStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	dttStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
	arrowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Bold(true)
)

func sprintPreviewImportFile(p app.Preview) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		getIcon(p.DttOp),
		dotStyle.Render(p.Dot.Path.Abs()),
		arrowStyle.Render("->"),
		dttStyle.Render(p.Dtt.Path.Abs()),
	)
}

func sprintPreviewImportLink(p app.Preview) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		getIcon(p.DotOp),
		dotStyle.Render(p.Dot.Path.Abs()),
		arrowStyle.Render("->"),
		getIcon(p.DttOp),
		dttStyle.Render(p.Dtt.Path.Abs()),
	)
}

func sprintPreviewExportFile(p app.Preview) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		getIcon(p.DotOp),
		dotStyle.Render(p.Dot.Path.Abs()),
		arrowStyle.Render("<-"),
		dttStyle.Render(p.Dtt.Path.Abs()),
	)
}

func sprintPreviewExportLink(p app.Preview) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		getIcon(p.DotOp),
		dotStyle.Render(p.Dot.Path.Abs()),
		arrowStyle.Render("<-"),
		getIcon(p.DttOp),
		dttStyle.Render(p.Dtt.Path.Abs()),
	)
}

func sprintPreviewUnlink(p app.Preview) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		getIcon(p.DotOp),
		dotStyle.Render(p.Dot.Path.Abs()),
		arrowStyle.Render("<-"),
		getIcon(p.DttOp),
		dttStyle.Render(p.Dtt.Path.Abs()),
	)
}
