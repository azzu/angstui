package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type ListModel struct {
	list  list.Model
	focus bool
}

func NewListModel() ListModel {
	// 리스트 아이템 생성
	items := []list.Item{
		item{title: "Item 1", description: "Description for Item 1"},
		item{title: "Item 2", description: "Description for Item 2"},
		item{title: "Item 3", description: "Description for Item 3"},
	}

	// 리스트 설정
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select an Item"
	l.SetShowHelp(true)
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.Styles.Title = lipgloss.NewStyle().
		MarginLeft(2).
		Foreground(lipgloss.Color("205"))

	// 첫 번째 아이템에 포커스 설정
	l.Select(0)

	return ListModel{
		list:  l,
		focus: true,
	}
}

func (m ListModel) Update(msg tea.Msg) (ListModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.focus {
			switch msg.String() {
			case "up", "k":
				m.list.CursorUp()
			case "down", "j":
				m.list.CursorDown()
			case "enter":
				// Enter 키를 눌렀을 때만 선택 이벤트를 발생시킴
				if selectedItem, ok := m.list.SelectedItem().(item); ok {
					return m, tea.Cmd(func() tea.Msg {
						return ItemSelectedMsg{Item: selectedItem}
					})
				}
			default:
				// Enter 키가 아닌 다른 키는 기본 리스트 동작으로 처리
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListModel) View() string {
	return m.list.View()
}

func (m ListModel) SelectedItem() (item, bool) {
	if i, ok := m.list.SelectedItem().(item); ok {
		return i, true
	}
	return item{}, false
}

func (m ListModel) SetSize(width, height int) {
	m.list.SetSize(width, height)
}

func (m ListModel) SetShowHelp(show bool) {
	m.list.SetShowHelp(show)
}

func (m ListModel) SetFocus(focus bool) {
	m.focus = focus
	if focus {
		m.list.SetShowHelp(true)
	} else {
		m.list.SetShowHelp(false)
	}
}

// ItemSelectedMsg는 아이템이 선택되었을 때 발생하는 메시지
type ItemSelectedMsg struct {
	Item item
}
