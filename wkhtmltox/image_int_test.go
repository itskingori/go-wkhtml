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

package wkhtmltox

import (
	"testing"
)

func TestImageOptionsConvertToFlagSet(t *testing.T) {
	cropH := 10
	cropW := 10
	cropX := 10
	cropY := 10
	format := "png"
	height := 10
	width := 10

	opts := ImageOptions{
		CropH:  &cropH,
		CropW:  &cropW,
		CropX:  &cropX,
		CropY:  &cropY,
		Format: &format,
		Height: &height,
		Width:  &width,
	}
	flags := opts.convertToFlagSet()

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
