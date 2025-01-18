package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type Assessment struct {
	name string
	path string
}

func semChoose(url string) {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Fetching assessments").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	params_url := url

	assessments, err := semChooseReq(url)
	if err != nil {
		fmt.Println(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	var selectedOption string
	var assessList []Assessment
	var options []huh.Option[string]

	// Convert assessments to huh options.
	for _, assessment := range assessments {
		assess := Assessment(assessment)
		assessList = append(assessList, assess)
		options = append(options, huh.NewOption(assess.name, assess.name))
	}
	// Add back option.
	options = append(options, huh.NewOption("Back", "Back"))

	// Create the form.
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Assessments").
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
	if selectedOption == "Back" {
		semTable(stack.Pop())
		return
	}

	// Find selected assessment and process it.
	for _, assess := range assessList {
		if assess.name == selectedOption {
			url := BASE_URL + assess.path
			year(url)
			break
		}
	}

	stack.Push(params_url)
}