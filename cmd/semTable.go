package cmd

import (
    "fmt"
    "github.com/anaskhan96/soup"
)

// semTable function - equivalent to your Python code
func semTable(url string) {
    fmt.Println("Fetching semesters...")
    res := fetchHTML(url)  // Fetch the HTML content

    // Check if the response is empty
    if res == "" {
        fmt.Println("Failed to fetch the HTML content. Exiting.")
        return
    }
    
    // Parse the HTML content using soup
    doc := soup.HTMLParse(res)
    div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")

    if div.Error != nil {
        fmt.Println("No semesters found on the page.")
        return
    }

    ul := div.Find("ul")
    li := ul.FindAll("li")

    if len(li) == 0 {
        fmt.Println("No semesters found on the page.")
        return
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

    // append to stack
    stack.Push(url)


    for {
        var ch int
        fmt.Print("\nEnter your semester: ")
        fmt.Scanln(&ch)

        if ch > 0 && ch <= len(li) {
            url := BASE_URL + li[ch-1].Find("a").Attrs()["href"]
            semChoose(url)
            break
        } else if ch == len(li)+1 {
            start()
            break
        } else {
            fmt.Println("Please enter a valid input!")
        }
    }
    
}
