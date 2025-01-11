package cmd

import (
	"fmt"
)

func semChoose(url string) {
	fmt.Println(fetchStatusStyle.Render("Fetching assesments..."))
	params_url := url

	assesments, err := semChooseReq(url)

	if err != nil {
		fmt.Println(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	// Display the found semesters.
	fmt.Print(titleStyle.Render("No\tSemesters\n"))
	for i, assesment := range assesments {
		fmt.Println(listInfo.Render(fmt.Sprintf("%d\t%s", i+1, assesment.name))) // Extract the text from the span element.
	}

	// Option to add "Back".
	fmt.Print(returnStyle.Render(fmt.Sprintf("%d\tBack\n", len(assesments)+1)))


	for {
		var ch int
		fmt.Print(fetchStatusStyle.Render("Enter your assesment: "))
		fmt.Scanln(&ch)

		if ch > 0 && ch <= len(assesments) {
			url = BASE_URL + assesments[ch-1].path
			break
		} else if ch == len(assesments)+1 {
			semTable(stack.Pop())
		} else {
			fmt.Println(errorStyle.Render("Please enter a valid input!"))
		}
	}

	// append to stack.
	stack.Push(params_url)

	year(url)
}
