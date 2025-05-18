package versioncmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/msisdev/dotato/internal/cli/ui"
)

var (
	currentStyle = lipgloss.NewStyle().Foreground(ui.SecondaryColor)
	latestStyle  = lipgloss.NewStyle().Foreground(ui.PrimaryColor)
)

func renderCurrent(current string) string {
	return currentStyle.Render(current)
}

func renderCurrentIsLatest(current string) string {
	return fmt.Sprintf("☀️ current version %s is latest", latestStyle.Render(current))
}

func renderLatestAvailable(current, latest string) string {
	return fmt.Sprintf("🥔 version upgrade %s -> %s is available", currentStyle.Render(current), latestStyle.Render(latest))
}

func renderMajorUpgradeWarning() string {
	return "⚠️ This is a major release, please check the changelog before upgrading."
}
