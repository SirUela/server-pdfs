package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

var (
	A4Width  = 8.5
	A4Height = 11.0

	//go:embed template.html
	html string

	r = render.New(render.Options{})
)

func SomeEndpoint(c buffalo.Context) error {
	data := map[string]interface{}{
		"random": "data",
		"more":   "data",
	}

	c.Set("needed", "data")
	c.Set("needed2", "data2")

	data2 := c.Data()
	fmt.Println(data2)

	html := &bytes.Buffer{}
	if err := r.HTML("/invoices/pdf.plush.html", "reports.plush.html").Render(html, data); err != nil {
		// handle error
	}

	file, err := Render(html.String())
	if err != nil {
		// handle error
	}

	c.Response().Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	c.Response().Header().Set("Content-Type", "application/pdf")

	// do something with the file
	fmt.Println(file)
	return nil
}

func Render(html string) (*bytes.Buffer, error) {
	u := launcher.New().
		Headless(true).
		MustLaunch()

	page := rod.New().ControlURL(u).Trace(true).MustConnect().MustPage("about:blank")

	err := page.MustWaitLoad().SetDocumentContent(html)
	if err != nil {
		return nil, fmt.Errorf("error setting document content: %w", err)
	}
	page.MustWaitLoad()

	margin := 0.0
	settings := &proto.PagePrintToPDF{
		PrintBackground: true,
		PaperWidth:      &A4Width,
		PaperHeight:     &A4Height,
		MarginTop:       &margin,
		MarginBottom:    &margin,
		MarginLeft:      &margin,
		MarginRight:     &margin,
	}

	pdf, err := page.PDF(settings)
	if err != nil {
		return nil, fmt.Errorf("error rendering PDF: %w", err)
	}
	defer pdf.Close()

	file := &bytes.Buffer{}
	_, err = io.Copy(file, pdf)
	if err != nil {
		return nil, fmt.Errorf("error copying PDF: %w", err)
	}

	return file, nil
}
