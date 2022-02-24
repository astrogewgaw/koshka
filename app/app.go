package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/evertras/bubble-table/table"
	"golang.org/x/term"

	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"
)

type App struct {
	Width  int
	Height int

	Status string

	Keys      KeyMap
	Help      help.Model
	Table     table.Model
	SearchBar textinput.Model
}

func Koshka() App {
	Width, Height, _ := term.GetSize(int(os.Stdout.Fd()))
	return App{
		Width:  Width,
		Height: Height,

		Status: "",

		Keys:      DefKeys(),
		Help:      DefHelp(),
		Table:     DefTable(),
		SearchBar: DefSearchBar(),
	}
}

func (A App) Init() tea.Cmd { return nil }
func (A App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	A.Table, cmd = A.Table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		A.Width = msg.Width
		A.Height = msg.Height
		A.Help.Width = msg.Width
		A.Table = A.Table.WithPageSize(
			msg.Height -
				lip.Height(A.TitleBar()) -
				lip.Height(A.SearchView()) -
				lip.Height(A.HelpView()) -
				lip.Height(A.StatusBar(A.Status)) -
				HeaderSize -
				FooterSize -
				(FullHelpSize - ShortHelpSize))
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, A.Keys.ForceQuit):
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, A.Keys.Quit):
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, A.Keys.ShowFullHelp):
			A.Help.ShowAll = !A.Help.ShowAll
		case key.Matches(msg, A.Keys.Filter):
			A.SearchBar.Focus()
			A.Table.Focused(false)
			A.Status = "Filtering"
			A.Keys.Filter.SetEnabled(false)
			A.Keys.Quit.SetEnabled(false)
			A.Keys.ShowFullHelp.SetEnabled(false)
			A.Keys.CloseFullHelp.SetEnabled(false)
		case key.Matches(msg, A.Keys.ClearFilter):
			A.SearchBar.Blur()
			A.Table.Focused(true)
			A.Status = "Browsing"
			A.Keys.Filter.SetEnabled(true)
			A.Keys.Quit.SetEnabled(true)
			A.Keys.ShowFullHelp.SetEnabled(true)
			A.Keys.CloseFullHelp.SetEnabled(true)
		case key.Matches(msg, A.Keys.CancelWhileFiltering):
			A.SearchBar.Blur()
			A.Table.Focused(true)
			A.Status = "Browsing"
			A.Keys.Filter.SetEnabled(true)
			A.Keys.Quit.SetEnabled(true)
			A.Keys.ShowFullHelp.SetEnabled(true)
			A.Keys.CloseFullHelp.SetEnabled(true)
		case key.Matches(msg, A.Keys.AcceptWhileFiltering):
		}
	case SearchMsg:
		A.Table = A.Table.WithRows(msg)
		return A, nil
	}

	if A.Status == "Filtering" {
		NewSearchBar, Input := A.SearchBar.Update(msg)
		InputChanged := A.SearchBar.Value() != NewSearchBar.Value()
		A.SearchBar = NewSearchBar
		cmds = append(cmds, Input)

		if InputChanged {
			cmds = append(cmds, Searching(A))
		}
	}

	return A, tea.Batch(cmds...)
}
func (A App) View() string {
	KoshkaView := lip.JoinVertical(
		lip.Left,
		A.TitleBar(),
		A.SearchView(),
		A.TableView(),
		A.HelpView(),
		A.StatusBar(A.Status))
	return KoshkaView
}

func RunApp() {
	if os.Getenv("KOSHKA_DEBUG") != "" {
		if f, err := tea.LogToFile("koshka.log", "help"); err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		} else {
			defer f.Close()
		}
	}

	if err := tea.NewProgram(
		Koshka(),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	).Start(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}
}
