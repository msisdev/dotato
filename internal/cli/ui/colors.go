package ui

import (
	"github.com/charmbracelet/lipgloss"
)

/*
https://www.realtimecolors.com/?colors=fcf4e8-090601-e39c16-0c4e73-3d12ab&fonts=Inter-Inter

var (
	PrimaryColor = lipgloss.AdaptiveColor{Light: "${primaryL.hex}", Dark: "${primaryD.hex}"}
	SecondaryColor = lipgloss.AdaptiveColor{Light: "${secondaryL.hex}", Dark: "${secondaryD.hex}"}
	AccentColor = lipgloss.AdaptiveColor{Light: "${accentL.hex}", Dark: "${accentD.hex}"}
	TextColor = lipgloss.AdaptiveColor{Light: "${textL.hex}", Dark: "${textD.hex}"}
	BackgroundColor = lipgloss.AdaptiveColor{Light: "${bgL.hex}", Dark: "${bgD.hex}"}
)
*/

var (
	PrimaryColor 		= lipgloss.AdaptiveColor{Light: "#e9a11c", Dark: "#e39c16"}
	SecondaryColor	= lipgloss.AdaptiveColor{Light: "#8ccdf3", Dark: "#0c4e73"}
	AccentColor 		= lipgloss.AdaptiveColor{Light: "#8054ed", Dark: "#3d12ab"}
	TextColor 			= lipgloss.AdaptiveColor{Light: "#170f03", Dark: "#fcf4e8"}
	BackgroundColor	= lipgloss.AdaptiveColor{Light: "#fefbf6", Dark: "#090601"}
	
	MutedColor 	= lipgloss.Color("7")		// gray
	PositiveColor = lipgloss.Color("10")		// green
	CriticalColor = lipgloss.Color("11")	// olive (yellow)
	NegativeColor = lipgloss.Color("9")		// red
	InfoColor 		= lipgloss.Color("12")		// navy
	PromoteColor = lipgloss.Color("13")	// purple
)
