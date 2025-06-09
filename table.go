package main

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TableModel struct {
	table table.Model
}

func NewTableModel() TableModel {
	// 테이블 컬럼 설정
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Name", Width: 20},
		{Title: "Value", Width: 10},
	}

	// 테이블 데이터 설정
	rows := []table.Row{
		{"1", "Sample 1", "100"},
		{"2", "Sample 2", "200"},
		{"3", "Sample 3", "300"},
	}

	// 테이블 생성
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	// 테이블 스타일 설정
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(s)

	return TableModel{
		table: t,
	}
}

func (m TableModel) Update(msg tea.Msg) (TableModel, tea.Cmd) {
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TableModel) View() string {
	return m.table.View()
}

func (m TableModel) SetRows(rows []table.Row) {
	m.table.SetRows(rows)
}

func (m TableModel) SetWidth(width int) {
	m.table.SetWidth(width)
}

func (m TableModel) Focus() {
	m.table.Focus()
}

func (m TableModel) Blur() {
	m.table.Blur()
}
