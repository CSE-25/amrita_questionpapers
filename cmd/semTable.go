package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type Semester struct {
	name string
	path string
}

func semTable(url string) {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Fetching Semesters").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	semesters, err := semTableReq(url)
	if err != nil {
		fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	var selectedOption string
	var sems []Semester
	var options []huh.Option[string]

	// Convert semesters to huh options.
	for _, sem := range semesters {
		semester := Semester(sem)
		sems = append(sems, semester)
		options = append(options, huh.NewOption(semester.name, semester.name))
	}
	// Add back option.
	options = append(options, huh.NewOption("Back", "Back"))

	// Create the form.
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Semesters").
				Options(options...).
				Value(&selectedOption),
		),
	)

	stack.Push(url) // Save current URL to stack.

	err = form.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	// Handle selection.
	if selectedOption == "Back" {
		huhMenuStart() // Go back to main menu.
		return
	}

	// Find selected semester and process it.
	for _, sem := range sems {
		if sem.name == selectedOption {
			url := BASE_URL + sem.path
			semChoose(url)
			break
		}
	}
}