package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strconv"
)

var baseStyle = lipgloss.NewStyle().
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func printVmStatus(node NodeInfo, vms []VmInfo) {
	columns := []table.Column{
		{Title: "VMID", Width: 4},
		{Title: "NAME", Width: 12},
		{Title: "STATUS", Width: 8},
		{Title: "CPUs", Width: 5},
		{Title: "CPU %", Width: 5},
		{Title: "MEM %", Width: 5},
	}

	rows := []table.Row{}
	for _, vm := range vms {
		rows = append(rows, table.Row{strconv.FormatInt(int64(vm.Vmid), 10),
			vm.Name,
			vm.Status,
			strconv.FormatInt(int64(vm.Cpus), 10),
			strconv.FormatFloat(float64(100.0*vm.Cpu), 'f', 1, 32),
			strconv.FormatFloat(float64(100.0*int64(vm.Mem)/vm.Maxmem), 'f', 1, 32)})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(vms)+1),
	)

	s := table.DefaultStyles()
	myBorder := lipgloss.Border{}

	s.Header = s.Header.Border(myBorder).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(false).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("255")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
