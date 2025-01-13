package cmd

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type File struct {
	name string
	path string
}

type fileModel struct {
	state    string // "main" or "message".
	cursor   int
	files    []File
	message  string // Temporary message to display.
	err      error
	quitting bool // Track if we're quitting.
	showMenu bool // Track if we should show the menu.
}

func initialFileModel(files []File, err error) fileModel {
	return fileModel{
		state:    "main",
		cursor:   0,
		files:    files,
		err:      err,
		quitting: false,
		showMenu: true,
	}
}

func (m fileModel) Init() tea.Cmd {
	return nil
}

func (m fileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.quitting {
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			m.showMenu = false
			m.message = "Exiting..."
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.files) + 1
			}

		case "down":
			if m.cursor < len(m.files)+1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		case "enter":
			if m.cursor >= 0 && m.cursor < len(m.files) {
				url := BASE_URL + m.files[m.cursor].path
				m.message = "Opening browser, please wait..."
				m.showMenu = true // Keep showing the menu.

				return m, tea.Tick(time.Millisecond*500, func(_ time.Time) tea.Msg {
					if err := openBrowser(url); err != nil {
						return clearMessage{message: fmt.Sprintf("Error: %v", err)}
					}
					return clearMessage{message: "Browser opened successfully!"}
				})

			} else if m.cursor == len(m.files) {
				m.message = "Returning to main menu..."
				m.quitting = true
				m.showMenu = false
				return m, tea.Sequence(
					tea.Tick(time.Millisecond*1000, func(_ time.Time) tea.Msg {
						bubbleTeaStart()
						return clearMessage{message: ""}
					}),
					tea.Quit,
				)

			} else {
				m.message = "Exiting..."
				m.quitting = true
				m.showMenu = false
				return m, tea.Sequence(
					tea.Tick(time.Millisecond*1000, func(_ time.Time) tea.Msg {
						os.Exit(0)
						return clearMessage{message: ""}
					}),
					tea.Quit,
				)
			}
		}

	case clearMessage:
		m.message = msg.message
		if m.quitting {
			return m, tea.Quit
		}
		// If not quitting, keep the program running.
		return m, nil
	}

	return m, nil
}

func (m fileModel) View() string {
	var s string

	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v\n", m.err))
	}

	s += titleStyle.Render("Files") + "\n"

	if m.message != "" {
		s += fetchStatusStyle.Render(m.message) + "\n\n"
	}

	if m.showMenu {
		for i, file := range m.files {
			prefix := "  "
			listStyle := listInfo
			if i == m.cursor {
				prefix = "→ "
				listStyle = highlightStyle
			}
			s += listStyle.Render(prefix+file.name) + "\n"
		}

		prefix := "  "
		if m.cursor == len(m.files) {
			s += highlightStyle.Render("→ Back to Main Menu") + "\n"
		} else {
			s += returnStyle.Render(prefix+"Back to Main Menu") + "\n"
		}

		if m.cursor == len(m.files)+1 {
			s += highlightStyle.Render("→ Quit")
		} else {
			s += returnStyle.Render(prefix + "Quit")
		}
	}

	return s
}

// Define a message type to clear the temporary message.
type clearMessage struct {
	message string
}

func year(url string) {
	fmt.Println(fetchStatusStyle.Render("Fetching..."))

	files, err := yearReq(url)
	if err != nil {
		fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	var fileList []File
	for _, file := range files {
		fileList = append(fileList, File(file))
	}

	p := tea.NewProgram(initialFileModel(fileList, nil))
	_, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
