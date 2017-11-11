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

	"encoding/json"
	"os/exec"
)

const (
	imageConverterBinary = "wkhtmltoimage"
)

// ImageFlagSets are key-value pair of flags
type ImageFlagSets map[string]interface{}

// ImageOptions represents a wkhtmlimage attributes
type ImageOptions struct {
	ImageGeneralOptions
}

// ImageGeneralOptions represents wkhtmltoimage general options
type ImageGeneralOptions struct {
	CropH *int `json:"crop_h,omitempty"` // Set height for cropping
	CropW *int `json:"crop_w,omitempty"` // Set width for cropping
	CropX *int `json:"crop_x,omitempty"` // Set x coordinate for cropping
	CropY *int `json:"crop_y,omitempty"` // Set y coordinate for cropping

	Format *string `json:"format,omitempty"` // Output file format
	Height *int    `json:"height,omitempty"` // Set screen height
	Width  *int    `json:"width,omitempty"`  // Set screen width
}

// NewImageOptionsFromJSON create image options from JSON
func NewImageOptionsFromJSON(data []byte) (*ImageOptions, error) {
	opts := &ImageOptions{}

	err := json.Unmarshal(data, &opts)
	if err != nil {
		return opts, err
	}

	return opts, nil
}

// Flags generates flag key-value-map pairs from ImageOptions
func (opts *ImageOptions) Flags() ImageFlagSets {
	ifs := make(ImageFlagSets)

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

// Flags converts a FlagSet to a String slice
func (ifs *ImageFlagSets) Flags() []string {
	var flags []string

	for flagKey, flagValue := range *ifs {
		flags = append(flags, fmt.Sprintf("--%s", flagKey), fmt.Sprintf("%v", flagValue))
	}

	return flags
}

// GetCropH retrieves the CropH from ImageFlagSets
func (ifs *ImageFlagSets) GetCropH() (int, bool) {
	value, exists := (*ifs)["crop-h"]

	return value.(int), exists
}

// GetCropW retrieves the CropW from ImageFlagSets
func (ifs *ImageFlagSets) GetCropW() (int, bool) {
	value, exists := (*ifs)["crop-w"]

	return value.(int), exists
}

// GetCropX retrieves the CropX from ImageFlagSets
func (ifs *ImageFlagSets) GetCropX() (int, bool) {
	value, exists := (*ifs)["crop-x"]

	return value.(int), exists
}

// GetCropY retrieves the CropY from ImageFlagSets
func (ifs *ImageFlagSets) GetCropY() (int, bool) {
	value, exists := (*ifs)["crop-y"]

	return value.(int), exists
}

// GetFormat retrieves the Format from ImageFlagSets
func (ifs *ImageFlagSets) GetFormat() (string, bool) {
	value, exists := (*ifs)["format"]

	return value.(string), exists
}

// GetHeight retrieves the Height from ImageFlagSets
func (ifs *ImageFlagSets) GetHeight() (int, bool) {
	value, exists := (*ifs)["height"]

	return value.(int), exists
}

// GetWidth retrieves the Width from ImageFlagSets
func (ifs *ImageFlagSets) GetWidth() (int, bool) {
	value, exists := (*ifs)["width"]

	return value.(int), exists
}

// GenerateImage creates an image based on the parameters passed in
func GenerateImage(ifs *ImageFlagSets, inputURL string, outputFile string) ([]byte, error) {
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
