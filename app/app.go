package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"golang.org/x/term"

	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"
)

type App struct {
	width  int
	height int
	name   string
	url    string
	desc   string
	author string
	keys   KeyMap
	help   help.Model
}

func Koshka() App {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	return App{
		width:  width,
		height: height,
		name:   "Кошка",
		keys:   DefKeys(),
		help:   DefHelp(),
		author: "Ujjwal Panda",
		desc:   "Meow-ster of cat-alogs!",
		url:    "https://github.com/astrogewgaw/koshka",
	}
}

func (A App) TitleBar() string {
	return lip.JoinVertical(
		lip.Center,
		TitleBarStyle.Render(A.name),
		DescStyle.Render(A.desc))
}
func (A App) DummyView() string {
	VW := A.width
	VH := A.height - lip.Height(A.TitleBar()) - lip.Height(A.HelpView()) - lip.Height(A.StatusBar())

	VW = VW - 2
	VH = VH - 2

	return ViewportStyle.
		Width(VW).
		Height(VH).
		Render(
			lip.Place(
				VW,
				VH,
				lip.Center,
				lip.Center,
				""))
}
func (A App) HelpView() string {
	HW := A.width

	HW = HW - 2

	return HelpStyle.
		Width(HW).
		Render(A.help.View(A.keys))
}
func (A App) StatusBar() string {
	W := lip.Width

	URLView := URLStyle.Render(A.url)
	CopyrightView := CopyrightStyle.Render(fmt.Sprintf("Copyright: %s, 2022", A.author))
	StatusView := StatusStyle.Width(A.width - W(URLView) - W(CopyrightView)).Render("")

	return StatusBarStyle.Width(A.width).Render(
		lip.JoinHorizontal(
			lip.Top,
			CopyrightView,
			StatusView,
			URLView))
}
func (A App) Init() tea.Cmd { return tea.EnterAltScreen }
func (A App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		A.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, A.keys.ForceQuit):
			return A, tea.Quit
		case key.Matches(msg, A.keys.Quit):
			return A, tea.Quit
		case key.Matches(msg, A.keys.ShowFullHelp):
			A.help.ShowAll = !A.help.ShowAll
		case key.Matches(msg, A.keys.CursorUp):
		case key.Matches(msg, A.keys.CursorDown):
		case key.Matches(msg, A.keys.PrevPage):
		case key.Matches(msg, A.keys.NextPage):
		case key.Matches(msg, A.keys.GoToStart):
		case key.Matches(msg, A.keys.GoToEnd):
		case key.Matches(msg, A.keys.Filter):
		case key.Matches(msg, A.keys.ClearFilter):
		case key.Matches(msg, A.keys.CancelWhileFiltering):
		case key.Matches(msg, A.keys.AcceptWhileFiltering):
		}
	}
	return A, nil
}
func (A App) View() string {
	KoshkaView := lip.JoinVertical(
		lip.Center,
		A.TitleBar(),
		A.DummyView(),
		A.HelpView(),
		A.StatusBar())
	return KoshkaView
}

func RunApp() {
	if os.Getenv("HELP_DEBUG") != "" {
		if f, err := tea.LogToFile("debug.log", "help"); err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		} else {
			defer f.Close()
		}
	}

	if err := tea.NewProgram(Koshka()).Start(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}
}
