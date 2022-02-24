package app

import "fmt"
import lip "github.com/charmbracelet/lipgloss"

const (
	C0 = "#000000"
	C1 = "#FF6600"
	C2 = "#990000"
	C3 = "#F7F7CF"
	C4 = "#3C3C3C"
)

var (
	KoshkaStyle = lip.NewStyle().Padding(1, 2, 1, 2)

	TitleStyle = lip.NewStyle().
			Bold(true).
			Padding(0, 1).
			Margin(1, 1, 1, 1).
			Foreground(lip.Color(C3)).
			Background(lip.Color(C2))

	DescStyle = lip.NewStyle().
			Italic(true).
			Margin(0, 0, 1, 0).
			Padding(0, 1, 0, 2).
			Foreground(lip.Color(C3))

	SearchBarStyle = lip.NewStyle().
			BorderStyle(lip.RoundedBorder()).
			BorderForeground(lip.Color(C3))

	SearchTextStyle   = lip.NewStyle()
	SearchPromptStyle = lip.NewStyle()
	SearchCursorStyle = lip.NewStyle()
	ActiveDotStyle    = lip.NewStyle()
	InactiveDotStyle  = lip.NewStyle()
	SpinnerStyle      = lip.NewStyle()

	TableStyle = lip.NewStyle().Foreground(lip.Color(C3))

	TableColumnStyle = lip.NewStyle().
				Inherit(TableStyle).
				Bold(true).
				Align(lip.Center)

	TableRowNormalStyle = lip.NewStyle().Inherit(TableStyle)

	TableRowSelectedStyle = lip.NewStyle().
				Inherit(TableStyle).
				Bold(true).
				Background(lip.Color(C2))

	HelpStyle = lip.NewStyle().
			BorderStyle(lip.RoundedBorder()).
			BorderForeground(lip.Color(C4))

	StatusBarStyle = lip.NewStyle().
			Foreground(lip.Color(C0)).
			Background(lip.Color(C3))

	URLStyle = lip.NewStyle().
			Padding(0, 2).
			Italic(true).
			Align(lip.Right).
			Foreground(lip.Color(C3)).
			Background(lip.Color(C0)).
			SetString(fmt.Sprintf("%s %s", "🌏", URL))

	CopyrightStyle = lip.NewStyle().
			Bold(true).
			Padding(0, 1).
			Foreground(lip.Color(C3)).
			Background(lip.Color(C2)).
			SetString(fmt.Sprintf("  %s, 2022", Author))

	StatusStyle = lip.NewStyle().
			Inherit(StatusBarStyle).
			Italic(true).
			Padding(0, 1).
			Align(lip.Center)
)
