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
	"fmt"
	"os/exec"
)

const (
	imageConverterBinary = "wkhtmltoimage"
)

// ImageFlagSet represents key-value pairs of image converter flags
type ImageFlagSet flagSet

// ImageOptions represents wkhtmlimage attributes
type ImageOptions struct {
	// General
	CropH  *int    `json:"crop_h,omitempty"` // Set height for cropping
	CropW  *int    `json:"crop_w,omitempty"` // Set width for cropping
	CropX  *int    `json:"crop_x,omitempty"` // Set x coordinate for cropping
	CropY  *int    `json:"crop_y,omitempty"` // Set y coordinate for cropping
	Format *string `json:"format,omitempty"` // Output file format
	Height *int    `json:"height,omitempty"` // Set screen height
	Width  *int    `json:"width,omitempty"`  // Set screen width
}

// FlagSet generates a FlagSet from ImageOptions
func (opts *ImageOptions) FlagSet() ImageFlagSet {
	ifs := make(ImageFlagSet)

	if opts.CropH != nil {
		ifs["crop-h"] = *opts.CropH
	}

	if opts.CropW != nil {
		ifs["crop-w"] = *opts.CropW
	}

	if opts.CropX != nil {
		ifs["crop-x"] = *opts.CropX
	}

	if opts.CropY != nil {
		ifs["crop-y"] = *opts.CropY
	}

	if opts.Format != nil {
		ifs["format"] = *opts.Format
	}

	if opts.Height != nil {
		ifs["height"] = *opts.Height
	}

	if opts.Width != nil {
		ifs["width"] = *opts.Width
	}

	return ifs
}

// Flags generates a String slice from a ImageFlagSet
func (ifs *ImageFlagSet) Flags() []string {
	var flags []string

	for flagKey, flagValue := range *ifs {
		flags = append(flags, fmt.Sprintf("--%s", flagKey), fmt.Sprintf("%v", flagValue))
	}

	return flags
}

// GetCropH retrieves the CropH from a ImageFlagSet
func (ifs *ImageFlagSet) GetCropH() (int, bool) {
	value, exists := (*ifs)["crop-h"]

	return value.(int), exists
}

// GetCropW retrieves the CropW from a ImageFlagSet
func (ifs *ImageFlagSet) GetCropW() (int, bool) {
	value, exists := (*ifs)["crop-w"]

	return value.(int), exists
}

// GetCropX retrieves the CropX from a ImageFlagSet
func (ifs *ImageFlagSet) GetCropX() (int, bool) {
	value, exists := (*ifs)["crop-x"]

	return value.(int), exists
}

// GetCropY retrieves the CropY from a ImageFlagSet
func (ifs *ImageFlagSet) GetCropY() (int, bool) {
	value, exists := (*ifs)["crop-y"]

	return value.(int), exists
}

// GetFormat retrieves the Format from a ImageFlagSet
func (ifs *ImageFlagSet) GetFormat() (string, bool) {
	value, exists := (*ifs)["format"]

	return value.(string), exists
}

// GetHeight retrieves the Height from a ImageFlagSet
func (ifs *ImageFlagSet) GetHeight() (int, bool) {
	value, exists := (*ifs)["height"]

	return value.(int), exists
}

// GetWidth retrieves the Width from a ImageFlagSet
func (ifs *ImageFlagSet) GetWidth() (int, bool) {
	value, exists := (*ifs)["width"]

	return value.(int), exists
}

// Generate performs the image conversion and saves the file to disk
func (ifs *ImageFlagSet) Generate(inputURL string, outputFile string) ([]byte, error) {
	var out []byte

	params := ifs.Flags()
	params = append(params, inputURL, outputFile)
	cmd := exec.Command(imageConverterBinary, params...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, err
}
