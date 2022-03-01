package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Mood int

const (
	Browsing  Mood = 0
	Searching Mood = 1
)

func (M Mood) String() string {
	return [...]string{
		"  Browsing...",
		"🔎 Searching...",
	}[M]
}

func (C *Cat) MoodyPaws() {
	switch C.Mood {
	case Searching:
		C.Finder.Focus()
		C.Table.Focused(false)
		C.Paws.Quit.SetEnabled(false)
		C.Paws.RowUp.SetEnabled(false)
		C.Paws.Search.SetEnabled(false)
		C.Paws.RowDown.SetEnabled(false)
		C.Paws.NextPage.SetEnabled(false)
		C.Paws.PrevPage.SetEnabled(false)
		C.Paws.ToggleFullHelp.SetEnabled(false)
		C.Paws.CancelSearching.SetEnabled(true)
	default:
		C.Finder.Reset()
		C.Finder.Blur()
		C.Table.Focused(true)
		C.Paws.Quit.SetEnabled(true)
		C.Paws.RowUp.SetEnabled(true)
		C.Paws.Search.SetEnabled(true)
		C.Paws.RowDown.SetEnabled(true)
		C.Paws.NextPage.SetEnabled(true)
		C.Paws.PrevPage.SetEnabled(true)
		C.Paws.ClearSearch.SetEnabled(false)
		C.Paws.ToggleFullHelp.SetEnabled(true)
		C.Paws.CancelSearching.SetEnabled(false)
	}
}

func (C *Cat) BrowseMood(msg tea.Msg) tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	C.Table, cmd = C.Table.Update(msg)
	cmds = append(cmds, cmd)
	C.UpdateFooter()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, C.Paws.Quit):
			return tea.Quit
		case key.Matches(msg, C.Paws.ToggleFullHelp):
			C.Help.ShowAll = !C.Help.ShowAll
		case key.Matches(msg, C.Paws.Search):
			C.Mood = Searching
			C.MoodyPaws()
			return textinput.Blink
		case key.Matches(msg, C.Paws.ClearSearch):
			C.Finder.Reset()
		}
	}

	return tea.Batch(cmds...)
}

func (C *Cat) SearchMood(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, C.Paws.CancelSearching):
			C.Mood = Browsing
			C.MoodyPaws()
			C.UpdateTable(AllData())
		}
	}

	if C.Mood == Searching {
		NewFinder, Input := C.Finder.Update(msg)
		InputChanged := C.Finder.Value() != NewFinder.Value()
		C.Finder = NewFinder
		cmds = append(cmds, Input)

		if InputChanged {
			cmds = append(cmds, ApplySearch(C))
		}
	}

	return tea.Batch(cmds...)
}
