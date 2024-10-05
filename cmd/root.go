package cmd

import (
	"fmt"
	"os"
	"github.com/anaskhan96/soup"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "ampyq",
    Short: "Amrita PYQ CLI",
    Long:  `A CLI application to access Amrita Repository for previous year question papers.`,
    Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(LOGO_ASCII)
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

// start function - equivalent to start() in Python
func start() {
    fmt.Println("Fetching Courses...")
    res := fetchHTML(COURSE_LIST_URL)
    fmt.Println("Available Courses:")

    // Check if the response is empty
    if res == "" {
        fmt.Println("Failed to fetch the HTML content. Exiting.")
        return
    }

    // Parse the HTML content using soup
    doc := soup.HTMLParse(res)
    div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")

    subs := div.FindAll("div","class","artifact-title")

    for i, item := range subs {
        sub := item.Find("span")
        if sub.Error == nil {
            fmt.Printf("%d.\t%s\n", i+1, sub.Text())
        }
    }

    // Option to quit.
    fmt.Printf("%d.\tQuit\n", len(subs)+1)

    for {
        var ch int
        fmt.Printf("\nEnter your choice: ")
        fmt.Scanln(&ch)

        if ch > 0 && ch <= len(subs) {
            a := subs[ch-1].Find("a")
            path := a.Attrs()["href"]
            url := BASE_URL + path
            semTable(url)
        } else if ch == len(subs)+1 {
            fmt.Println("Goodbye!")
            os.Exit(0)
        } else {
            fmt.Println("Please enter a valid input!")
        }
    }
}