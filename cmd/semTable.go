package cmd

import (
	"fmt"
)

func semTable(url string) {
	fmt.Println(fetchStatusStyle.Render("Fetching semesters..."))

	semesters, err := semTableReq(url)

	if err != nil {
		fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	fmt.Print(titleStyle.Render("No\tSemesters\n"))
	for i, semester := range semesters {
		fmt.Println(listInfo.Render(fmt.Sprintf("%d\t%s", i+1, semester.name)))
	}

	// Option to add "Back".
	fmt.Print(returnStyle.Render(fmt.Sprintf("%d\tBack", len(semesters)+1)))

	stack.Push(url)

	for {
		var ch int
		fmt.Printf("\n%s",fetchStatusStyle.Render("Enter your semester: "))
		fmt.Scanln(&ch)

		if ch > 0 && ch <= len(semesters) {
			url := BASE_URL + semesters[ch-1].path
			semChoose(url)
			break
		} else if ch == len(semesters)+1 {
			start()
			break
		} else {
			fmt.Println(errorStyle.Render("Please enter a valid input!"))
		}
	}

}
