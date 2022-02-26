package app

import lip "github.com/charmbracelet/lipgloss"

// A few size constants.
const (
	BS  = 1 // Border size.
	THH = 3 // Height of the table header, with borders.
	TFH = 3 // Height of the table footer, with borders.
	FHH = 6 // Height of the full help menu, with borders.
	SHH = 3 // Height of the short help menu, with borders.
)

// Color scheme.
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

	ConsoleStyle = lip.NewStyle().
			Foreground(lip.Color(C3)).
			BorderForeground(lip.Color(C3)).
			BorderStyle(lip.RoundedBorder())

	TableColumnStyle = lip.NewStyle().
				Bold(true).
				Align(lip.Center).
				Foreground(lip.Color(C3))

	TableRowNormalStyle = lip.NewStyle().
				Foreground(lip.Color(C3))

	TableRowSelectedStyle = lip.NewStyle().
				Bold(true).
				Background(lip.Color(C2)).
				Foreground(lip.Color(C3))

	TableFooterStyle = lip.NewStyle().
				Bold(true).
				Foreground(lip.Color(C3))

	HelpStyle = lip.NewStyle().
			BorderStyle(lip.RoundedBorder()).
			BorderForeground(lip.Color(C4))

	MoodyBarStyle = lip.NewStyle().
			Foreground(lip.Color(C0)).
			Background(lip.Color(C3))

	MoodyStyle = lip.NewStyle().
			Inherit(MoodyBarStyle).
			Italic(true).
			Padding(0, 1).
			Align(lip.Center)
)

func (C Cat) TitleBar() string { return TitleStyle.Render(C.Name) }

func (C Cat) TableView() string {
	var txt string

	W := C.W - (2 * BS)
	H := C.H -
		lip.Height(C.TitleBar()) -
		lip.Height(C.HelpView()) -
		lip.Height(C.FinderView()) -
		lip.Height(C.MoodyBar()) -
		(2 * BS)

	if C.Table.TotalRows() != 0 {
		txt = C.Table.View()
	} else {
		txt = "No pulsars found."
	}

	return ConsoleStyle.Render(
		lip.Place(
			W,
			H,
			lip.Center,
			lip.Center,
			txt))
}

func (C Cat) FinderView() string {
	return SearchBarStyle.
		Width(C.W - 2).
		Render(C.Finder.View())
}

func (C Cat) HelpView() string {
	return HelpStyle.
		Width(C.W - 2).
		Height(FHH - 2).
		Render(C.Help.View(C.Paws))
}

func (C Cat) MoodyBar() string {
	return MoodyBarStyle.
		Width(C.W).
		Render(
			lip.JoinHorizontal(
				lip.Top,
				MoodyStyle.
					Width(C.W).
					Render(C.Mood.String())))
}
