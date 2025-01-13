package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/spf13/cobra"
)

type Subject struct {
	name string // Course name.
	path string // Course path.
}

type model struct {
	cursor   int
	subjects []Subject
	err      error
}

func initialModel(subjects []Subject, err error) model {
	return model{
		cursor:   0,
		subjects: subjects,
		err:      err,
	}
}

var (
	logoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#01FAC6")). // Aqua Green.
			Bold(true)

	fetchStatusStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.Color("6")). // Dark Green.
			Bold(true).
			Margin(1, 0)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")). // Red.
			Bold(true).
			Underline(true).
			Padding(0, 1).
			Margin(1, 0, 1, 0).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("1")) // Bright Red.

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("11")). // Yellow.
			PaddingLeft(2)

	listInfo = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")). // Green.
			Italic(true).
			Bold(true)

	returnStyle = lipgloss.NewStyle().
			Bold(true).
			Italic(true).
			Foreground(lipgloss.Color("201")) // Pink

	highlightStyle = lipgloss.NewStyle().
			Inherit(listInfo).
			UnsetAlign().
			Foreground(lipgloss.Color("212")) //Light Pink
)

var rootCmd = &cobra.Command{
	Use:   "ampyq",
	Short: "Amrita PYQ CLI",
	Long:  `A CLI application to access Amrita Repository for previous year question papers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(logoStyle.Render(LOGO_ASCII))
		bubbleTeaStart()
	},
}

func bubbleTeaStart() {
	fmt.Println(fetchStatusStyle.Render("Fetching Courses..."))

	resources, err := getCoursesReq(COURSE_LIST_URL)

	var subjects []Subject
	for _, res := range resources {
		subjects = append(subjects, Subject(res))
	}
	p := tea.NewProgram(initialModel(subjects, err))
	_, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.subjects) // Move to the bottom when at the top.
			}
		case "down":
			if m.cursor < len(m.subjects) {
				m.cursor++
			} else {
				m.cursor = 0 // Move to the top when at the bottom.
			}
		case "enter":
			if m.cursor >= 0 && m.cursor < len(m.subjects) {
				url := BASE_URL + m.subjects[m.cursor].path
				semTable(url)
			} else if m.cursor == len(m.subjects) {
				fmt.Print(fetchStatusStyle.Render("Goodbye!\n"))
				os.Exit(0)
			} else {
				fmt.Print(errorStyle.Render("Please enter a valid input!\n"))
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v\n", m.err))
	}

	s := titleStyle.Render("Available Courses:\n") + "\n"
	for i, course := range m.subjects {
		prefix := "  "
		listStyle := listInfo
		if i == m.cursor {
			prefix = "→ "
			listStyle = highlightStyle
		}
		s += listStyle.Render(prefix+course.name) + "\n"
	}
	prefix := "  "
	if m.cursor == len(m.subjects) {
		s += highlightStyle.Render("→ Quit")
	} else {
		s += returnStyle.Render(prefix + "Quit")
	}

	return s
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
