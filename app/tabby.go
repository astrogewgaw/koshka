package app

import (
	"fmt"
	"github.com/astrogewgaw/koshka/data"
	"github.com/evertras/bubble-table/table"

	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"
)

type SearchMsg []table.Row

func DataIntoTable(search string) []table.Row {
	var rows []table.Row
	psrs := data.GetPulsars(search)
	for _, psr := range psrs {
		rows = append(
			rows,
			table.NewRow(
				table.RowData{
					"PSRJ": psr.PSRJ.String,
					"P0":   fmt.Sprintf("%f", psr.P0.Float64),
					"P1":   fmt.Sprintf("%f", psr.P1.Float64),
					"DM":   fmt.Sprintf("%f", psr.DM.Float64),
					"GL":   fmt.Sprintf("%f", psr.GL.Float64),
					"GB":   fmt.Sprintf("%f", psr.GB.Float64),
				}).WithStyle(TableRowNormalStyle))
	}
	return rows
}

func AllData() []table.Row { return DataIntoTable("") }

func ApplySearch(C *Cat) tea.Cmd {
	str := C.Finder.Value()
	msg := SearchMsg(DataIntoTable(str))
	cmd := func() tea.Msg { return msg }
	return cmd
}

func CreateEmptyTable() table.Model {
	cols := []table.Column{
		table.NewColumn("PSRJ", "Name", 15),
		table.NewColumn("P0", "P0", 15),
		table.NewColumn("P1", "P1", 15),
		table.NewColumn("DM", "DM", 15),
		table.NewColumn("GL", "GL", 15),
		table.NewColumn("GB", "GB", 15),
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
		WithBaseStyle(TableStyle).
		HeaderStyle(TableColumnStyle).
		HighlightStyle(TableRowSelectedStyle).
		SortByAsc("PSRJ")
}

func CreateFilledTable() table.Model { return CreateEmptyTable().WithRows(AllData()) }

func (C *Cat) UpdateFooter() {
	C.Table = C.Table.WithStaticFooter(
		TableFooterStyle.Render(
			fmt.Sprintf(
				"Pg. %d of %d | Total number of pulsars: %d.",
				C.Table.CurrentPage(),
				C.Table.MaxPages(),
				C.Table.TotalRows())))
}

func (C *Cat) UpdateTable(Msg SearchMsg) {
	C.Table = CreateEmptyTable().
		WithPageSize(C.H -
			lip.Height(C.TitleBar()) -
			lip.Height(C.FinderView()) -
			lip.Height(C.HelpView()) -
			lip.Height(C.MoodyBar()) -
			THH - TFH - (FHH - SHH)).
		WithRows(Msg)
	C.UpdateFooter()
}
