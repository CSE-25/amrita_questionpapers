package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Semester struct {
	name string
	path string
}

type semModel struct {
	cursor    int
	semesters []Semester
	err       error
}

func initialSemModel(semesters []Semester, err error) semModel {
	return semModel{
		cursor:    0,
		semesters: semesters,
		err:       err,
	}
}

func (m semModel) Init() tea.Cmd {
	return nil
}

func (m semModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.semesters) // Move to the bottom when at the top.
			}
		case "down":
			if m.cursor < len(m.semesters) {
				m.cursor++
			} else {
				m.cursor = 0 // Move to the top when at the bottom.
			}
		case "enter":
			if m.cursor >= 0 && m.cursor < len(m.semesters) {
				url := BASE_URL + m.semesters[m.cursor].path
				semChoose(url)
				break
			} else if m.cursor == len(m.semesters) {
				bubbleTeaStart()
				return m, tea.Quit
			} else {
				fmt.Print(errorStyle.Render("Please enter a valid input!\n"))
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m semModel) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v\n", m.err))
	}

	s := titleStyle.Render("Semesters\n") + "\n"
	for i, semester := range m.semesters {
		prefix := "  "
		listStyle := listInfo
		if i == m.cursor {
			prefix = "→ "
			listStyle = highlightStyle
		}
		s += listStyle.Render(prefix+semester.name) + "\n"
	}

	prefix := "  "
	if m.cursor == len(m.semesters) {
		s += highlightStyle.Render("→ Back")
	} else {
		s += returnStyle.Render(prefix + "Back")
	}

	return s
}

func semTable(url string) {
	fmt.Print(fetchStatusStyle.Render("Fetching semesters...\n"))

	semesters, err := semTableReq(url)

	if err != nil {
		fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	var sems []Semester
	for _, sem := range semesters {
		sems = append(sems, Semester(sem)) // Convert resource to Semester.
	}
	p := tea.NewProgram(initialSemModel(sems, err))
	stack.Push(url)
	_, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
