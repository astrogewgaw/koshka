package app

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/evertras/bubble-table/table"
	"golang.org/x/term"

	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"
)

func CreateFinder() textinput.Model {
	tx := textinput.New()
	tx.Prompt = "❯ "
	tx.CharLimit = 50
	return tx
}

func CreatePaws() Paws {
	return Paws{
		Quit:            key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "Quit.")),
		ForceQuit:       key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("Ctrl+C", "Force quit.")),
		RowUp:           key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Up")),
		RowDown:         key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Down")),
		NextPage:        key.NewBinding(key.WithKeys("left", "pgup"), key.WithHelp("←/PgUp", "Prev Page")),
		PrevPage:        key.NewBinding(key.WithKeys("right", "pgdown"), key.WithHelp("→/PgDn", "Next Page")),
		Search:          key.NewBinding(key.WithKeys("ctrl+f"), key.WithHelp("Ctrl+F", "Search the cat-alogs!")),
		ClearSearch:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("Esc", "Clear search.")),
		ToggleFullHelp:  key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "Toggle full help.")),
		CancelSearching: key.NewBinding(key.WithKeys("esc"), key.WithHelp("Esc", "Cancel search.")),
	}
}

// The application model.
type Cat struct {
	W int
	H int

	Name    string
	Author  string
	Website string

	Mood   Mood
	Paws   Paws
	Help   help.Model
	Table  table.Model
	Finder textinput.Model
}

// The default application state.
func Кошка() Cat {
	W, H, _ := term.GetSize(int(os.Stdout.Fd()))
	return Cat{
		W: W,
		H: H,

		Name:    "Кошка",
		Author:  "Ujjwal Panda",
		Website: "https://github.com/astrogewgaw/koshka",

		Mood:   Browsing,
		Help:   help.New(),
		Paws:   CreatePaws(),
		Finder: CreateFinder(),
		Table:  CreateFilledTable(),
	}
}

func (C Cat) Init() tea.Cmd { return nil }

func (C Cat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	C.MoodyPaws()
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		C.W = msg.Width
		C.H = msg.Height
		C.Help.Width = msg.Width
		C.Table = C.Table.WithPageSize(
			C.H -
				lip.Height(C.TitleBar()) -
				lip.Height(C.FinderView()) -
				lip.Height(C.HelpView()) -
				lip.Height(C.MoodyBar()) -
				THH - TFH - (FHH - SHH))

	case tea.KeyMsg:
		if key.Matches(msg, C.Paws.ForceQuit) {
			return C, tea.Quit
		}
	case SearchMsg:
		C.UpdateTable(msg)
		return C, nil
	}

	if C.Mood == Browsing {
		cmds = append(cmds, C.BrowseMood(msg))
	} else {
		cmds = append(cmds, C.SearchMood(msg))
	}

	return C, tea.Batch(cmds...)
}

func (C Cat) View() string {
	return lip.JoinVertical(
		lip.Left,
		C.TitleBar(),
		C.FinderView(),
		C.TableView(),
		lip.JoinHorizontal(
			lip.Bottom,
			C.HelpView(),
			C.UnitsView(),
		),
		C.MoodyBar())
}
