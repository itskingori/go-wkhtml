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
	"testing"

	"github.com/itskingori/go-wkhtml/wkhtmltox"
)

func TestPDFNewPDFFlagSetFromOptions(t *testing.T) {
	pageSize := "A4"

	opts := wkhtmltox.PDFOptions{
		PageSize: &pageSize,
	}
	pfs := wkhtmltox.NewPDFFlagSetFromOptions(&opts)

	// when set
	if value, exists := pfs["page-size"]; !exists && value != "A4" {
		t.Fatalf("expected %s to be %s, got %s", "width", "A4", pfs["page-size"])
	}

	// when not set
	if _, exists := pfs["page-width"]; exists {
		t.Fatalf("expected %s to be %d, got %d", "height", 10, pfs["page-width"])
	}
}

func TestPDFFlagSetFlags(t *testing.T) {
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["cache-dir"] = "/some/dir"
	expected := []string{"--cache-dir", "/some/dir"}
	got := pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["cookie"] = []wkhtmltox.CookieSet{
		wkhtmltox.CookieSet{
			Name:  "cookie1",
			Value: "value1",
		},
		wkhtmltox.CookieSet{
			Name:  "cookie2",
			Value: "value2",
		},
	}
	expected = []string{"--cookie", "cookie1", "value1", "--cookie", "cookie2", "value2"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["custom-header"] = []wkhtmltox.HeaderSet{
		wkhtmltox.HeaderSet{
			Name:  "header1",
			Value: "value1",
		},
		wkhtmltox.HeaderSet{
			Name:  "header2",
			Value: "value2",
		},
	}
	expected = []string{"--custom-header", "header1", "value1", "--custom-header", "header2", "value2"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["custom-header-propagation"] = true
	expected = []string{"--custom-header-propagation"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["custom-header-propagation"] = false
	expected = []string{"--no-custom-header-propagation"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["debug-javascript"] = true
	expected = []string{"--debug-javascript"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["debug-javascript"] = false
	expected = []string{"--no-debug-javascript"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["dpi"] = 600
	expected = []string{"--dpi", "600"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["encoding"] = "utf-8"
	expected = []string{"--encoding", "utf-8"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["external-links"] = false
	expected = []string{"--disable-external-links"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["external-links"] = true
	expected = []string{"--enable-external-links"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["forms"] = false
	expected = []string{"--disable-forms"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["forms"] = true
	expected = []string{"--enable-forms"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["grayscale"] = true
	expected = []string{"--grayscale"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["grayscale"] = false
	got = pfs.Flags()
	if len(got) != 0 {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["images"] = true
	expected = []string{"--images"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["images"] = false
	expected = []string{"--no-images"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["image-dpi"] = 600
	expected = []string{"--image-dpi", "600"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["image-quality"] = 94
	expected = []string{"--image-quality", "94"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["internal-links"] = false
	expected = []string{"--disable-internal-links"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["internal-links"] = true
	expected = []string{"--enable-internal-links"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["javascript"] = false
	expected = []string{"--disable-javascript"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["javascript"] = true
	expected = []string{"--enable-javascript"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["javascript-delay"] = 200
	expected = []string{"--javascript-delay", "200"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["load-error-handling"] = "abort"
	expected = []string{"--load-error-handling", "abort"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["load-media-error-handling"] = "abort"
	expected = []string{"--load-media-error-handling", "abort"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["lowquality"] = true
	expected = []string{"--lowquality"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["lowquality"] = false
	got = pfs.Flags()
	if len(got) != 0 {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["margin-bottom"] = 10
	expected = []string{"--margin-bottom", "10"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["margin-left"] = 10
	expected = []string{"--margin-left", "10"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["margin-right"] = 10
	expected = []string{"--margin-right", "10"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["margin-top"] = 10
	expected = []string{"--margin-top", "10"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["minimum-font-size"] = 10
	expected = []string{"--minimum-font-size", "10"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["orientation"] = "landscape"
	expected = []string{"--orientation", "landscape"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["page-height"] = 10
	expected = []string{"--page-height", "10"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["page-size"] = "A4"
	expected = []string{"--page-size", "A4"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["page-width"] = 10
	expected = []string{"--page-width", "10"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["no-pdf-compression"] = true
	expected = []string{"--no-pdf-compression"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["no-pdf-compression"] = false
	expected = []string{"--no-pdf-compression"}
	got = pfs.Flags()
	if len(got) != 0 {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["password"] = "sekret"
	expected = []string{"--password", "sekret"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["smart-shrinking"] = false
	expected = []string{"--disable-smart-shrinking"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["smart-shrinking"] = true
	expected = []string{"--enable-smart-shrinking"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["stop-slow-scripts"] = true
	expected = []string{"--stop-slow-scripts"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["stop-slow-scripts"] = false
	expected = []string{"--no-stop-slow-scripts"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["title"] = "Some Title"
	expected = []string{"--title", "Some Title"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["use-xserver"] = true
	expected = []string{"--use-xserver"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["use-xserver"] = false
	got = pfs.Flags()
	if len(got) != 0 {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["username"] = "some_name"
	expected = []string{"--username", "some_name"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	pfs = make(wkhtmltox.PDFFlagSet)
	pfs["zoom"] = 0.5
	expected = []string{"--zoom", "0.5"}
	got = pfs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}
}

func TestPDFFlagSetGetCacheDir(t *testing.T) {
	attribute := "cache-dir"
	dir := "/tmp/xyz"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["cache-dir"] = dir
	result, exists := pfs.GetCacheDir()

	if !exists || result != dir {
		t.Fatalf("expected %s to be %s, got %s", attribute, dir, result)
	}
}

func TestPDFFlagSetGetCookie(t *testing.T) {
	attribute := "cookie"
	cookies := []wkhtmltox.CookieSet{
		wkhtmltox.CookieSet{
			Name:  "cookie1",
			Value: "value1",
		},
		wkhtmltox.CookieSet{
			Name:  "cookie2",
			Value: "value2",
		},
	}
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["cookie"] = cookies
	result, exists := pfs.GetCookie()

	if !exists && !reflect.DeepEqual(result, cookies) {
		t.Fatalf("expected %s to be %s, got %s", attribute, cookies, result)
	}
}

func TestPDFFlagSetGetCustomHeader(t *testing.T) {
	attribute := "custom-header"
	headers := []wkhtmltox.HeaderSet{
		wkhtmltox.HeaderSet{
			Name:  "header1",
			Value: "value1",
		},
		wkhtmltox.HeaderSet{
			Name:  "header2",
			Value: "value2",
		},
	}
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["custom-header"] = headers
	result, exists := pfs.GetCustomHeader()

	if !exists && !reflect.DeepEqual(result, headers) {
		t.Fatalf("expected %s to be %s, got %s", attribute, headers, result)
	}
}

func TestPDFFlagSetGetCustomHeaderPropagation(t *testing.T) {
	attribute := "custom-header-propagation"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["custom-header-propagation"] = value
	result, exists := pfs.GetCustomHeaderPropagation()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetDebugJavascript(t *testing.T) {
	attribute := "debug-javascript"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["debug-javascript"] = value
	result, exists := pfs.GetDebugJavascript()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetDPI(t *testing.T) {
	attribute := "dpi"
	dpi := 600
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["dpi"] = dpi
	result, exists := pfs.GetDPI()

	if !exists || result != dpi {
		t.Fatalf("expected %s to be %d, got %d", attribute, dpi, result)
	}
}

func TestPDFFlagSetGetEncoding(t *testing.T) {
	attribute := "encoding"
	encoding := "utf-8"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["encoding"] = encoding
	result, exists := pfs.GetEncoding()

	if !exists || result != encoding {
		t.Fatalf("expected %s to be %s, got %s", attribute, encoding, result)
	}
}

func TestPDFFlagSetGetExternalLinks(t *testing.T) {
	attribute := "external-links"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["external-links"] = value
	result, exists := pfs.GetExternalLinks()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetForms(t *testing.T) {
	attribute := "forms"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["forms"] = value
	result, exists := pfs.GetForms()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetGrayscale(t *testing.T) {
	attribute := "grayscale"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["grayscale"] = value
	result, exists := pfs.GetGrayscale()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetImages(t *testing.T) {
	attribute := "images"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["images"] = value
	result, exists := pfs.GetImages()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetImageDPI(t *testing.T) {
	attribute := "image-dpi"
	dpi := 600
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["image-dpi"] = dpi
	result, exists := pfs.GetImageDPI()

	if !exists || result != dpi {
		t.Fatalf("expected %s to be %d, got %d", attribute, dpi, result)
	}
}

func TestPDFFlagSetGetImageQuality(t *testing.T) {
	attribute := "image-quality"
	quality := 94
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["image-quality"] = quality
	result, exists := pfs.GetImageQuality()

	if !exists || result != quality {
		t.Fatalf("expected %s to be %d, got %d", attribute, quality, result)
	}
}

func TestPDFFlagSetGetInternalLinks(t *testing.T) {
	attribute := "internal-links"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["internal-links"] = value
	result, exists := pfs.GetInternalLinks()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetJavascript(t *testing.T) {
	attribute := "javascript"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["javascript"] = value
	result, exists := pfs.GetJavascript()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetJavascriptDelay(t *testing.T) {
	attribute := "javascript-delay"
	milliseconds := 200
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["javascript-delay"] = milliseconds
	result, exists := pfs.GetJavascriptDelay()

	if !exists || result != milliseconds {
		t.Fatalf("expected %s to be %d, got %d", attribute, milliseconds, result)
	}
}

func TestPDFFlagSetGetLoadErrorHandling(t *testing.T) {
	attribute := "load-error-handling"
	handling := "abort"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["load-error-handling"] = handling
	result, exists := pfs.GetLoadErrorHandling()

	if !exists || result != handling {
		t.Fatalf("expected %s to be %s, got %s", attribute, handling, result)
	}
}

func TestPDFFlagSetGetLoadMediaErrorHandling(t *testing.T) {
	attribute := "load-media-error-handling"
	handling := "abort"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["load-media-error-handling"] = handling
	result, exists := pfs.GetLoadMediaErrorHandling()

	if !exists || result != handling {
		t.Fatalf("expected %s to be %s, got %s", attribute, handling, result)
	}
}

func TestPDFFlagSetGetLowQuality(t *testing.T) {
	attribute := "lowquality"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["lowquality"] = value
	result, exists := pfs.GetLowQuality()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetMarginBottom(t *testing.T) {
	attribute := "margin-bottom"
	millimetres := 10
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["margin-bottom"] = millimetres
	result, exists := pfs.GetMarginBottom()

	if !exists || result != millimetres {
		t.Fatalf("expected %s to be %d, got %d", attribute, millimetres, result)
	}
}

func TestPDFFlagSetGetMarginLeft(t *testing.T) {
	attribute := "margin-left"
	millimetres := 10
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["margin-left"] = millimetres
	result, exists := pfs.GetMarginLeft()

	if !exists || result != millimetres {
		t.Fatalf("expected %s to be %d, got %d", attribute, millimetres, result)
	}
}

func TestPDFFlagSetGetMarginRight(t *testing.T) {
	attribute := "margin-right"
	millimetres := 10
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["margin-right"] = millimetres
	result, exists := pfs.GetMarginRight()

	if !exists || result != millimetres {
		t.Fatalf("expected %s to be %d, got %d", attribute, millimetres, result)
	}
}

func TestPDFFlagSetGetMarginTop(t *testing.T) {
	attribute := "margin-top"
	millimetres := 10
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["margin-top"] = millimetres
	result, exists := pfs.GetMarginTop()

	if !exists || result != millimetres {
		t.Fatalf("expected %s to be %d, got %d", attribute, millimetres, result)
	}
}

func TestPDFFlagSetGetMinimumFontSize(t *testing.T) {
	attribute := "minimum-font-size"
	size := 12
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["minimum-font-size"] = size
	result, exists := pfs.GetMinimumFontSize()

	if !exists || result != size {
		t.Fatalf("expected %s to be %d, got %d", attribute, size, result)
	}
}

func TestPDFFlagSetGetNoPDFCompression(t *testing.T) {
	attribute := "no-pdf-compression"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["no-pdf-compression"] = value
	result, exists := pfs.GetNoPDFCompression()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetOrientation(t *testing.T) {
	attribute := "orientation"
	orientation := "landscape"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["orientation"] = orientation
	result, exists := pfs.GetOrientation()

	if !exists || result != orientation {
		t.Fatalf("expected %s to be %s, got %s", attribute, orientation, result)
	}
}

func TestPDFFlagSetGetPageHeight(t *testing.T) {
	attribute := "page-height"
	height := 300
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["page-height"] = height
	result, exists := pfs.GetPageHeight()

	if !exists || result != height {
		t.Fatalf("expected %s to be %d, got %d", attribute, height, result)
	}
}

func TestPDFFlagSetGetPageSize(t *testing.T) {
	attribute := "page-size"
	size := "A4"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["page-size"] = size
	result, exists := pfs.GetPageSize()

	if !exists || result != size {
		t.Fatalf("expected %s to be %s, got %s", attribute, size, result)
	}
}

func TestPDFFlagSetGetPageWidth(t *testing.T) {
	attribute := "page-width"
	width := 300
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["page-width"] = width
	result, exists := pfs.GetPageWidth()

	if !exists || result != width {
		t.Fatalf("expected %s to be %d, got %d", attribute, width, result)
	}
}

func TestPDFFlagSetGetPassword(t *testing.T) {
	attribute := "password"
	password := "sekret"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["password"] = password
	result, exists := pfs.GetPassword()

	if !exists || result != password {
		t.Fatalf("expected %s to be %s, got %s", attribute, password, result)
	}
}

func TestPDFFlagSetGetSmartShrinking(t *testing.T) {
	attribute := "smart-width"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["smart-width"] = value
	result, exists := pfs.GetSmartShrinking()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetStopSlowScripts(t *testing.T) {
	attribute := "stop-slow-scripts"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["stop-slow-scripts"] = value
	result, exists := pfs.GetStopSlowScripts()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetTitle(t *testing.T) {
	attribute := "title"
	title := "Some Title"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["title"] = title
	result, exists := pfs.GetTitle()

	if !exists || result != title {
		t.Fatalf("expected %s to be %s, got %s", attribute, title, result)
	}
}

func TestPDFFlagSetGetUseXServer(t *testing.T) {
	attribute := "use-xserver"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["use-xserver"] = value
	result, exists := pfs.GetUseXServer()

	if !exists || result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestPDFFlagSetGetUsername(t *testing.T) {
	attribute := "username"
	username := "some_name"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["username"] = username
	result, exists := pfs.GetUsername()

	if !exists || result != username {
		t.Fatalf("expected %s to be %s, got %s", attribute, username, result)
	}
}

func TestPDFFlagSetGetZoom(t *testing.T) {
	attribute := "zoom"
	zoom := 1.5
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs["zoom"] = zoom
	result, exists := pfs.GetZoom()

	if !exists || result != zoom {
		t.Fatalf("expected %s to be %f, got %f", attribute, zoom, result)
	}
}

func TestPDFFlagSetSetCacheDir(t *testing.T) {
	attribute := "cache-dir"
	dir := "/tmp/xyz"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetCacheDir(dir)

	if pfs[attribute] != dir {
		t.Fatalf("expected %s to be %s, got %s", attribute, dir, pfs[attribute])
	}
}

func TestPDFFlagSetSetCookie(t *testing.T) {
	attribute := "cookie"
	cookies := []wkhtmltox.CookieSet{
		wkhtmltox.CookieSet{
			Name:  "cookie1",
			Value: "value1",
		},
		wkhtmltox.CookieSet{
			Name:  "cookie2",
			Value: "value2",
		},
	}
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetCookie(cookies)

	if !reflect.DeepEqual(pfs[attribute], cookies) {
		t.Fatalf("expected %s to be %s, got %s", attribute, cookies, pfs[attribute])
	}
}

func TestPDFFlagSetSetCustomHeader(t *testing.T) {
	attribute := "custom-header"
	headers := []wkhtmltox.HeaderSet{
		wkhtmltox.HeaderSet{
			Name:  "header1",
			Value: "value1",
		},
		wkhtmltox.HeaderSet{
			Name:  "header2",
			Value: "value2",
		},
	}
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetCustomHeader(headers)

	if !reflect.DeepEqual(pfs[attribute], headers) {
		t.Fatalf("expected %s to be %s, got %s", attribute, headers, pfs[attribute])
	}
}

func TestPDFFlagSetSetCustomHeaderPropagation(t *testing.T) {
	attribute := "custom-header-propagation"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetCustomHeaderPropagation(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetDebugJavascript(t *testing.T) {
	attribute := "debug-javascript"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetDebugJavascript(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetDPI(t *testing.T) {
	attribute := "dpi"
	dpi := 600
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetDPI(dpi)

	if pfs[attribute] != dpi {
		t.Fatalf("expected %s to be %d, got %d", attribute, dpi, pfs[attribute])
	}
}

func TestPDFFlagSetSetEncoding(t *testing.T) {
	attribute := "encoding"
	encoding := "utf-8"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetEncoding(encoding)

	if pfs[attribute] != encoding {
		t.Fatalf("expected %s to be %s, got %s", attribute, encoding, pfs[attribute])
	}
}

func TestPDFFlagSetSetExternalLinks(t *testing.T) {
	attribute := "external-links"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetExternalLinks(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetForms(t *testing.T) {
	attribute := "forms"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetForms(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetGrayscale(t *testing.T) {
	attribute := "grayscale"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetGrayscale(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetImages(t *testing.T) {
	attribute := "images"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetImages(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetImageDPI(t *testing.T) {
	attribute := "image-dpi"
	dpi := 600
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetImageDPI(dpi)

	if pfs[attribute] != dpi {
		t.Fatalf("expected %s to be %d, got %d", attribute, dpi, pfs[attribute])
	}
}

func TestPDFFlagSetSetImageQuality(t *testing.T) {
	attribute := "image-quality"
	quality := 94
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetImageQuality(quality)

	if pfs[attribute] != quality {
		t.Fatalf("expected %s to be %d, got %d", attribute, quality, pfs[attribute])
	}
}

func TestPDFFlagSetSetInternalLinks(t *testing.T) {
	attribute := "internal-links"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetInternalLinks(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetJavascript(t *testing.T) {
	attribute := "javascript"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetJavascript(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetJavascriptDelay(t *testing.T) {
	attribute := "javascript-delay"
	milliseconds := 200
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetJavascriptDelay(milliseconds)

	if pfs[attribute] != milliseconds {
		t.Fatalf("expected %s to be %d, got %d", attribute, milliseconds, pfs[attribute])
	}
}

func TestPDFFlagSetSetLoadErrorHandling(t *testing.T) {
	attribute := "load-error-handling"
	handling := "abort"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetLoadErrorHandling(handling)

	if pfs[attribute] != handling {
		t.Fatalf("expected %s to be %s, got %s", attribute, handling, pfs[attribute])
	}
}

func TestPDFFlagSetSetLoadMediaErrorHandling(t *testing.T) {
	attribute := "load-media-error-handling"
	handling := "abort"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetLoadMediaErrorHandling(handling)

	if pfs[attribute] != handling {
		t.Fatalf("expected %s to be %s, got %s", attribute, handling, pfs[attribute])
	}
}

func TestPDFFlagSetSetLowQuality(t *testing.T) {
	attribute := "lowquality"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetLowQuality(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetMarginBottom(t *testing.T) {
	attribute := "margin-bottom"
	millimetres := 10
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetMarginBottom(millimetres)

	if pfs[attribute] != millimetres {
		t.Fatalf("expected %s to be %d, got %d", attribute, millimetres, pfs[attribute])
	}
}

func TestPDFFlagSetSetMarginLeft(t *testing.T) {
	attribute := "margin-left"
	millimetres := 10
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetMarginLeft(millimetres)

	if pfs[attribute] != millimetres {
		t.Fatalf("expected %s to be %d, got %d", attribute, millimetres, pfs[attribute])
	}
}

func TestPDFFlagSetSetMarginRight(t *testing.T) {
	attribute := "margin-right"
	millimetres := 10
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetMarginRight(millimetres)

	if pfs[attribute] != millimetres {
		t.Fatalf("expected %s to be %d, got %d", attribute, millimetres, pfs[attribute])
	}
}

func TestPDFFlagSetSetMarginTop(t *testing.T) {
	attribute := "margin-top"
	millimetres := 10
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetMarginTop(millimetres)

	if pfs[attribute] != millimetres {
		t.Fatalf("expected %s to be %d, got %d", attribute, millimetres, pfs[attribute])
	}
}

func TestPDFFlagSetSetMinimumFontSize(t *testing.T) {
	attribute := "minimum-font-size"
	size := 12
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetMinimumFontSize(size)

	if pfs[attribute] != size {
		t.Fatalf("expected %s to be %d, got %d", attribute, size, pfs[attribute])
	}
}

func TestPDFFlagSetSetNoPDFCompression(t *testing.T) {
	attribute := "no-pdf-compression"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetNoPDFCompression(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetOrientation(t *testing.T) {
	attribute := "orientation"
	orientation := "landscape"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetOrientation(orientation)

	if pfs[attribute] != orientation {
		t.Fatalf("expected %s to be %s, got %s", attribute, orientation, pfs[attribute])
	}
}

func TestPDFFlagSetSetPageHeight(t *testing.T) {
	attribute := "page-height"
	height := 300
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetPageHeight(height)

	if pfs[attribute] != height {
		t.Fatalf("expected %s to be %d, got %d", attribute, height, pfs[attribute])
	}
}

func TestPDFFlagSetSetPageSize(t *testing.T) {
	attribute := "page-size"
	size := "A4"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetPageSize(size)

	if pfs[attribute] != size {
		t.Fatalf("expected %s to be %s, got %s", attribute, size, pfs[attribute])
	}
}

func TestPDFFlagSetSetPageWidth(t *testing.T) {
	attribute := "page-width"
	width := 300
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetPageWidth(width)

	if pfs[attribute] != width {
		t.Fatalf("expected %s to be %d, got %d", attribute, width, pfs[attribute])
	}
}

func TestPDFFlagSetSetPassword(t *testing.T) {
	attribute := "password"
	password := "sekret"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetPassword(password)

	if pfs[attribute] != password {
		t.Fatalf("expected %s to be %s, got %s", attribute, password, pfs[attribute])
	}
}

func TestPDFFlagSetSetSmartShrinking(t *testing.T) {
	attribute := "smart-width"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetSmartShrinking(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetStopSlowScripts(t *testing.T) {
	attribute := "stop-slow-scripts"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetStopSlowScripts(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetTitle(t *testing.T) {
	attribute := "title"
	title := "Some Title"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetTitle(title)

	if pfs[attribute] != title {
		t.Fatalf("expected %s to be %s, got %s", attribute, title, pfs[attribute])
	}
}

func TestPDFFlagSetSetUseXServer(t *testing.T) {
	attribute := "use-xserver"
	value := true
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetUseXServer(value)

	if pfs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, pfs[attribute])
	}
}

func TestPDFFlagSetSetUsername(t *testing.T) {
	attribute := "username"
	username := "some_name"
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetUsername(username)

	if pfs[attribute] != username {
		t.Fatalf("expected %s to be %s, got %s", attribute, username, pfs[attribute])
	}
}

func TestPDFFlagSetSetZoom(t *testing.T) {
	attribute := "zoom"
	zoom := 1.5
	pfs := make(wkhtmltox.PDFFlagSet)
	pfs.SetZoom(zoom)

	if pfs[attribute] != zoom {
		t.Fatalf("expected %s to be %f, got %f", attribute, zoom, pfs[attribute])
	}
}
