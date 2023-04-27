package main

import (
	"github.com/go-rod/rod"
)

const (
	url1 = "https://www.google.com"
	url2 = "https://wawand.co/"
)

func main() {
	page := rod.New(). // Creates a new rod instance.
				MustConnect(). // Connects to the browser.
				MustPage(url2) // Creates a new page and navigate to it.

	page.MustWaitLoad() // Waits for the page to load.

	page.MustPDF("test.pdf") // Saves the page as a PDF.
}
