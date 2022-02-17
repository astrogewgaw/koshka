package app

import lip "github.com/charmbracelet/lipgloss"

const (
	C0  = "#000000"
	C1  = "#03071e"
	C2  = "#370617"
	C3  = "#6a040f"
	C4  = "#9d0208"
	C5  = "#d00000"
	C6  = "#dc2f02"
	C7  = "#e85d04"
	C8  = "#f48c06"
	C9  = "#faa307"
	C10 = "#ffba08"
	C11 = "#e9d8a6"
	C12 = "#9b2226"
	C13 = "#3C3C3C"
)

var (
	KoshkaStyle = lip.NewStyle().
			Padding(1, 2, 1, 2).
			Foreground(lip.Color(C8))

	TitleBarStyle = lip.NewStyle().
			Italic(true).
			Padding(0, 1).
			Margin(1, 0, 1, 0).
			Foreground(lip.Color(C11)).
			Background(lip.Color(C2))

	DescStyle = lip.NewStyle().
			Italic(true).
			Padding(0, 0, 1, 0).
			Foreground(lip.Color(C11))

	SearchBarStyle = lip.NewStyle().
			BorderStyle(lip.RoundedBorder()).
			BorderForeground(lip.Color(C11))

	ViewportStyle = lip.NewStyle().
			BorderStyle(lip.RoundedBorder()).
			BorderForeground(lip.Color(C11))

	TableViewStyle        = lip.NewStyle()
	TableRowNormalStyle   = lip.NewStyle()
	TableRowSelectedStyle = lip.NewStyle()

	HelpStyle = lip.NewStyle().
			BorderStyle(lip.RoundedBorder()).
			BorderForeground(lip.Color(C13))

	StatusBarStyle = lip.NewStyle().
			Foreground(lip.Color(C11)).
			Background(lip.Color(C4))

	URLStyle = lip.NewStyle().
			Italic(true).
			Padding(0, 2).
			Align(lip.Right).
			Foreground(lip.Color(C0)).
			Background(lip.Color(C11))

	CopyrightStyle = lip.NewStyle().
			Bold(true).
			Italic(true).
			Padding(0, 1).
			Foreground(lip.Color(C11)).
			Background(lip.Color(C1))

	StatusStyle = lip.NewStyle().
			Inherit(StatusBarStyle).
			Italic(true).
			Padding(0, 1).
			Align(lip.Center)
)
