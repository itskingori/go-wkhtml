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

func TestImageNewImageFlagSetFromOptions(t *testing.T) {
	format := "png"

	opts := wkhtmltox.ImageOptions{
		Format: &format,
	}
	ifs := wkhtmltox.NewImageFlagSetFromOptions(&opts)

	// when set
	if value, exists := ifs["format"]; !exists && value != "png" {
		t.Fatalf("expected %s to be %s, got %s", "format", "png", ifs["format"])
	}

	// when not set
	if _, exists := ifs["height"]; exists {
		t.Fatalf("expected %s to be %d, got %d", "height", 10, ifs["height"])
	}
}

func TestImageFlagSetFlags(t *testing.T) {
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs["cache-dir"] = "/some/dir"
	expected := []string{"--cache-dir", "/some/dir"}
	got := ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["cookie"] = []wkhtmltox.CookieSet{
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
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["crop-h"] = 10
	expected = []string{"--crop-h", "10"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["crop-w"] = 10
	expected = []string{"--crop-w", "10"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["crop-x"] = 10
	expected = []string{"--crop-x", "10"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["crop-y"] = 10
	expected = []string{"--crop-y", "10"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["custom-header"] = []wkhtmltox.HeaderSet{
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
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["custom-header-propagation"] = true
	expected = []string{"--custom-header-propagation"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["custom-header-propagation"] = false
	expected = []string{"--no-custom-header-propagation"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["debug-javascript"] = true
	expected = []string{"--debug-javascript"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["debug-javascript"] = false
	expected = []string{"--no-debug-javascript"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["encoding"] = "utf-8"
	expected = []string{"--encoding", "utf-8"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["format"] = "png"
	expected = []string{"--format", "png"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["height"] = 10
	expected = []string{"--height", "10"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["images"] = true
	expected = []string{"--images"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["images"] = false
	expected = []string{"--no-images"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["javascript"] = false
	expected = []string{"--disable-javascript"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["javascript"] = true
	expected = []string{"--enable-javascript"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["javascript-delay"] = 200
	expected = []string{"--javascript-delay", "200"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["load-error-handling"] = "abort"
	expected = []string{"--load-error-handling", "abort"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["load-media-error-handling"] = "abort"
	expected = []string{"--load-media-error-handling", "abort"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["minimum-font-size"] = 10
	expected = []string{"--minimum-font-size", "10"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["password"] = "sekret"
	expected = []string{"--password", "sekret"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["quality"] = 94
	expected = []string{"--quality", "94"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["smart-width"] = false
	expected = []string{"--disable-smart-width"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["smart-width"] = true
	expected = []string{"--enable-smart-width"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["stop-slow-scripts"] = true
	expected = []string{"--stop-slow-scripts"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["stop-slow-scripts"] = false
	expected = []string{"--no-stop-slow-scripts"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["transparent"] = true
	expected = []string{"--transparent"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["transparent"] = false
	got = ifs.Flags()
	if len(got) != 0 {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["use-xserver"] = true
	expected = []string{"--use-xserver"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["use-xserver"] = false
	got = ifs.Flags()
	if len(got) != 0 {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["username"] = "sekret"
	expected = []string{"--username", "sekret"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["width"] = 10
	expected = []string{"--width", "10"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

	ifs = make(wkhtmltox.ImageFlagSet)
	ifs["zoom"] = 0.5
	expected = []string{"--zoom", "0.5"}
	got = ifs.Flags()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}

}

func TestImageFlagSetGetCacheDir(t *testing.T) {
	attribute := "cache-dir"
	dir := "/tmp/xyz"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = dir
	result, exists := ifs.GetCacheDir()

	if !exists && result != dir {
		t.Fatalf("expected %s to be %s, got %s", attribute, dir, result)
	}
}

func TestImageFlagSetGetCookie(t *testing.T) {
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
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = cookies

	result, exists := ifs.GetCookie()
	if !exists && !reflect.DeepEqual(result, cookies) {
		t.Fatalf("expected %s to be %s, got %s", attribute, cookies, result)
	}
}

func TestImageFlagSetGetCropH(t *testing.T) {
	attribute := "crop-h"
	height := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = height
	result, exists := ifs.GetCropH()

	if !exists && result != height {
		t.Fatalf("expected %s to be %d, got %d", attribute, height, result)
	}
}

func TestImageFlagSetGetCropW(t *testing.T) {
	attribute := "crop-w"
	width := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = width
	result, exists := ifs.GetCropW()

	if !exists && result != width {
		t.Fatalf("expected %s to be %d, got %d", attribute, width, result)
	}
}

func TestImageFlagSetGetCropX(t *testing.T) {
	attribute := "crop-x"
	xc := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = xc
	result, exists := ifs.GetCropX()

	if !exists && result != xc {
		t.Fatalf("expected %s to be %d, got %d", attribute, xc, result)
	}
}

func TestImageFlagSetGetCropY(t *testing.T) {
	attribute := "crop-y"
	yc := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = yc
	result, exists := ifs.GetCropY()

	if !exists && result != yc {
		t.Fatalf("expected %s to be %d, got %d", attribute, yc, result)
	}
}

func TestImageFlagSetGetCustomHeader(t *testing.T) {
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
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = headers
	result, exists := ifs.GetCustomHeader()

	if !exists && !reflect.DeepEqual(result, headers) {
		t.Fatalf("expected %s to be %s, got %s", attribute, headers, result)
	}
}

func TestImageFlagSetGetCustomHeaderPropagation(t *testing.T) {
	attribute := "custom-header-propagation"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = value
	result, exists := ifs.GetCustomHeaderPropagation()

	if !exists && result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestImageFlagSetGetDebugJavascript(t *testing.T) {
	attribute := "debug-javascript"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = value
	result, exists := ifs.GetDebugJavascript()

	if !exists && result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestImageFlagSetGetEncoding(t *testing.T) {
	attribute := "encoding"
	encoding := "utf-8"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = encoding
	result, exists := ifs.GetEncoding()

	if !exists && result != encoding {
		t.Fatalf("expected %s to be %s, got %s", attribute, encoding, result)
	}
}

func TestImageFlagSetGetFormat(t *testing.T) {
	attribute := "format"
	format := "png"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = format
	result, exists := ifs.GetFormat()

	if !exists && result != format {
		t.Fatalf("expected %s to be %s, got %s", attribute, format, result)
	}
}

func TestImageFlagSetGetHeight(t *testing.T) {
	attribute := "height"
	height := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = height
	result, exists := ifs.GetHeight()

	if !exists && result != height {
		t.Fatalf("expected %s to be %d, got %d", attribute, height, result)
	}
}

func TestImageFlagSetGetImages(t *testing.T) {
	attribute := "images"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = value
	result, exists := ifs.GetImages()

	if !exists && result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestImageFlagSetGetJavascript(t *testing.T) {
	attribute := "javascript"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = value
	result, exists := ifs.GetJavascript()

	if !exists && result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestImageFlagSetGetJavascriptDelay(t *testing.T) {
	attribute := "javascript-delay"
	milliseconds := 200
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = milliseconds
	result, exists := ifs.GetJavascriptDelay()

	if !exists && result != milliseconds {
		t.Fatalf("expected %s to be %d, got %d", attribute, milliseconds, result)
	}
}

func TestImageFlagSetGetLoadErrorHandling(t *testing.T) {
	attribute := "load-error-handling"
	handling := "abort"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = handling
	result, exists := ifs.GetLoadErrorHandling()

	if !exists && result != handling {
		t.Fatalf("expected %s to be %s, got %s", attribute, handling, result)
	}
}

func TestImageFlagSetGetLoadMediaErrorHandling(t *testing.T) {
	attribute := "load-media-error-handling"
	handling := "abort"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = handling
	result, exists := ifs.GetLoadMediaErrorHandling()

	if !exists && result != handling {
		t.Fatalf("expected %s to be %s, got %s", attribute, handling, result)
	}
}

func TestImageFlagSetGetMinimumFontSize(t *testing.T) {
	attribute := "minimum-font-size"
	size := 12
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = size
	result, exists := ifs.GetMinimumFontSize()

	if !exists && result != size {
		t.Fatalf("expected %s to be %d, got %d", attribute, size, result)
	}
}

func TestImageFlagSetGetPassword(t *testing.T) {
	attribute := "password"
	password := "sekret"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = password
	result, exists := ifs.GetPassword()

	if !exists && result != password {
		t.Fatalf("expected %s to be %s, got %s", attribute, password, result)
	}
}

func TestImageFlagSetGetQuality(t *testing.T) {
	attribute := "quality"
	quality := 95
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = quality
	result, exists := ifs.GetQuality()

	if !exists && result != quality {
		t.Fatalf("expected %s to be %d, got %d", attribute, quality, result)
	}
}

func TestImageFlagSetGetSmartWidth(t *testing.T) {
	attribute := "smart-width"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = value
	result, exists := ifs.GetSmartWidth()

	if !exists && result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestImageFlagSetGetStopSlowScripts(t *testing.T) {
	attribute := "stop-slows-cripts"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = value
	result, exists := ifs.GetStopSlowScripts()

	if !exists && result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestImageFlagSetGetTransparent(t *testing.T) {
	attribute := "transparent"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = value
	result, exists := ifs.GetTransparent()

	if !exists && result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestImageFlagSetGetUseXServer(t *testing.T) {
	attribute := "use-xserver"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = value
	result, exists := ifs.GetUseXServer()

	if !exists && result != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, result)
	}
}

func TestImageFlagSetGetUsername(t *testing.T) {
	attribute := "username"
	username := "some_name"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = username
	result, exists := ifs.GetUsername()

	if !exists && result != username {
		t.Fatalf("expected %s to be %s, got %s", attribute, username, result)
	}
}

func TestImageFlagSetGetWidth(t *testing.T) {
	attribute := "width"
	width := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = width
	result, exists := ifs.GetWidth()

	if !exists && result != width {
		t.Fatalf("expected %s to be %d, got %d", attribute, width, result)
	}
}

func TestImageFlagSetGetZoom(t *testing.T) {
	attribute := "zoom"
	zoom := 0.2
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs[attribute] = zoom
	result, exists := ifs.GetZoom()

	if !exists && result != zoom {
		t.Fatalf("expected %s to be %f, got %f", attribute, zoom, result)
	}
}

func TestImageFlagSetSetCacheDir(t *testing.T) {
	attribute := "cache-dir"
	dir := "/tmp/xyz"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetCacheDir(dir)

	if ifs[attribute] != dir {
		t.Fatalf("expected %s to be %s, got %s", attribute, dir, ifs[attribute])
	}
}

func TestImageFlagSetSetCookie(t *testing.T) {
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
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetCookie(cookies)

	if !reflect.DeepEqual(ifs[attribute], cookies) {
		t.Fatalf("expected %s to be %s, got %s", attribute, cookies, ifs[attribute])
	}
}

func TestImageFlagSetSetCropH(t *testing.T) {
	attribute := "crop-h"
	height := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetCropH(height)

	if ifs[attribute] != height {
		t.Fatalf("expected %s to be %d, got %d", attribute, height, ifs[attribute])
	}
}

func TestImageFlagSetSetCropW(t *testing.T) {
	attribute := "crop-w"
	width := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetCropW(width)

	if ifs[attribute] != width {
		t.Fatalf("expected %s to be %d, got %d", attribute, width, ifs[attribute])
	}
}

func TestImageFlagSetSetCropX(t *testing.T) {
	attribute := "crop-x"
	xc := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetCropX(xc)

	if ifs[attribute] != xc {
		t.Fatalf("expected %s to be %d, got %d", attribute, xc, ifs[attribute])
	}
}

func TestImageFlagSetSetCropY(t *testing.T) {
	attribute := "crop-y"
	yc := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetCropY(yc)

	if ifs[attribute] != yc {
		t.Fatalf("expected %s to be %d, got %d", attribute, yc, ifs[attribute])
	}
}

func TestImageFlagSetSetCustomHeader(t *testing.T) {
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
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetCustomHeader(headers)

	if !reflect.DeepEqual(ifs[attribute], headers) {
		t.Fatalf("expected %s to be %s, got %s", attribute, headers, ifs[attribute])
	}
}

func TestImageFlagSetSetCustomHeaderPropagation(t *testing.T) {
	attribute := "custom-header-propagation"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetCustomHeaderPropagation(value)

	if ifs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, ifs[attribute])
	}
}

func TestImageFlagSetSetDebugJavascript(t *testing.T) {
	attribute := "debug-javascript"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetDebugJavascript(value)

	if ifs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, ifs[attribute])
	}
}

func TestImageFlagSetSetEncoding(t *testing.T) {
	attribute := "encoding"
	encoding := "utf-8"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetEncoding(encoding)

	if ifs[attribute] != encoding {
		t.Fatalf("expected %s to be %s, got %s", attribute, encoding, ifs[attribute])
	}
}

func TestImageFlagSetSetFormat(t *testing.T) {
	attribute := "format"
	format := "png"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetFormat(format)

	if ifs[attribute] != format {
		t.Fatalf("expected %s to be %s, got %s", attribute, format, ifs[attribute])
	}
}

func TestImageFlagSetSetHeight(t *testing.T) {
	attribute := "height"
	height := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetHeight(height)

	if ifs[attribute] != height {
		t.Fatalf("expected %s to be %d, got %d", attribute, height, ifs[attribute])
	}
}

func TestImageFlagSetSetImages(t *testing.T) {
	attribute := "images"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetImages(value)

	if ifs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, ifs[attribute])
	}
}

func TestImageFlagSetSetJavascript(t *testing.T) {
	attribute := "javascript"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetJavascript(value)

	if ifs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, ifs[attribute])
	}
}

func TestImageFlagSetSetJavascriptDelay(t *testing.T) {
	attribute := "javascript-delay"
	milliseconds := 200
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetJavascriptDelay(milliseconds)

	if ifs[attribute] != milliseconds {
		t.Fatalf("expected %s to be %d, got %d", attribute, milliseconds, ifs[attribute])
	}
}

func TestImageFlagSetSetLoadErrorHandling(t *testing.T) {
	attribute := "load-error-handling"
	handling := "abort"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetLoadErrorHandling(handling)

	if ifs[attribute] != handling {
		t.Fatalf("expected %s to be %s, got %s", attribute, handling, ifs[attribute])
	}
}

func TestImageFlagSetSetLoadMediaErrorHandling(t *testing.T) {
	attribute := "load-media-error-handling"
	handling := "abort"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetLoadMediaErrorHandling(handling)

	if ifs[attribute] != handling {
		t.Fatalf("expected %s to be %s, got %s", attribute, handling, ifs[attribute])
	}
}

func TestImageFlagSetSetMinimumFontSize(t *testing.T) {
	attribute := "minimum-font-size"
	size := 12
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetMinimumFontSize(size)

	if ifs[attribute] != size {
		t.Fatalf("expected %s to be %d, got %d", attribute, size, ifs[attribute])
	}
}

func TestImageFlagSetSetPassword(t *testing.T) {
	attribute := "password"
	password := "sekret"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetPassword(password)

	if ifs[attribute] != password {
		t.Fatalf("expected %s to be %s, got %s", attribute, password, ifs[attribute])
	}
}

func TestImageFlagSetSetQuality(t *testing.T) {
	attribute := "quality"
	quality := 95
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetQuality(quality)

	if ifs[attribute] != quality {
		t.Fatalf("expected %s to be %d, got %d", attribute, quality, ifs[attribute])
	}
}

func TestImageFlagSetSetSmartWidth(t *testing.T) {
	attribute := "smart-width"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetSmartWidth(value)

	if ifs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, ifs[attribute])
	}
}

func TestImageFlagSetSetStopSlowScripts(t *testing.T) {
	attribute := "stop-slows-cripts"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetStopSlowScripts(value)

	if ifs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, ifs[attribute])
	}
}

func TestImageFlagSetSetTransparent(t *testing.T) {
	attribute := "transparent"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetTransparent(value)

	if ifs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, ifs[attribute])
	}
}

func TestImageFlagSetSetUseXServer(t *testing.T) {
	attribute := "use-xserver"
	value := true
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetUseXServer(value)

	if ifs[attribute] != value {
		t.Fatalf("expected %s to be %t, got %t", attribute, value, ifs[attribute])
	}
}

func TestImageFlagSetSetUsername(t *testing.T) {
	attribute := "username"
	username := "some_name"
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetUsername(username)

	if ifs[attribute] != username {
		t.Fatalf("expected %s to be %s, got %s", attribute, username, ifs[attribute])
	}
}

func TestImageFlagSetSetWidth(t *testing.T) {
	attribute := "width"
	width := 10
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetWidth(width)

	if ifs[attribute] != width {
		t.Fatalf("expected %s to be %d, got %d", attribute, width, ifs[attribute])
	}
}

func TestImageFlagSetSetZoom(t *testing.T) {
	attribute := "zoom"
	zoom := 0.2
	ifs := make(wkhtmltox.ImageFlagSet)
	ifs.SetZoom(zoom)

	if ifs[attribute] != zoom {
		t.Fatalf("expected %s to be %f, got %f", attribute, zoom, ifs[attribute])
	}
}
