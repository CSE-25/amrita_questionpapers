package cmd

import (
	"fmt"
	"log"
    "github.com/anaskhan96/soup"
	"os/exec"
    "runtime"
)

// fetchHTML - fetches and parses HTML from the given URL
func fetchHTML(url string) string {
    doc, err := soup.Get(url)
    if err != nil {
        fmt.Println("Error fetching the URL. Make sure you're connected to Amrita WiFi or VPN.")
        return ""
    }
    return doc
}


// openBrowser - opens a URL in the default web browser
func openBrowser(url string) {
    var err error
    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    }
    if err != nil {
        log.Fatal(err)
    }
}

// Stack implementation using Go slices
type Stack struct {
	items []string
}

// Push - adds an item to the stack
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop - removes and returns the top item from the stack
func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// IsEmpty - returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Peek - returns the top item without removing it
func (s *Stack) Peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

// NewStack - creates a new stack
func NewStack() *Stack {
	return &Stack{}
}

var stack = NewStack()




