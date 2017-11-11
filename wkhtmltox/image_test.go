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

func TestImageOptionsFlagSet(t *testing.T) {
	cropH := 10
	cropW := 10
	cropX := 10
	cropY := 10
	format := "png"
	height := 10
	width := 10

	opts := wkhtmltox.ImageOptions{
		CropH:  &cropH,
		CropW:  &cropW,
		CropX:  &cropX,
		CropY:  &cropY,
		Format: &format,
		Height: &height,
		Width:  &width,
	}
	flags := opts.FlagSet()

	if flags["crop-h"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-h", 10, flags["crop-h"])
	}

	if flags["crop-w"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-w", 10, flags["crop-w"])
	}

	if flags["crop-x"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-x", 10, flags["crop-x"])
	}

	if flags["crop-y"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-y", 10, flags["crop-y"])
	}

	if flags["format"] != "png" {
		t.Fatalf("expected %s to be %s, got %s", "format", "png", flags["format"])
	}

	if flags["height"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "height", 10, flags["height"])
	}

	if flags["width"] != 10 {
		t.Fatalf("expected %s to be %d, got %d", "width", 10, flags["width"])
	}
}

func TestImageFlagSetFlags(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSet)
	fss["crop-h"] = 10
	fss["format"] = "png"

	expected := []string{"--crop-h", "10", "--format", "png"}
	got := fss.Flags()

	sort.Strings(expected)
	sort.Strings(got)

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}
}

func TestImageFlagSetGetCropH(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSet)
	fss["crop-h"] = 10

	value, exists := fss.GetCropH()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-h", fss["crop-h"], value)
	}
}

func TestImageFlagSetGetCropW(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSet)
	fss["crop-w"] = 10

	value, exists := fss.GetCropW()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-w", fss["crop-w"], value)
	}
}

func TestImageFlagSetGetCropX(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSet)
	fss["crop-x"] = 10

	value, exists := fss.GetCropX()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-x", fss["crop-x"], value)
	}
}

func TestImageFlagSetGetCropY(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSet)
	fss["crop-y"] = 10

	value, exists := fss.GetCropY()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-y", fss["crop-y"], value)
	}
}

func TestImageFlagSetGetFormat(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSet)
	fss["format"] = "png"

	value, exists := fss.GetFormat()
	if !exists || value != "png" {
		t.Fatalf("expected %s to be %s, got %s", "format", fss["format"], value)
	}
}

func TestImageFlagSetGetHeight(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSet)
	fss["height"] = 10

	value, exists := fss.GetHeight()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "height", fss["height"], value)
	}
}

func TestImageFlagSetGetWidth(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSet)
	fss["width"] = 10

	value, exists := fss.GetWidth()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "width", fss["width"], value)
	}
}
