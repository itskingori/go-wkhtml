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

var imageInJSON = `
{
  "crop_h": 10,
  "crop_w": 10,
  "crop_x": 10,
  "crop_y": 10,
  "format": "png",
  "height": 10,
  "width": 10
}`

func TestNewImageOptionsFromJSON(t *testing.T) {
	d := []byte(imageInJSON)
	i, err := wkhtmltox.NewImageOptionsFromJSON(d)
	if err != nil {
		t.Fatal(err)
	}

	if *i.CropH != 10 {
		t.Fatalf("expected %s to be %d, got %d", "CropH", 10, *i.CropH)
	}

	if *i.CropW != 10 {
		t.Fatalf("expected %s to be %d, got %d", "CropW", 10, *i.CropW)
	}

	if *i.CropX != 10 {
		t.Fatalf("expected %s to be %d, got %d", "CropX", 10, *i.CropX)
	}

	if *i.CropY != 10 {
		t.Fatalf("expected %s to be %d, got %d", "CropY", 10, *i.CropY)
	}

	if *i.Format != "png" {
		t.Fatalf("expected %s to be %s, got %s", "Format", "png", *i.Format)
	}

	if *i.Height != 10 {
		t.Fatalf("expected %s to be %d, got %d", "Height", 10, *i.Height)
	}

	if *i.Width != 10 {
		t.Fatalf("expected %s to be %d, got %d", "Width", 10, *i.Width)
	}
}

func TestImageFlagSetsFlags(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSets)
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

func TestImageFlagSetsGetCropH(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSets)
	fss["crop-h"] = 10

	value, exists := fss.GetCropH()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-h", fss["crop-h"], value)
	}
}

func TestImageFlagSetsGetCropW(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSets)
	fss["crop-w"] = 10

	value, exists := fss.GetCropW()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-w", fss["crop-w"], value)
	}
}

func TestImageFlagSetsGetCropX(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSets)
	fss["crop-x"] = 10

	value, exists := fss.GetCropX()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-x", fss["crop-x"], value)
	}
}

func TestImageFlagSetsGetCropY(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSets)
	fss["crop-y"] = 10

	value, exists := fss.GetCropY()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-y", fss["crop-y"], value)
	}
}

func TestImageFlagSetsGetFormat(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSets)
	fss["format"] = "png"

	value, exists := fss.GetFormat()
	if !exists || value != "png" {
		t.Fatalf("expected %s to be %s, got %s", "format", fss["format"], value)
	}
}

func TestImageFlagSetsGetHeight(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSets)
	fss["height"] = 10

	value, exists := fss.GetHeight()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "height", fss["height"], value)
	}
}

func TestImageFlagSetsGetWidth(t *testing.T) {
	fss := make(wkhtmltox.ImageFlagSets)
	fss["width"] = 10

	value, exists := fss.GetWidth()
	if !exists || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "width", fss["width"], value)
	}
}
