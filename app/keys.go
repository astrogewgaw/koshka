package app

import "github.com/charmbracelet/bubbles/key"
import "github.com/charmbracelet/bubbles/help"

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
		CursorUp:             key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "Up")),
		CursorDown:           key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/k", "Down")),
		NextPage:             key.NewBinding(key.WithKeys("left", "h", "pgup"), key.WithHelp("←/h/PgUp", "Prev Page")),
		PrevPage:             key.NewBinding(key.WithKeys("right", "l", "pgdown"), key.WithHelp("→/l/PgDn", "Next Page")),
		GoToStart:            key.NewBinding(key.WithKeys("home", "g"), key.WithHelp("g/Home", "Go to start.")),
		GoToEnd:              key.NewBinding(key.WithKeys("end", "G"), key.WithHelp("G/End", "Go to end.")),
		Filter:               key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "Search the database.")),
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
