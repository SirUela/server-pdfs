package main

import (
	_ "embed"

	"github.com/go-rod/rod"
)

//go:embed template.html
var html string

func main() {
	page := rod.New(). // Creates a new rod instance.
				MustConnect(). // Connects to the browser.
				MustPage()     // Creates a new page and navigate to it. Since no URL is provided, it will navigate to "about:blank".

	page.SetDocumentContent(html) // Sets the page content.

	page.MustWaitLoad() // Waits for the page to load.

	page.MustPDF("test.pdf") // Saves the page as a PDF.
}
