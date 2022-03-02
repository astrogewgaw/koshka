package app

import lip "github.com/charmbracelet/lipgloss"

// A few size constants.
const (
	BS  = 1 // Border size.
	THH = 3 // Height of the table header, with borders.
	TFH = 3 // Height of the table footer, with borders.
	FHH = 7 // Height of the full help menu, with borders.
	SHH = 3 // Height of the short help menu, with borders.
)

// Color scheme.
const (
	C0 = "#000000"
	C1 = "#DF2E2E"
	C2 = "#FFF7AE"
	C3 = "#3C3C3C"
)

var (
	KoshkaStyle = lip.NewStyle().Padding(1, 2, 1, 2)

	TitleStyle = lip.NewStyle().
			Bold(true).
			Padding(0, 1).
			Margin(1, 1, 1, 1).
			Foreground(lip.Color(C2)).
			Background(lip.Color(C1))

	DescStyle = lip.NewStyle().
			Italic(true).
			Margin(0, 0, 1, 0).
			Padding(0, 1, 0, 2).
			Foreground(lip.Color(C2))

	SearchBarStyle = lip.NewStyle().
			BorderStyle(lip.RoundedBorder()).
			BorderForeground(lip.Color(C2))

	SearchTextStyle   = lip.NewStyle()
	SearchPromptStyle = lip.NewStyle()
	SearchCursorStyle = lip.NewStyle()
	ActiveDotStyle    = lip.NewStyle()
	InactiveDotStyle  = lip.NewStyle()
	SpinnerStyle      = lip.NewStyle()

	ConsoleStyle = lip.NewStyle().
			Foreground(lip.Color(C2)).
			BorderForeground(lip.Color(C2)).
			BorderStyle(lip.RoundedBorder())

	TableStyle = lip.NewStyle().
			Foreground(lip.Color(C2)).
			BorderForeground(lip.Color(C2))

	TableColumnStyle = lip.NewStyle().
				Bold(true).
				Align(lip.Center).
				Foreground(lip.Color(C2))

	TableRowNormalStyle = lip.NewStyle().
				Foreground(lip.Color(C2))

	TableRowSelectedStyle = lip.NewStyle().
				Bold(true).
				Background(lip.Color(C1)).
				Foreground(lip.Color(C2))

	TableFooterStyle = lip.NewStyle().
				Bold(true).
				Foreground(lip.Color(C2))

	HelpStyle = lip.NewStyle().
			BorderStyle(lip.RoundedBorder()).
			BorderForeground(lip.Color(C3))

	CommentsStyle = lip.NewStyle().
			Foreground(lip.Color(C2)).
			BorderForeground(lip.Color(C2)).
			BorderStyle(lip.RoundedBorder()).
			SetString(
			"P: Period, in seconds.\n" +
				"DM: Dispersion Measure, in pc per cm³\n" +
				"P1: Period derivative; dimensionless.\n" +
				"GL/GB: Galactic coordinates, in degrees.")

	MoodyBarStyle = lip.NewStyle().
			Foreground(lip.Color(C0)).
			Background(lip.Color(C2))

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
		Height(FHH - 2).
		Render(C.Help.View(C.Paws))
}

func (C Cat) CommentsView() string {
	return CommentsStyle.
		Height(FHH - 2).
		Align(lip.Right).
		Width(C.W - lip.Width(C.HelpView()) - 2).
		String()
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
