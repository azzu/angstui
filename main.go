package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	list   ListModel
	table  TableModel
	width  int
	height int
	focus  string // "list" or "table"
}

func initialModel() model {
	return model{
		list:  NewListModel(),
		table: NewTableModel(),
		focus: "list", // 초기 포커스는 리스트에
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width/2-h, msg.Height-v)
		m.table.SetWidth(msg.Width/2 - h)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			// 탭 키로 포커스 전환
			if m.focus == "list" {
				m.focus = "table"
				m.list.SetFocus(false)
				m.table.Focus()
			} else {
				m.focus = "list"
				m.list.SetFocus(true)
				m.table.Blur()
			}
			return m, nil
		}
	}

	// 리스트 업데이트
	var listCmd tea.Cmd
	m.list, listCmd = m.list.Update(msg)
	cmds = append(cmds, listCmd)

	// 리스트 아이템 선택 처리
	if m.focus == "list" {
		if i, ok := m.list.SelectedItem(); ok {
			// 선택된 아이템에 따라 테이블 데이터 업데이트
			rows := []table.Row{
				{"1", i.title + " - Detail 1", "100"},
				{"2", i.title + " - Detail 2", "200"},
				{"3", i.title + " - Detail 3", "300"},
			}
			m.table.SetRows(rows)
			// 자동으로 테이블로 포커스 이동
			m.focus = "table"
			m.list.SetShowHelp(false)
			m.table.Focus()
		}
	}

	// 테이블 업데이트
	var tableCmd tea.Cmd
	m.table, tableCmd = m.table.Update(msg)
	cmds = append(cmds, tableCmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// 리스트와 테이블의 스타일 설정
	listStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))

	tableStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))

	// 포커스에 따른 스타일 변경
	if m.focus == "list" {
		listStyle = listStyle.BorderForeground(lipgloss.Color("205"))
	} else {
		tableStyle = tableStyle.BorderForeground(lipgloss.Color("205"))
	}

	// 화면을 두 부분으로 나누기
	left := listStyle.Render(m.list.View())
	right := tableStyle.Render(m.table.View())

	// 두 부분을 나란히 배치
	return docStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		left,
		strings.Repeat(" ", 2), // 간격
		right,
	))
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
