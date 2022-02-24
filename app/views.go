package app

import lip "github.com/charmbracelet/lipgloss"

const (
	Name   = "Кошка"
	Author = "Ujjwal Panda"
	URL    = "https://github.com/astrogewgaw/koshka"

	HeaderSize    = 3
	FooterSize    = 3
	FullHelpSize  = 6
	ShortHelpSize = 3
)

func (A App) TitleBar() string { return TitleStyle.Render(Name) }

func (A App) TableView() string {
	W := A.Width - 2
	H := A.Height -
		lip.Height(A.TitleBar()) -
		lip.Height(A.HelpView()) -
		lip.Height(A.SearchView()) -
		lip.Height(A.StatusBar(A.Status))

	return TableStyle.Render(
		lip.Place(
			W,
			H,
			lip.Center,
			lip.Center,
			A.Table.View()))
}

func (A App) SearchView() string {
	return SearchBarStyle.
		Width(A.Width - 2).
		Render(A.SearchBar.View())
}

func (A App) HelpView() string {
	return HelpStyle.
		Width(A.Width - 2).
		Height(FullHelpSize - 2).
		Render(A.Help.View(A.Keys))
}

func (A App) StatusBar(Status string) string {
	return StatusBarStyle.
		Width(A.Width).
		Render(
			lip.JoinHorizontal(
				lip.Top,
				CopyrightStyle.String(),
				StatusStyle.
					Width(
						A.Width-
							lip.Width(CopyrightStyle.String())-
							lip.Width(URLStyle.String())).
					Render(Status),
				URLStyle.String()))
}
