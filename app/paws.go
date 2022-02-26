package app

import "github.com/charmbracelet/bubbles/key"

type Paws struct {
	Quit            key.Binding
	ForceQuit       key.Binding
	RowUp           key.Binding
	RowDown         key.Binding
	NextPage        key.Binding
	PrevPage        key.Binding
	Search          key.Binding
	ClearSearch     key.Binding
	ToggleFullHelp  key.Binding
	CancelSearching key.Binding
	AcceptSearching key.Binding
}

func (P Paws) ShortHelp() []key.Binding {
	return []key.Binding{P.ToggleFullHelp, P.Quit, P.ForceQuit}
}

func (P Paws) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{P.ToggleFullHelp, P.Quit, P.ForceQuit},
		{P.RowUp, P.RowDown, P.NextPage, P.PrevPage},
		{P.Search, P.ClearSearch, P.CancelSearching, P.AcceptSearching},
	}
}
