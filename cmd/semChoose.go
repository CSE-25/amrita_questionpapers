package cmd

import (
	"fmt"
	"github.com/anaskhan96/soup"
)

func semChoose(url string) {
	fmt.Println("Fetching assesments...")
	params_url := url
    resp := fetchHTML(url)  // Fetch the HTML content

	// Check if the response is empty
    if resp == "" {
        fmt.Println("Failed to fetch the HTML content. Exiting.")
        return
    }
    
    // Parse the HTML content using soup
    doc := soup.HTMLParse(resp)
    div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")

	if div.Error != nil {
        fmt.Println("No assesments found on the page.")
        return
    }

	ul := div.FindAll("ul")
	li := ul[0].FindAll("li")

	if len(ul)>1 {
		li = ul[1].FindAll("li")
	} else {
		li = ul[0].FindAll("li")
	}

	// Display the found items
    fmt.Printf("No\tSemesters\n")
    for i, link := range li {
        a := link.Find("a")
        if a.Error == nil {
            span := a.Find("span")
            if span.Error == nil {
                fmt.Printf("%d\t%s\n", i+1, span.Text())  // Extract the text from the span element
            }
        }
    }

	// Option to add "Back"
    fmt.Printf("%d\tBack\n", len(li)+1)


	for {
        var ch int
        fmt.Print("\nEnter your assesment: ")
        fmt.Scanln(&ch)

        if ch > 0 && ch <= len(li) {
            url = base_url + li[ch-1].Find("a").Attrs()["href"]
            break
        } else if ch == len(li)+1 {
            semTable(stack.Pop())
        } else {
            fmt.Println("Please enter a valid input!")
        }
    }

	// append to stack
    stack.Push(params_url)

	year(url)
}