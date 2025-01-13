package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Assessment struct {
	name string
	path string
}

type assessModel struct {
	cursor      int
	assessments []Assessment
	err         error
}

func initialAssessModel(assessments []Assessment, err error) assessModel {
	return assessModel{
		cursor:      0,
		assessments: assessments,
		err:         err,
	}
}

func (m assessModel) Init() tea.Cmd {
	return nil
}

func (m assessModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.assessments)
			}
		case "down":
			if m.cursor < len(m.assessments) {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "enter":
			if m.cursor >= 0 && m.cursor < len(m.assessments) {
				url := BASE_URL + m.assessments[m.cursor].path
				year(url)
				break
			} else if m.cursor == len(m.assessments) {
				semTable(stack.Pop())
				return m, tea.Quit
			} else {
				fmt.Print(errorStyle.Render("Please enter a valid input!\n"))
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m assessModel) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v\n", m.err))
	}

	s := titleStyle.Render("Assessments\n") + "\n"
	for i, assessment := range m.assessments {
		prefix := "  "
		listStyle := listInfo
		if i == m.cursor {
			prefix = "→ "
			listStyle = highlightStyle
		}
		s += listStyle.Render(prefix+assessment.name) + "\n"
	}
	prefix := "  "
	if m.cursor == len(m.assessments) {
		s += highlightStyle.Render("→ Back") + "\n"
	} else {
		s += returnStyle.Render(prefix+"Back") + "\n"
	}

	return s
}

func semChoose(url string) {
	fmt.Println(fetchStatusStyle.Render("Fetching assessments..."))
	params_url := url

	assessments, err := semChooseReq(url)

	if err != nil {
		fmt.Println(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	var assessList []Assessment
	for _, assessment := range assessments {
		assessList = append(assessList, Assessment(assessment))
	}
	p := tea.NewProgram(initialAssessModel(assessList, err))
	_, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	stack.Push(params_url)
}
