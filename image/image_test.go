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

package image_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/itskingori/go-wkhtml/image"
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

func TestNewOptionsFromJSON(t *testing.T) {
	d := []byte(imageInJSON)
	i, err := image.NewOptionsFromJSON(d)
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

func TestFlags(t *testing.T) {
	cropH := 10
	cropW := 10
	cropX := 10
	cropY := 10
	format := "png"
	height := 10
	width := 10

	gopts := image.GeneralOptions{
		CropH:  &cropH,
		CropW:  &cropW,
		CropX:  &cropX,
		CropY:  &cropY,
		Format: &format,
		Height: &height,
		Width:  &width,
	}

	opts := image.Options{GeneralOptions: gopts}
	flags := opts.Flags()

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

func TestFlagSetFlags(t *testing.T) {
	fss := make(image.FlagSets)
	fss["crop-h"] = 10
	fss["format"] = "png"

	expected := []string{"--crop-h=10", "--format=png"}
	got := fss.Flags()

	sort.Strings(expected)
	sort.Strings(got)

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected '%s' but got '%s'", expected, got)
	}
}

func TestGetCropH(t *testing.T) {
	fss := make(image.FlagSets)
	fss["crop-h"] = 10

	value, exists := fss.GetCropH()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-h", fss["crop-h"], value)
	}
}

func TestGetCropW(t *testing.T) {
	fss := make(image.FlagSets)
	fss["crop-w"] = 10

	value, exists := fss.GetCropW()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-w", fss["crop-w"], value)
	}
}

func TestGetCropX(t *testing.T) {
	fss := make(image.FlagSets)
	fss["crop-x"] = 10

	value, exists := fss.GetCropX()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-x", fss["crop-x"], value)
	}
}

func TestGetCropY(t *testing.T) {
	fss := make(image.FlagSets)
	fss["crop-y"] = 10

	value, exists := fss.GetCropY()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "crop-y", fss["crop-y"], value)
	}
}

func TestGetFormat(t *testing.T) {
	fss := make(image.FlagSets)
	fss["format"] = "png"

	value, exists := fss.GetFormat()
	if exists != true || value != "png" {
		t.Fatalf("expected %s to be %s, got %s", "format", fss["format"], value)
	}
}

func TestGetHeight(t *testing.T) {
	fss := make(image.FlagSets)
	fss["height"] = 10

	value, exists := fss.GetHeight()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "height", fss["height"], value)
	}
}

func TestGetWidth(t *testing.T) {
	fss := make(image.FlagSets)
	fss["width"] = 10

	value, exists := fss.GetWidth()
	if exists != true || value != 10 {
		t.Fatalf("expected %s to be %d, got %d", "width", fss["width"], value)
	}
}
