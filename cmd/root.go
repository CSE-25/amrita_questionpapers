package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "ampyq",
    Short: "Amrita PYQ CLI",
    Long:  `A CLI application to access Amrita Repository for previous year question papers.`,
    Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(logo_ascii)
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
    fmt.Println("Available Courses:\n")
    fmt.Println("1. B.Tech")
    fmt.Println("2. BA Communication")
    fmt.Println("3. MA Communication")
    fmt.Println("4. Integrated MSc & MA")
    fmt.Println("5. MCA")
    fmt.Println("6. MSW")
    fmt.Println("7. M.Tech")
    fmt.Println("8. Exit")

    var choice int
    fmt.Printf("\nEnter your choice: ")
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        semTable(course_url+code_btech)
    case 2:
        semTable(course_url+code_ba_communication)
    case 3:
        semTable(course_url+code_ma_communication)
    case 4:
        semTable(course_url+code_integrated_msc_ma)
    case 5:
        semTable(course_url+code_mca)
    case 6:
        semTable(course_url+code_msw)
    case 7:
        semTable(course_url+code_mtech)
    case 8:
        fmt.Println("Goodbye!")
        os.Exit(0)
    default:
        fmt.Println("Invalid option!")
    }
}