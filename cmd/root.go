package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type Subject struct {
	name string
	path string
}

var (
	logoStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#01FAC6")).
		Bold(true)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Bold(true).
		Underline(true).
		Padding(0, 1).
		Margin(1, 0, 1, 0).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("1"))

	fetchStatusStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.Color("6")).
		Bold(true).
		Margin(1, 0)
)

var rootCmd = &cobra.Command{
	Use:   "ampyq",
	Short: "Amrita PYQ CLI",
	Long:  `A CLI application to access Amrita Repository for previous year question papers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(logoStyle.Render(LOGO_ASCII))
		huhMenuStart()
	},
}

func huhMenuStart() {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Fetching Courses").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

    resources, err := getCoursesReq(COURSE_LIST_URL)
    if err != nil {
        fmt.Println(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
        os.Exit(1)
    }

    var selectedOption string
    var subjects []Subject
    var options []huh.Option[string]

    // Convert courses to huh options.
    for _, res := range resources {
        subject := Subject(res)
        subjects = append(subjects, subject)
        options = append(options, huh.NewOption(subject.name, subject.name))
    }
    // Add quit option.
    options = append(options, huh.NewOption("Quit", "Quit"))

    // Create the form.
    form := huh.NewForm(
        huh.NewGroup(
            huh.NewSelect[string]().
                Title("Available Courses").
                Options(options...).
                Value(&selectedOption),
        ),
    )

    err = form.Run()
    if err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }

    // Handle selection.
    if selectedOption == "Quit" {
        fmt.Print(fetchStatusStyle.Render("Goodbye!\n"))
        os.Exit(0)
    }

    // Find selected subject and process it.
    for _, subject := range subjects {
        if subject.name == selectedOption {
            url := BASE_URL + subject.path
            semTable(url)
            break
        }
    }
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}