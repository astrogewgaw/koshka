package app

import (
	"fmt"

	"github.com/astrogewgaw/koshka/data"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/evertras/bubble-table/table"

	tea "github.com/charmbracelet/bubbletea"
)

type SearchMsg []table.Row

type KeyMap struct {
	CursorUp             key.Binding
	CursorDown           key.Binding
	NextPage             key.Binding
	PrevPage             key.Binding
	GoToStart            key.Binding
	GoToEnd              key.Binding
	Filter               key.Binding
	ClearFilter          key.Binding
	CancelWhileFiltering key.Binding
	AcceptWhileFiltering key.Binding
	ShowFullHelp         key.Binding
	CloseFullHelp        key.Binding
	Quit                 key.Binding
	ForceQuit            key.Binding
}

func DefKeys() KeyMap {
	return KeyMap{
		CursorUp:             key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Up")),
		CursorDown:           key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Down")),
		NextPage:             key.NewBinding(key.WithKeys("left", "pgup"), key.WithHelp("←/PgUp", "Prev Page")),
		PrevPage:             key.NewBinding(key.WithKeys("right", "pgdown"), key.WithHelp("→/PgDn", "Next Page")),
		GoToStart:            key.NewBinding(key.WithKeys("home"), key.WithHelp("Home", "Go to start.")),
		GoToEnd:              key.NewBinding(key.WithKeys("end"), key.WithHelp("End", "Go to end.")),
		Filter:               key.NewBinding(key.WithKeys("ctrl+f"), key.WithHelp("Ctrl+F", "Search the database.")),
		ClearFilter:          key.NewBinding(key.WithKeys("esc"), key.WithHelp("Esc", "Clear search.")),
		CancelWhileFiltering: key.NewBinding(key.WithKeys("esc"), key.WithHelp("Esc", "Cancel search.")),
		AcceptWhileFiltering: key.NewBinding(key.WithKeys("enter"), key.WithHelp("Enter", "Get searching!")),
		ShowFullHelp:         key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "Show full help.")),
		CloseFullHelp:        key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "Close full help.")),
		Quit:                 key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "Quit")),
		ForceQuit:            key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("Ctrl+C", "Force quit.")),
	}
}

func DefHelp() help.Model { return help.New() }

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.ShowFullHelp, k.Quit, k.ForceQuit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.CursorUp, k.CursorDown},
		{k.NextPage, k.PrevPage, k.GoToStart, k.GoToEnd},
		{k.Filter, k.ClearFilter, k.CancelWhileFiltering, k.AcceptWhileFiltering},
		{k.ShowFullHelp, k.CloseFullHelp, k.Quit, k.ForceQuit},
	}
}

func DefSpinner() spinner.Model {
	S := spinner.New()
	S.Style = SpinnerStyle
	S.Spinner = spinner.Dot
	return S
}

func DefSearchBar() textinput.Model {
	T := textinput.New()
	T.Prompt = "❯ "
	T.CharLimit = 50
	T.TextStyle = SearchTextStyle
	T.PromptStyle = SearchPromptStyle
	T.CursorStyle = SearchCursorStyle
	return T
}

func DefTable() table.Model {
	cols := []table.Column{
		table.NewColumn("PSRJ", "Name", 15),
		table.NewColumn("P0", "Period (in s)", 15),
		table.NewColumn("DM", "DM (in pc per cm^-3)", 25),
	}

	keys := table.DefaultKeyMap()
	keys.RowUp.SetKeys("up")
	keys.RowDown.SetKeys("down")
	keys.PageUp.SetKeys("left", "pgup")
	keys.PageDown.SetKeys("right", "pgdown")

	return table.New(cols).
		Focused(true).
		WithKeyMap(keys).
		WithPageSize(10).
		WithRows(MakeRows("")).
		HeaderStyle(TableColumnStyle).
		HighlightStyle(TableRowSelectedStyle)
}

func MakeRows(Input string) []table.Row {
	rows := []table.Row{}

	pulsars := data.GetPulsars(Input)

	for _, pulsar := range pulsars {
		rows = append(
			rows,
			table.NewRow(
				table.RowData{
					"PSRJ": pulsar.PSRJ.String,
					"P0":   fmt.Sprintf("%f", pulsar.P0.Float64),
					"DM":   fmt.Sprintf("%f", pulsar.DM.Float64),
				}).WithStyle(TableRowNormalStyle))
	}

	return rows
}

func Searching(A App) tea.Cmd {
	return func() tea.Msg {
		rows := MakeRows(A.SearchBar.Value())
		return SearchMsg(rows)
	}
}
