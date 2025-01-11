package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/lipgloss/v2"
)

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
    fetchStatusStyle = lipgloss.NewStyle().
        PaddingLeft(1).
        Foreground(lipgloss.Color("6")).	// Cyan
        Bold(true).
        MarginTop(1)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).    // Bright Red
		Bold(true).                         
		Underline(true).                   
		Padding(0, 1).                     
		Margin(1, 0, 1, 0).                 
		Border(lipgloss.RoundedBorder()).   
		BorderForeground(lipgloss.Color("1")) 
	
    titleStyle = lipgloss.NewStyle().
        PaddingLeft(1).
		Bold(true).
        Foreground(lipgloss.Color("11")).   // Yellow
        MarginTop(1)

	listInfo = lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")). 	// Green
		PaddingLeft(1).
		Italic(true).
		Bold(true).
		MarginTop(1)

	returnStyle =lipgloss.NewStyle().
		PaddingLeft(1).
		Bold(true).
		Italic(true).
		Foreground(lipgloss.Color("201")).   // Orange
		MarginTop(1)
)

var rootCmd = &cobra.Command{
	Use:   "ampyq",
	Short: "Amrita PYQ CLI",
	Long:  `A CLI application to access Amrita Repository for previous year question papers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(logoStyle.Render(LOGO_ASCII))
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// start lists the available courses and redirects to chosen course.
func start() {
	fmt.Println(fetchStatusStyle.Render("Fetching Courses..."))

	subjects, err := getCoursesReq(COURSE_LIST_URL)

	if err != nil {
		fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	fmt.Println(titleStyle.Render(`Available Courses:`))

	for i, subject := range subjects {
		fmt.Println(listInfo.Render(fmt.Sprintf("%d.\t%s", i+1, subject.name)))
	}

	// Option to quit.
	fmt.Print(returnStyle.Render(fmt.Sprintf("%d.\tQuit", len(subjects)+1)))

	for {
		var ch int
		fmt.Printf("\n%s", fetchStatusStyle.Render("Enter your choice: "))
		fmt.Scanln(&ch)

		if ch > 0 && ch <= len(subjects) {
			path := subjects[ch-1].path
			url := BASE_URL + path
			semTable(url)
		} else if ch == len(subjects)+1 {
			fmt.Print(returnStyle.Render("Goodbye!\n"))	
			os.Exit(0)
		} else {
			fmt.Print(errorStyle.Render("Please enter a valid input!\n"))		}
	}
}
