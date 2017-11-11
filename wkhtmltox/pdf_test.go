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

func TestPDFOptionsFlagSet(t *testing.T) {
	marginTop := 10
	marginBottom := 10
	marginLeft := 10
	marginRight := 10
	pageHeight := 10
	pageWidth := 10
	pageSize := "A4"

	opts := wkhtmltox.PDFOptions{
		MarginTop:    &marginTop,
		MarginBottom: &marginBottom,
		MarginLeft:   &marginLeft,
		MarginRight:  &marginRight,
		PageHeight:   &pageHeight,
		PageWidth:    &pageWidth,
		PageSize:     &pageSize,
	}
	pfs := opts.FlagSet()

	if pfs["margin-top"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-h", 10, pfs["margin-top"])
	}

	if pfs["margin-bottom"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-w", 10, pfs["margin-bottom"])
	}

	if pfs["margin-left"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-x", 10, pfs["margin-left"])
	}

	if pfs["margin-right"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-y", 10, pfs["margin-right"])
	}

	if pfs["page-height"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "format", 10, pfs["page-height"])
	}

	if pfs["page-width"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "height", 10, pfs["page-width"])
	}

	if pfs["page-size"] != "A4" {
		t.Fatalf("expected %s to be %s, got %s", "width", "A4", pfs["page-size"])
	}
}

func TestPDFFlagSetFlags(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSet)
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

func TestPDFFlagSetGetMarginTop(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSet)
	fss["margin-top"] = 10

	value, exists := fss.GetMarginTop()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "margin-top", fss["margin-top"], value)
	}
}

func TestPDFFlagSetGetMarginBottom(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSet)
	fss["margin-bottom"] = 10

	value, exists := fss.GetMarginBottom()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "margin-bottom", fss["margin-bottom"], value)
	}
}

func TestPDFFlagSetGetMarginLeft(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSet)
	fss["margin-left"] = 10

	value, exists := fss.GetMarginLeft()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "margin-left", fss["margin-left"], value)
	}
}

func TestPDFFlagSetGetMarginRight(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSet)
	fss["margin-right"] = 10

	value, exists := fss.GetMarginRight()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "margin-right", fss["margin-right"], value)
	}
}

func TestPDFFlagSetGetPageHeight(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSet)
	fss["page-height"] = 10

	value, exists := fss.GetPageHeight()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "page-height", fss["page-height"], value)
	}
}

func TestPDFFlagSetGetPageWidth(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSet)
	fss["page-width"] = 10

	value, exists := fss.GetPageWidth()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "page-width", fss["page-width"], value)
	}
}

func TestPDFFlagSetGetPageSize(t *testing.T) {
	fss := make(wkhtmltox.PDFFlagSet)
	fss["page-size"] = "A4"

	value, exists := fss.GetPageSize()
	if !exists || value != "A4" {
		t.Fatalf("expected %s to be %s, got %s", "page-size", fss["page-size"], value)
	}
}
