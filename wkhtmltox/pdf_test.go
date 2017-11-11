// Copyright © 2017 Job King'ori Maina <j@kingori.co>
//
// This file is part of go-wkhtml.
//
// go-wkhtml is free software: you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version.
//
// go-wkhtml is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE.  See the GNU Lesser General Public License for more
// details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with go-wkhtml. If not, see <http://www.gnu.org/licenses/>.

package wkhtmltox_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/itskingori/go-wkhtml/wkhtmltox"
)

var pdfInJSON = `
{
  "margin_top": 10,
  "margin_bottom": 10,
  "margin_left": 10,
  "margin_right": 10,
  "page_height": 10,
  "page_width": 10,
  "page_size": "A4"
}`

func TestNewPDFOptionsFromJSON(t *testing.T) {
	d := []byte(pdfInJSON)
	p, err := wkhtmltox.NewPDFOptionsFromJSON(d)
	if err != nil {
		t.Fatal(err)
	}

	if *p.MarginTop != 10 {
		t.Fatalf("expected %s to be %d, got %d", "MarginTop", 10, *p.MarginTop)
	}

	if *p.MarginBottom != 10 {
		t.Fatalf("expected %s to be %d, got %d", "MarginBottom", 10, *p.MarginBottom)
	}

	if *p.MarginLeft != 10 {
		t.Fatalf("expected %s to be %d, got %d", "MarginLeft", 10, *p.MarginLeft)
	}

	if *p.MarginRight != 10 {
		t.Fatalf("expected %s to be %d, got %d", "MarginRight", 10, *p.MarginRight)
	}

	if *p.PageHeight != 10 {
		t.Fatalf("expected %s to be %d, got %d", "PageHeight", 10, *p.PageHeight)
	}

	if *p.PageWidth != 10 {
		t.Fatalf("expected %s to be %d, got %d", "PageWidth", 10, *p.PageWidth)
	}

	if *p.PageSize != "A4" {
		t.Fatalf("expected %s to be %s, got %s", "PageSize", "A4", *p.PageSize)
	}
}

func TestPDFOptionsFlags(t *testing.T) {
	marginTop := 10
	marginBottom := 10
	marginLeft := 10
	marginRight := 10
	pageHeight := 10
	pageWidth := 10
	pageSize := "A4"

	gopts := wkhtmltox.PDFGlobalOptions{
		MarginTop:    &marginTop,
		MarginBottom: &marginBottom,
		MarginLeft:   &marginLeft,
		MarginRight:  &marginRight,
		PageHeight:   &pageHeight,
		PageWidth:    &pageWidth,
		PageSize:     &pageSize,
	}

	opts := wkhtmltox.PDFOptions{PDFGlobalOptions: gopts}
	flags := opts.Flags()

	if flags["margin-top"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-h", 10, flags["margin-top"])
	}

	if flags["margin-bottom"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-w", 10, flags["margin-bottom"])
	}

	if flags["margin-left"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-x", 10, flags["margin-left"])
	}

	if flags["margin-right"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-y", 10, flags["margin-right"])
	}

	if flags["page-height"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "format", 10, flags["page-height"])
	}

	if flags["page-width"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "height", 10, flags["page-width"])
	}

	if flags["page-size"] != "A4" {
		t.Fatalf("expected %s to be %s, got %s", "width", "A4", flags["page-size"])
	}
}

func TestPDFFlagSetsFlags(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSets)
	fss["margin-top"] = 10
	fss["page-size"] = "A4"

	expected := []string{"--margin-top", "10", "--page-size", "A4"}
	got := fss.Flags()

	sort.Strings(expected)
	sort.Strings(got)

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}
}

func TestPDFFlagSetsGetMarginTop(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSets)
	fss["margin-top"] = 10

	value, exists := fss.GetMarginTop()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "margin-top", fss["margin-top"], value)
	}
}

func TestPDFFlagSetsGetMarginBottom(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSets)
	fss["margin-bottom"] = 10

	value, exists := fss.GetMarginBottom()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "margin-bottom", fss["margin-bottom"], value)
	}
}

func TestPDFFlagSetsGetMarginLeft(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSets)
	fss["margin-left"] = 10

	value, exists := fss.GetMarginLeft()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "margin-left", fss["margin-left"], value)
	}
}

func TestPDFFlagSetsGetMarginRight(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSets)
	fss["margin-right"] = 10

	value, exists := fss.GetMarginRight()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "margin-right", fss["margin-right"], value)
	}
}

func TestPDFFlagSetsGetPageHeight(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSets)
	fss["page-height"] = 10

	value, exists := fss.GetPageHeight()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "page-height", fss["page-height"], value)
	}
}

func TestPDFFlagSetsGetPageWidth(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSets)
	fss["page-width"] = 10

	value, exists := fss.GetPageWidth()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "page-width", fss["page-width"], value)
	}
}

func TestPDFFlagSetsGetPageSize(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSets)
	fss["page-size"] = "A4"

	value, exists := fss.GetPageSize()
	if exists != true || value != "A4" {
		t.Fatalf("expected %s to be %s, got %s", "page-size", fss["page-size"], value)
	}
}
