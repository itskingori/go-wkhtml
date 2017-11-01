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

package pdf_test

import (
	"testing"

	"github.com/itskingori/go-wkhtml/pdf"
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

func TestNewOptionsFromJSON(t *testing.T) {
	d := []byte(pdfInJSON)
	p, err := pdf.NewOptionsFromJSON(d)
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
