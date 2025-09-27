package output

import "github.com/charmbracelet/lipgloss/v2"

var (
	green       = lipgloss.Color("#04B575")
	blue        = lipgloss.Color("#00FFFF")
	red         = lipgloss.Color("#D4634C")
	purple      = lipgloss.Color("#9D69FF")
	blueExtra   = lipgloss.Color("#9DA2F5")
	pink        = lipgloss.Color("#F59DF2")
	orange      = lipgloss.Color("#F5A73D")
	whiteRed    = lipgloss.Color("#FFD0C4")
	whiteOrange = lipgloss.Color("#FF9D00")
)

func WhiteOrange(str string) string {
	style := lipgloss.NewStyle().Foreground(whiteOrange)
	return style.Render(str)
}

func WhiteRed(str string) string {
	style := lipgloss.NewStyle().Foreground(whiteRed)
	return style.Render(str)
}

func Orange(str string) string {
	style := lipgloss.NewStyle().Foreground(orange)
	return style.Render(str)
}

func Pink(str string) string {
	style := lipgloss.NewStyle().Foreground(pink)
	return style.Render(str)
}

func Green(str string) string {
	style := lipgloss.NewStyle().Foreground(green)
	return style.Render(str)
}

func Blue(str string) string {
	style := lipgloss.NewStyle().Foreground(blue)
	return style.Render(str)
}

func BlueExtra(str string) string {
	style := lipgloss.NewStyle().Foreground(blueExtra)
	return style.Render(str)
}

func Red(str string) string {
	style := lipgloss.NewStyle().Foreground(red)
	return style.Render(str)
}

func Purple(str string) string {
	style := lipgloss.NewStyle().Foreground(purple)
	return style.Render(str)
}
