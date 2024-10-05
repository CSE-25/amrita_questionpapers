package cmd

import (
	"fmt"
	"os"
	"github.com/anaskhan96/soup"
)


func year(url string) {
    fmt.Println("Fetching...")
    year_url := url
    res := fetchHTML(url)
    
    // Check if the response is empty
    if res == "" {
        fmt.Println("Failed to fetch the HTML content. Exiting.")
        return
    }

    // Parse the HTML content using soup
    doc := soup.HTMLParse(res)
    div := doc.Find("div", "xmlns","http://di.tamu.edu/DRI/1.0/")

    ul := div.Find("ul")
    li := ul.Find("li")
    hyper := li.Find("a").Attrs()["href"]

    url = BASE_URL + hyper
    page := fetchHTML(url)
    
    doc = soup.HTMLParse(page)
    div = doc.Find("div", "class","file-list")

    subdiv := div.FindAll("div","class","file-wrapper")

    // Display the found items
    fmt.Printf("No\tFiles\n")
    for i, item := range subdiv {
        title := item.FindAll("div")
        indiv := title[1].Find("div")
        span := indiv.FindAll("span")
        if span[0].Error == nil {
            fmt.Printf("%d\t%s\n", i+1, span[1].Attrs()["title"])  // Extract the text from the span element
        }
    }

    // Option to add "Back"
    fmt.Printf("%d\tBack\n", len(subdiv)+1)

    for {
        var ch int
        fmt.Print("\nEnter your choice: ")
        fmt.Scanln(&ch)

        if ch > 0 && ch <= len(subdiv) {
            title := subdiv[ch-1].FindAll("div")
            link := title[0].Find("a").Attrs()["href"]
            url = BASE_URL + link
            break
        } else if ch == len(subdiv)+1 {
            semChoose(stack.Pop())
        } else {
            fmt.Println("Please enter a valid input!")
        }
    }

    fmt.Println("Please wait until browser opens !")
    openBrowser(url)
    
    var ch int
    fmt.Println("Do you want to continue ? \nPress 1 for Yes and 0 for No : ");
    fmt.Scanln(&ch)

    if ch == 0 {
        fmt.Println("Exiting...")
        os.Exit(0)
    } else {
        year(year_url)
    }
}
