package app

import (
	"fmt"
	"github.com/astrogewgaw/koshka/data"
	"github.com/evertras/bubble-table/table"

	tea "github.com/charmbracelet/bubbletea"
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
					"ID":   fmt.Sprintf("%d", psr.ID+1),
					"P0":   fmt.Sprintf("%f", psr.P0.Float64),
					"DM":   fmt.Sprintf("%f", psr.DM.Float64),
				}).WithStyle(TableRowNormalStyle))
	}
	return rows
}

func ApplySearch(C Cat) tea.Cmd {
	return func() tea.Msg { return SearchMsg(DataIntoTable(C.Finder.Value())) }
}
