package cmd

import (
	"fmt"
	"os"
)

func year(url string) {
	fmt.Println(fetchStatusStyle.Render("Fetching..."))
	year_url := url

	files, err := yearReq(url)

	if err != nil {
		fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	fmt.Print(titleStyle.Render("No\tFiles\n"))
	for i, file := range files {
		fmt.Println(listInfo.Render(fmt.Sprintf("%d\t%s", i+1, file.name)))
	}

	// Option to add "Back".
	fmt.Print(returnStyle.Render(fmt.Sprintf("%d\tBack\n", len(files)+1)))

	for {
		var ch int
		fmt.Print(fetchStatusStyle.Render("Enter your choice: "))
		fmt.Scanln(&ch)

		if ch > 0 && ch <= len(files) {
			link := files[ch-1].path
			url = BASE_URL + link
			break
		} else if ch == len(files)+1 {
			semChoose(stack.Pop())
		} else {
			fmt.Println(errorStyle.Render("Please enter a valid input!"))
		}
	}

	fmt.Println(fetchStatusStyle.Render("Please wait until browser opens !"))
	if err := openBrowser(url); err != nil {
		fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
	}

	var ch int
	fmt.Print(fetchStatusStyle.Render("Do you want to continue ? \nPress 1 for Yes and 0 for No : "))
	fmt.Scanln(&ch)

	if ch == 0 {
		fmt.Println(returnStyle.Render("Exiting..."))
		os.Exit(0)
	} else {
		year(year_url)
	}
}
